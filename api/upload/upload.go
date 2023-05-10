package upload

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/lang"
	"go-speed/model/request"
	"go-speed/model/response"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PathUpload(c *gin.Context) {
	var uploadImgParam response.PathUploadParam
	param := new(request.PathUpload)
	err := c.ShouldBind(param)
	if err != nil {
		global.Logger.Err(err).Send()
		return
	}

	//只允许白名单上传
	reqIP := c.RemoteIP()
	ipAddrs := global.Viper.GetStringSlice("ip_white_list")
	fmt.Println(reqIP)
	fmt.Println(ipAddrs)
	if len(ipAddrs) != 0 {
		var flag = false
		for _, item := range ipAddrs {
			if reqIP == item {
				flag = true
				break
			}
		}
		if !flag {
			global.Logger.Err(err).Msgf("白名单错误！%s", ipAddrs)
			response.RespFail(c, lang.Translate("cn", "fail"), nil)
			return
		}
	}

	dirName := constant.FilePath + "/" + constant.UploadFilePath + "/" + param.DirName
	dirName = strings.ReplaceAll(dirName, "..", "")
	dirName = filepath.Clean(dirName)
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		// 不存在则创建目录
		err = os.MkdirAll(dirName, 0775)
	}
	if err != nil {
		global.Logger.Err(err).Msg("文件夹创建失败")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	// 图片路径：public/目录名/新文件名
	path := fmt.Sprintf("%s/%s", dirName, param.NewFilename)
	pathName := path + fmt.Sprint(time.Now().UnixNano())
	// 保存图片
	fileBytes, err := base64.StdEncoding.DecodeString(param.FileContent)
	if err != nil {
		global.Logger.Err(err).Send()
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	out, err := os.Create(pathName)
	if err != nil {
		global.Logger.Err(err).Send()
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	_, err = io.Copy(out, bytes.NewReader(fileBytes))
	if err != nil {
		global.Logger.Err(err).Send()
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	out.Close()
	err = os.Rename(pathName, path)
	if err != nil {
		global.Logger.Err(err).Send()
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	path = "/" + constant.FilePath + "/" + constant.UploadFilePath + "/" + param.DirName + "/" + param.NewFilename
	uploadImgParam.Url = path
	response.RespOk(c, lang.Translate("cn", "success"), uploadImgParam)
}
