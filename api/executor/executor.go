package executor

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var v2rayJson = "{\n\t\"inbounds\": [{\n\t\t\"tag\": \"tcp-ws\",\n\t\t\"port\": 11111,\n\t\t\"listen\": \"127.0.0.1\",\n\t\t\"protocol\": \"vmess\",\n\t\t\"settings\": {\n\t\t\t\"clients\": [{\n\t\t\t\t\t\"email\": \"###\",\n\t\t\t\t\t\"id\": \"***\",\n\t\t\t\t\t\"alterId\": 0,\n\t\t\t\t\t\"level\": 0\n\t\t\t\t}\n\n\t\t\t]\n\t\t},\n\t\t\"streamSettings\": {\n\t\t\t\"network\": \"ws\",\n\t\t\t\"wsSettings\": {\n\t\t\t\t\"path\": \"/work\"\n\t\t\t}\n\t\t}\n\t}]\n\n}\n"

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
	path := ""
	if param.Tag == "1" {
		path = fmt.Sprintf("/v2rayJsonAdd/%s.json", param.Uuid)
	} else {
		path = fmt.Sprintf("/v2rayJsonSub/%s.json", param.Uuid)
	}
	v2rayJson = strings.ReplaceAll(v2rayJson, "###", param.Email)
	v2rayJson = strings.ReplaceAll(v2rayJson, "***", param.Uuid)
	fmt.Println(v2rayJson)
	fmt.Printf("111TTTTTTThistest, %s", param.Email)
	fmt.Printf("222TTTTTTThistest, %s", param.Uuid)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		global.Logger.Err(err).Msg("添加失败")
		response.ResFail(c, "添加失败")
		return
	}
	defer file.Close()
	wr := bufio.NewWriter(file)
	_, err = wr.WriteString(v2rayJson) //注意这里是写在缓存中的，而不是直接落盘的
	if err != nil {
		global.Logger.Err(err).Msg("添加失败")
		response.ResFail(c, "添加失败")
		return
	}
	wr.Flush() //将缓存的内容写入文件

	if param.Tag == "1" {
		_ = os.Remove(fmt.Sprintf("/v2rayJsonSub/%s.json", param.Uuid))
		cmds := exec.Command("/usr/local/bin/v2ray", "  api adi -s 127.0.0.1:10085 /v2rayJsonAdd")
		err = cmds.Start()
		if err != nil {
			global.Logger.Err(err).Msg("添加失败")
			response.ResFail(c, "添加udid启动失败")
			return
		}

	} else {
		_ = os.Remove(fmt.Sprintf("/v2rayJsonAdd/%s.json", param.Uuid))
		cmds := exec.Command("/usr/local/bin/v2ray", "  api rmi -s 127.0.0.1:10085 /v2rayJsonSub")
		err = cmds.Start()
		if err != nil {
			global.Logger.Err(err).Msg("删除udid启动失败")
			response.ResFail(c, "删除udid启动失败")
			return
		}
	}
	global.Logger.Info().Msg("添加成功")
	_ = os.Remove(path)
	response.ResOk(c, "成功")
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
