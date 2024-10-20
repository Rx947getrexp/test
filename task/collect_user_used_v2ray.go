package task

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/types/api"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service/api/speed_api"
	"net/http/httptest"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	TaskTypeCollectUserUsed2ray = "node-user"

	TaskStatusFinished = 1

	TaskIPAll = "node-all"
)

func getYesterday(format string) string {
	location, _ := time.LoadLocation("Europe/Moscow")
	now := time.Now().In(location)
	return now.AddDate(0, 0, -1).Format(format)
}

func isTaskSuccess(date int) bool {
	var task *entity.TTask
	err := dao.TTask.Ctx(context.Background()).Where(do.TTask{
		Ip:   TaskIPAll,
		Date: date,
		Type: TaskTypeCollectUserUsed2ray,
	}).Scan(&task)
	if err != nil {
		global.Logger.Err(err).Msgf("get %s failed", TaskIPAll)
		return false
	}
	if task != nil {
		return true
	} else {
		return false
	}
}

func CollectUserUsedV2ray() {
	global.Recovery()
	global.Logger.Info().Msg("CollectUserTraffic start...")
	for {
		location, _ := time.LoadLocation("Europe/Moscow")
		yesterdayTime := time.Now().In(location).AddDate(0, 0, -1)

		yesterdayDate, _ := strconv.Atoi(yesterdayTime.Format("20060102"))
		if isTaskSuccess(yesterdayDate) {
			global.Logger.Info().Msgf("%d task already success. start sleep ...", yesterdayDate)
		} else {
			// 获取莫斯科的时区
			location, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				global.Logger.Err(err).Msgf("time.LoadLocation failed")
			} else {
				if time.Now().In(location).Hour() < 4 {
					global.Logger.Info().Msgf("时间未到莫斯科时间 凌晨4点 以后开始执行")
				} else {
					doCollectUserUsedV2ray(yesterdayTime)
				}
			}
		}

		global.Logger.Info().Msgf("start sleep ...")
		time.Sleep(intervalTraffic)
	}
}

func doCollectUserUsedV2ray(yesterdayTime time.Time) {
	// 查询v2ray数据节点
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	resp, err := speed_api.DescribeNodeList(ctx)
	if err != nil {
		global.Logger.Err(err).Msg("get node ip list failed")
		return
	}

	wg := &sync.WaitGroup{}
	for i, _ := range resp.Items {
		wg.Add(1)
		go doCollectUserUsedV2rayOneNode(ctx, wg, resp.Items[i], yesterdayTime)
	}
	wg.Wait()

	var cnt uint
	for _, item := range resp.Items {
		task, err := getIpTask(ctx, item.Ip, yesterdayTime)
		if err != nil {
			return
		}
		if task == nil {
			global.Logger.Warn().Msgf("还有节点没有统计完成")
			return
		}
		cnt += task.UserCnt
	}

	_, err = dao.TTask.Ctx(ctx).Data(do.TTask{
		Ip:        TaskIPAll,
		Date:      yesterdayTime.Format("20060102"),
		UserCnt:   cnt,
		Status:    TaskStatusFinished,
		Type:      TaskTypeCollectUserUsed2ray,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		global.Logger.Err(err).Msgf(
			"save TTask failed, date: %s, ip: %s", yesterdayTime.Format("20060102"), TaskIPAll)
		return
	}
}

func getIpTask(ctx context.Context, ip string, yesterdayTime time.Time) (taskEntity *entity.TTask, err error) {
	var (
		yesterdayDate, _ = strconv.Atoi(yesterdayTime.Format("20060102"))

		taskDesc = fmt.Sprintf("%s_%d_%s", ip, yesterdayDate, TaskTypeCollectUserUsed2ray)
	)
	err = dao.TTask.Ctx(ctx).Where(do.TTask{
		Ip:   ip,
		Date: yesterdayDate,
		Type: TaskTypeCollectUserUsed2ray,
	}).Scan(&taskEntity)
	if err != nil {
		global.Logger.Err(err).Msgf("get t_task failed, %s", taskDesc)
		return nil, err
	}
	if taskEntity == nil {
		global.Logger.Info().Msgf("%s t_task not found", taskDesc)
	} else {
		global.Logger.Info().Msgf("%s have finished", taskDesc)
	}
	return taskEntity, nil
}

func doCollectUserUsedV2rayOneNode(ctx context.Context, wg *sync.WaitGroup, node api.NodeItem, yesterdayTime time.Time) {
	defer wg.Done()

	var (
		err              error
		yesterdayDate, _ = strconv.Atoi(yesterdayTime.Format("20060102"))
		taskEntity       *entity.TTask
		taskDesc         = fmt.Sprintf("%s_%d_%s", node.Ip, yesterdayDate, TaskTypeCollectUserUsed2ray)
	)
	err = dao.TTask.Ctx(ctx).Where(do.TTask{
		Ip:   node.Ip,
		Date: yesterdayDate,
		Type: TaskTypeCollectUserUsed2ray,
	}).Scan(&taskEntity)
	if err != nil {
		global.Logger.Err(err).Msgf("get t_task failed, %s", taskDesc)
		return
	}
	if taskEntity != nil {
		global.Logger.Info().Msgf("%s already finished", taskDesc)
		return
	}

	emails, err := collectUserUsedV2rayByIp(node.Ip, yesterdayTime)
	if err != nil {
		global.Logger.Err(err).Msgf("collectUserUsedV2rayByIp failed, %s", taskDesc)
		return
	}

	for _, email := range emails {
		_, err = dao.TV2RayUserTraffic.Ctx(ctx).Data(do.TV2RayUserTraffic{
			Email:     email,
			Date:      yesterdayDate,
			Ip:        node.Ip,
			Uplink:    1,
			Downlink:  1,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}).Save()
		if err != nil {
			global.Logger.Err(err).Msgf(
				"save TV2RayUserTraffic failed, email: %s, date: %d, ip: %s",
				email, yesterdayDate, node.Ip)
			return
		}
		global.Logger.Info().Msgf(
			"save TV2RayUserTraffic success, email: %s, date: %d, ip: %s",
			email, yesterdayDate, node.Ip)
	}

	_, err = dao.TTask.Ctx(ctx).Data(do.TTask{
		Ip:        node.Ip,
		Date:      yesterdayDate,
		UserCnt:   len(emails),
		Status:    TaskStatusFinished,
		Type:      TaskTypeCollectUserUsed2ray,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		global.Logger.Err(err).Msgf("save TTask failed, date: %d, ip: %s", yesterdayDate, node.Ip)
		return
	}
}

func collectUserUsedV2rayByIp(ip string, yesterdayTime time.Time) (emails []string, err error) {
	cmdName := "ssh"
	cmdArgs := []string{
		fmt.Sprintf("root@%s", ip),
		fmt.Sprintf(
			"grep -a '%s' /var/log/v2ray/access.log | grep email | grep accepted | awk '{print $7}' | sort | uniq",
			yesterdayTime.Format("2006/01/02")),
	}
	global.Logger.Debug().Msgf("ssh cmdArgs: %+v", cmdArgs)

	// 创建Cmd结构体
	cmd := exec.Command(cmdName, cmdArgs...)

	// 创建一个缓冲区来存储命令的输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 运行命令
	err = cmd.Run()
	if err != nil {
		return nil, gerror.Wrap(err, "cmd.Run() failed")
	}
	// 打印命令的输出
	global.Logger.Debug().Msgf("Output: %+v", out.String())

	// 将输出的字符串按照换行符分割，然后将结果存储到一个[]string变量中
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			emails = append(emails, line)
		}
	}
	global.Logger.Debug().Msgf("emails: %+v", emails)
	return
}
