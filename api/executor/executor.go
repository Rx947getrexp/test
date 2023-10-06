package executor

import (
	"bufio"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
)

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
