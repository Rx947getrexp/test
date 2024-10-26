package tools

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/spf13/cobra"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service/api/speed_api"
	"log"
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

var nDaysBefore uint32
var insertDB bool
var printDate bool

func init() {
	countUserAccessV2ray.Flags().StringVarP(&ip, "ip", "i", "", "节点IP(optional), 不传则为统计全部节点")
	countUserAccessV2ray.Flags().Uint32VarP(&nDaysBefore, "n_days_before", "d", 0, "统计几天前的数据(optional)，统计今天则不用设置，昨天为1，前天为2，以此类推...")
	countUserAccessV2ray.Flags().BoolVarP(&insertDB, "save_into_db", "s", false, "是否需保存到DB(optional), bool类型")
	countUserAccessV2ray.Flags().BoolVarP(&printDate, "print_date", "p", false, "打印数据日期(optional), bool类型")
	AddConfigCmd.AddCommand(countUserAccessV2ray)
}

var countUserAccessV2ray = &cobra.Command{
	Use:   "countUserAccess",
	Short: "count user access",
	Run: func(cmd *cobra.Command, args []string) {
		doCountUserAccessV2ray()
	},
}

func doCountUserAccessV2ray() {
	//ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//defer ctx.Done()

	//fmt.Printf("ip: %s\n", ip)

	//ips := getNodeIps(ip)

	// 获取当前时间
	now := time.Now()

	// 计算昨天的日期
	yesterday := now.AddDate(0, 0, -1*int(nDaysBefore))
	if printDate {
		fmt.Printf("data: %s\n", yesterday.Format("2006/01/02"))
		return
	}

	doCollectUserUsedV2ray(yesterday)

	//// 获取年、月和日
	////year, month, day := yesterday.Date()
	//
	//for _, ip := range ips {
	//	// 定义命令和参数
	//	//cmd := exec.Command("scp",
	//	//	fmt.Sprintf("root@%s:/usr/local/etc/v2ray/config.json", ip),
	//	//	fmt.Sprintf("/wwwroot/go/go-api/config/v2ray/users-%s.json", ip),
	//	//)
	//
	//	cmdName := "ssh"
	//	cmdArgs := []string{
	//		fmt.Sprintf("root@%s", ip),
	//		fmt.Sprintf("grep -a '%s' /var/log/v2ray/access.log | grep email | grep accepted | awk '{print $7}' | sort | uniq", yesterday.Format("2006/01/02")),
	//	}
	//	fmt.Println(cmdName, cmdArgs)
	//
	//	// 创建Cmd结构体
	//	cmd := exec.Command(cmdName, cmdArgs...)
	//
	//	// 创建一个缓冲区来存储命令的输出
	//	var out bytes.Buffer
	//	cmd.Stdout = &out
	//
	//	// 运行命令
	//	err := cmd.Run()
	//	if err != nil {
	//		log.Fatalf("cmd.Run() failed with %s\n", err)
	//	}
	//
	//	// 打印命令的输出
	//	fmt.Printf("Output: %q\n", out.String())
	//
	//	// 将输出的字符串按照换行符分割，然后将结果存储到一个[]string变量中
	//	lines := strings.Split(out.String(), "\n")
	//	var emails []string
	//	for _, line := range lines {
	//		emails = append(emails, line)
	//	}
	//	fmt.Printf("emails: %q\n", emails)
	//}
}

func doCollectUserUsedV2ray(yesterdayTime time.Time) {
	// 查询v2ray数据节点
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	var ips []string
	if ip == "" {
		resp, err := speed_api.DescribeNodeList(ctx)
		if err != nil {
			log.Fatalf("DescribeNodeList failed with %s\n", err.Error())
		}
		for _, item := range resp.Items {
			ips = append(ips, item.Ip)
		}
	} else {
		ips = append(ips, ip)
	}

	wg := &sync.WaitGroup{}
	for _, ip := range ips {
		wg.Add(1)
		go doCollectUserUsedV2rayOneNode(ctx, wg, ip, yesterdayTime)
	}
	wg.Wait()

	var cnt uint
	for _, ip := range ips {
		task, err := getIpTask(ctx, ip, yesterdayTime)
		if err != nil {
			log.Fatalf("getIpTask failed with %s\n", err.Error())
		}
		if task == nil {
			log.Fatalf("还有节点没有统计完成\n")
		}
		cnt += task.UserCnt
	}

	_, err := dao.TTask.Ctx(ctx).Data(do.TTask{
		Ip:        TaskIPAll,
		Date:      yesterdayTime.Format("20060102"),
		UserCnt:   cnt,
		Status:    TaskStatusFinished,
		Type:      TaskTypeCollectUserUsed2ray,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		log.Fatalf("save TTask failed with %s\n", err.Error())
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
		log.Fatalf("get t_task failed, %s, %s\n", taskDesc, err.Error())
	}
	if taskEntity == nil {
		fmt.Printf("%s t_task not found\n", taskDesc)
	} else {
		fmt.Printf("%s t_task have finished\n", taskDesc)
	}
	return taskEntity, nil
}

func doCollectUserUsedV2rayOneNode(ctx context.Context, wg *sync.WaitGroup, ip string, yesterdayTime time.Time) {
	defer wg.Done()

	var (
		err              error
		yesterdayDate, _ = strconv.Atoi(yesterdayTime.Format("20060102"))
		taskEntity       *entity.TTask
		taskDesc         = fmt.Sprintf("%s_%d_%s", ip, yesterdayDate, TaskTypeCollectUserUsed2ray)
	)
	err = dao.TTask.Ctx(ctx).Where(do.TTask{
		Ip:   ip,
		Date: yesterdayDate,
		Type: TaskTypeCollectUserUsed2ray,
	}).Scan(&taskEntity)
	if err != nil {
		log.Fatalf("get t_task failed, %s, %s\n", taskDesc, err.Error())
	}
	if taskEntity != nil {
		fmt.Printf("%s already finished\n", taskDesc)
		return
	}

	emails, err := collectUserUsedV2rayByIp(ip, yesterdayTime)
	if err != nil {
		log.Fatalf("collectUserUsedV2rayByIp failed with, %s, %s\n", taskDesc, err.Error())
	}

	for _, email := range emails {
		_, err = dao.TV2RayUserTraffic.Ctx(ctx).Data(do.TV2RayUserTraffic{
			Email:     email,
			Date:      yesterdayDate,
			Ip:        ip,
			Uplink:    1,
			Downlink:  1,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}).Save()
		if err != nil {
			log.Fatalf("save TV2RayUserTraffic failed, taskDesc: %s, %s\n", taskDesc, err.Error())
		}
		fmt.Printf("%s save TV2RayUserTraffic success\n", taskDesc)
	}

	_, err = dao.TTask.Ctx(ctx).Data(do.TTask{
		Ip:        ip,
		Date:      yesterdayDate,
		UserCnt:   len(emails),
		Status:    TaskStatusFinished,
		Type:      TaskTypeCollectUserUsed2ray,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		log.Fatalf("save TTask failed, taskDesc: %s, err: %s", taskDesc, err.Error())
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
	fmt.Printf("ssh cmdArgs: %+v\n", cmdArgs)

	// 创建Cmd结构体
	cmd := exec.Command(cmdName, cmdArgs...)

	// 创建一个缓冲区来存储命令的输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 运行命令
	err = cmd.Run()
	if err != nil {
		log.Fatalf("ip: %s, cmd.Run() failed: %+v\n", ip, err.Error())
	}
	// 打印命令的输出
	//global.Logger.Debug().Msgf("Output: %+v", out.String())

	// 将输出的字符串按照换行符分割，然后将结果存储到一个[]string变量中
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			emails = append(emails, line)
		}
	}
	fmt.Printf("ip: %s, emails: %+v\n", ip, emails)
	return
}

/*
+----------------+
| ip             |
+----------------+
| 110.42.42.229  |
| 147.45.178.51  |
| 154.93.104.115 |
| 185.39.207.104 |
| 185.39.207.20  |
| 193.124.41.88  |
| 193.233.48.70  |
| 207.90.237.91  |
| 212.18.104.23  |
| 213.159.68.106 |
| 38.55.136.95   |
| 45.147.200.112 |
| 45.147.201.21  |
| 46.17.41.7     |
| 46.17.44.132   |
| 62.133.60.81   |
| 62.133.63.237  |
| 91.149.218.194 |
| 92.118.112.133 |
| 92.118.112.89  |
+----------------+

authorized_keys


ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCb5Om4AiG92MQtCoOPFtc2AK4oXBgygRZcRN794bmOZyfAjA0VZeXk5lX7O9njMBy2smACj8EfXhNezEXlpCkEy6CbPdfic9ecvEEpuh1LNrFDNkt1hcnK4azErbAK9gOYI1DFU7+EAQWo+72VTsF944JS/BfpfETcczDhky5UyoVU++46OSTIJOmCQBB72TK2yh1NoSyhPHFoK7K5Ky82NQlfXsQs+nLhfAnKvsACCZlzWAw9Z1HtFoV4hwEJKbnD8yMqwbDv9GabmQHUt/GfRHHTkG33ecFpGzWjhvpMCyZk0OmQ9IWStAvc0ogCgGPvKN0lq4cIUb0qkS2qJ+kQFCcFl2ZBIdyPBEfD1LdaJoSo+y4WqOMwREZZhLEowpN7lXzTV7HBhFClHBumF5IHxPjObPDjkfO2yReogFGu5+dBISX6KDsX9eS+rUsJpi0BFvDpuWK2FQ93us4MOkABO4MySskdrNKxTN0LogQ4wFQ0hEIEMe/bZIPv+p9C00U= root@vm857751


*/
