package executor

import (
	"bufio"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// TODO: 简单实现添加、删除账号时的串行处理，防止并发执行时的相互覆盖问题，但用锁会影响性能，后续成为瓶颈时需要优化处理。
var addSubMutex sync.Mutex

var v2rayJson = ""

//"{\n\t\"inbounds\": [{\n\t\t\"tag\": \"tcp-ws\",\n\t\t\"port\": 11111,\n\t\t\"listen\": \"127.0.0.1\",\n\t\t\"protocol\": \"vmess\",\n\t\t\"settings\": {\n\t\t\t\"clients\": [{\n\t\t\t\t\t\"email\": \"%s\",\n\t\t\t\t\t\"id\": \"%s\",\n\t\t\t\t\t\"alterId\": 0,\n\t\t\t\t\t\"level\": 0\n\t\t\t\t}\n\n\t\t\t]\n\t\t},\n\t\t\"streamSettings\": {\n\t\t\t\"network\": \"ws\",\n\t\t\t\"wsSettings\": {\n\t\t\t\t\"path\": \"/work\"\n\t\t\t}\n\t\t}\n\t}]\n\n}\n"

func NodeAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("accessToken")
		timestamp := c.GetHeader("timestamp")
		md5Str := util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
		if accessToken != md5Str {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "token鉴权失败，无权限访问",
			})
			c.Abort()
			return
		}
	}
}

func AddSub(c *gin.Context) {
	// 加锁处理，防止配置文件并发请求处理时相互覆盖.
	addSubMutex.Lock()
	defer addSubMutex.Unlock()

	param := new(request.NodeAddSubRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	global.Logger.Info().Msgf(">>>>>>>> Tag: %s, Email: %s, Uuid: %s, Level: %s", param.Tag, param.Email, param.Uuid, param.Level)
	if param.Tag == "1" { // TODO： Tag定义是啥？
		err := addUser(param)
		if err != nil {
			global.Logger.Err(err).Msg("添加失败")
			response.ResFail(c, "添加失败, "+err.Error())
			return
		}
	} else {
		err := delUser(param)
		if err != nil {
			global.Logger.Err(err).Msg("删除失败")
			response.ResFail(c, "删除失败, "+err.Error())
			return
		}
	}
	response.ResOk(c, "成功")
	return
}

func AddSubBak(c *gin.Context) {
	param := new(request.NodeAddSubRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	name := param.Uuid
	path := ""
	if param.Tag == "1" {
		v2rayJson = fmt.Sprintf("{\"email\": \"%s\",\"password\": \"%s\"}", param.Email, param.Uuid)
		path = fmt.Sprintf("/v2rayJsonAdd/%s.json", name)
		_, err := os.Stat(path)

		//isnotexist来判断，是不是不存在的错误
		if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
			fmt.Printf(path + " is not exists")
		} else {
			fmt.Printf(path + " is exists")
			response.ResOk(c, "成功")
			return
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		wr := bufio.NewWriter(file)
		fmt.Printf(v2rayJson)
		_, err = wr.WriteString(v2rayJson) //注意这里是写在缓存中的，而不是直接落盘的
		wr.Flush()                         //将缓存的内容写入文件
		defer file.Close()
		if err != nil {
			global.Logger.Err(err).Msg("添加失败")
			response.ResFail(c, "添加失败")
			return
		}
	} else {
		path = fmt.Sprintf("/v2rayJsonAdd/%s.json", name)
		_ = os.Remove(path)
	}
	errs := Command("/usr/bin/python3 /shell/addAccount.py")
	if errs != nil {
		response.ResFail(c, "添加失败")
		return
	}
	/*
		fmt.Printf(v2rayJson)
		fmt.Printf("---------------------")
		dnsList, _ := service.FindExpireUsers()
		fmt.Printf(v2rayJson)
		for _, item := range dnsList {
			_ = os.Remove(fmt.Sprintf("/v2rayJsonAdd/%s", item.V2rayUuid))
		}*/
	response.ResOk(c, "成功")
	return
}

func AddSub2(c *gin.Context) {
	param := new(request.NodeAddSubRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	response.ResOk(c, "成功")

	path := ""
	name := param.Uuid
	email := param.Email
	path2 := fmt.Sprintf("/v2rayJsonAddBak/%s.json", name)
	if param.Tag == "1" {
		path = fmt.Sprintf("/v2rayJsonAdd/%s.json", name)
	} else {
		path = fmt.Sprintf("/v2rayJsonSub/%s.json", name)

	}

	//v2rayJson = strings.ReplaceAll(v2rayJson, "###", email)
	//v2rayJson = strings.ReplaceAll(v2rayJson, "***", name)

	v2rayJson = fmt.Sprintf("{\"inbounds\":[{\"tag\":\"tcp-ws\",\"port\":10085,\"listen\":\"127.0.0.1\",\"protocol\":\"trojan\",\"settings\":{\"clients\":[{\"email\":\"%s\",\"password\":\"%s\"}],\"fallbacks\":[{\"dest\":80}]},\"streamSettings\":{\"network\":\"ws\",\"security\":\"none\",\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"certificates\":[{\"certificateFile\":\"/usr/local/cert/cert.crt\",\"keyFile\":\"/usr/local/cert/private.key\"}]},\"wsSettings\":{\"path\":\"/work\",\"headers\": {}}}}]}", email, name)
	fmt.Printf("111TTTTTTThistest, Email:%s,uuid:%s,Tag:%s", email, name, param.Tag)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		global.Logger.Err(err).Msg("添加失败")
		response.ResFail(c, "添加失败")
		return
	}
	defer file.Close()
	wr := bufio.NewWriter(file)
	_, err = wr.WriteString(v2rayJson) //注意这里是写在缓存中的，而不是直接落盘的
	fmt.Printf(v2rayJson)
	if err != nil {
		global.Logger.Err(err).Msg("添加失败")
		response.ResFail(c, "添加失败")
		return
	}
	wr.Flush() //将缓存的内容写入文件

	if param.Tag == "1" {
		_ = os.Remove(fmt.Sprintf("/v2rayJsonSub/%s.json", param.Uuid))

		//cmds := exec.Command("/usr/local/bin/v2ray", "  api adi -s 127.0.0.1:10085 /v2rayJsonAdd")
		//err = cmds.Start()
		err := Command("/usr/local/bin/v2ray api adi -s 127.0.0.1:10088 /v2rayJsonAdd")
		err1 := Command("ps -e|grep v2ray|awk '{print $1}'|xargs kill -9")
		err2 := Command("nohup /usr/local/bin/v2ray run -c /usr/local/etc/v2ray/config.json > /dev/null 2>&1 &")

		if err != nil || err1 != nil || err2 != nil {
			global.Logger.Err(err).Msg("添加失败")
			response.ResFail(c, "添加失败")
			return
		}
		global.Logger.Info().Msg("添加成功")
		file2, errx := os.OpenFile(path2, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		wrx := bufio.NewWriter(file2)
		_, err = wrx.WriteString(v2rayJson) //注意这里是写在缓存中的，而不是直接落盘的
		defer file2.Close()
		if errx != nil {

		}
	} else {
		_ = os.Remove(fmt.Sprintf("/v2rayJsonAdd/%s.json", param.Uuid))
		err := Command("/usr/local/bin/v2ray api rmi -s 127.0.0.1:10088 /v2rayJsonSub")
		_ = os.Remove(path2)
		err1 := Command("ps -e|grep v2ray|awk '{print $1}'|xargs kill -9")
		err2 := Command("nohup /usr/local/bin/v2ray run -c /usr/local/etc/v2ray/config.json > /dev/null 2>&1 &")

		if err != nil || err1 != nil || err2 != nil {
			global.Logger.Err(err).Msg("删除udid启动失败")
			response.ResFail(c, "操作失败")
			return
		}
		global.Logger.Info().Msg(" /usr/local/bin/v2ray api rmi -s 127.0.0.1:10088 /v2rayJsonSub 删除成功")
		response.ResFail(c, "已到期，请充值")
	}

	_ = os.Remove(path)
	response.ResOk(c, "成功")
}
func Command(cmd string) error {
	//c := exec.Command("cmd", "/C", cmd)   // windows
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			byte2String := ConvertByte2String([]byte(readString))
			fmt.Print(byte2String)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}
func ConvertByte2String(byte []byte) string {
	var str string

	str = string(byte)

	return str
}
func AddEmail(c *gin.Context) {
	param := new(request.NodeAddEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	response.ResOk(c, "成功")
}

func RemoveEmail(c *gin.Context) {
	param := new(request.NodeRemoveEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	response.ResOk(c, "成功")
}

func GetUserList(c *gin.Context) {
	conf, err := ReadV2rayConfig(global.Config.System.V2rayConfigPath)
	if err != nil {
		global.Logger.Err(err).Msg("read v2ray config failed, err: " + err.Error())
		response.ResFail(c, "read v2ray config failed, "+err.Error())
		return
	}
	var items []response.ClientItem
	for _, client := range conf.GetClients() {
		items = append(items, response.ClientItem{
			Email:    client.Email,
			Password: client.Password,
		})
	}
	resp := response.GetUserListResponse{Items: items}
	response.RespOk(c, i18n.RetMsgSuccess, resp)
	return
}

func GetUserTraffic(c *gin.Context) {
	param := new(request.GetUserTrafficRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	global.Logger.Info().Msgf(">>>>>>>> param: %+v", *param)
	clientEmails := param.Emails
	if len(param.Emails) == 0 && param.All {
		conf, err := ReadV2rayConfig(global.Config.System.V2rayConfigPath)
		if err != nil {
			global.Logger.Err(err).Msg("read v2ray config failed, err: " + err.Error())
			response.ResFail(c, "read v2ray config failed, "+err.Error())
			return
		}
		for _, client := range conf.GetClients() {
			clientEmails = append(clientEmails, client.Email)
		}
	}
	global.Logger.Info().Msgf(">>>>>>>> len(clientEmails): %+v", len(clientEmails))
	resp := response.GetUserTrafficResponse{}
	if len(clientEmails) == 0 {
		response.RespOk(c, i18n.RetMsgSuccess, resp)
		return
	}

	var patterns []string
	for _, email := range clientEmails {
		patterns = append(patterns, fmt.Sprintf("user>>>%s>>>traffic>>>uplink", email))
		patterns = append(patterns, fmt.Sprintf("user>>>%s>>>traffic>>>downlink", email))
	}
	stats, err := QueryUserStats(nil, patterns, param.Reset)
	if err != nil {
		response.ResFail(c, "查询用户Traffic失败, "+err.Error())
		return
	}

	userTrafficMap := make(map[string]*response.UserTrafficItem)
	for _, stat := range stats {
		splits := strings.Split(stat.Name, ">>>")
		if len(splits) != 4 {
			global.Logger.Warn().Msgf("traffic data invalid, %s : %d", stat.Name, stat.Value)
			continue
		}
		var (
			prefix      string
			tag         string
			email       string
			trafficType string
			ok          bool
			userTraffic *response.UserTrafficItem
		)
		prefix = splits[0]
		email = splits[1]
		tag = splits[2]
		trafficType = splits[3]
		if prefix != "user" || tag != "traffic" {
			global.Logger.Warn().Msgf("traffic data invalid, %s : %d", stat.Name, stat.Value)
			continue
		}

		userTraffic, ok = userTrafficMap[email]
		if !ok {
			userTraffic = &response.UserTrafficItem{Email: email}
			userTrafficMap[email] = userTraffic
		}
		if trafficType == "uplink" {
			userTraffic.UpLink = uint64(stat.Value)
		} else if trafficType == "downlink" {
			userTraffic.DownLink = uint64(stat.Value)
		}
	}

	var items []response.UserTrafficItem
	for key, _ := range userTrafficMap {
		items = append(items, *userTrafficMap[key])
	}

	resp = response.GetUserTrafficResponse{Items: items}
	response.RespOk(c, i18n.RetMsgSuccess, resp)
	return
}

func GetV2raySysStats(c *gin.Context) {
	stats, err := service.GetSysStats(c)
	if err != nil {
		response.ResFail(c, "GetSysStats failed, "+err.Error())
		return
	}
	response.RespOk(c, i18n.RetMsgSuccess, response.GetV2raySysStatsResponse{
		NumGoroutine: stats.NumGoroutine,
		NumGC:        stats.NumGC,
		Alloc:        stats.Alloc,
		TotalAlloc:   stats.TotalAlloc,
		Sys:          stats.Sys,
		Mallocs:      stats.Mallocs,
		Frees:        stats.Frees,
		LiveObjects:  stats.LiveObjects,
		PauseTotalNs: stats.PauseTotalNs,
		Uptime:       stats.Uptime,
	})
	return
}
