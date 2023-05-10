package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go-speed/global"
	"go-speed/lang"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func Upload(param *request.UploadFile, fileExtMap map[string]bool) (string, error) {
	var result string
	fileHeader := param.Files
	// 文件扩展名
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == "" {
		return result, errors.New(lang.Translate("cn", "file ext error"))
	}
	if _, ok := fileExtMap[fileExt]; !ok {
		return result, errors.New(lang.Translate("cn", "file ext error"))
	}
	file, err := fileHeader.Open()
	if err != nil {
		return result, errors.New(lang.Translate("cn", "fail"))
	}
	defer file.Close()
	var fileBytes bytes.Buffer
	_, err = io.Copy(&fileBytes, file)
	if err != nil {
		return result, errors.New(lang.Translate("cn", "fail"))
	}
	fileExt = strings.ToLower(fileExt)
	// 时间戳+文件名，md5生成uuid
	newFilename := util.MD5(fmt.Sprintf("%d%s", time.Now().Unix(), fileHeader.Filename)) + fileExt
	//文件扩展名
	switch fileExt {
	case ".jpg2", ".jpeg2":
		_, err := jpeg.Decode(bytes.NewReader(fileBytes.Bytes()))
		if err != nil {
			global.Logger.Err(err).Msg("图片格式验证失败")
			return result, errors.New(lang.Translate("cn", "fail"))
		}
	case ".png2":
		_, err := png.Decode(bytes.NewReader(fileBytes.Bytes()))
		if err != nil {
			global.Logger.Err(err).Msg("图片格式验证失败")
			return result, errors.New(lang.Translate("cn", "fail"))
		}
	case ".jpg", ".jpeg", ".png", ".pdf", ".mp4", ".avi", ".dat", ".mkv", ".flv", ".vob", ".mp3", ".wav", ".wma", ".mp2", ".ra", ".ape", ".aac", ".cda", ".mov", ".gif", ".ppt", ".pptx":
		// todo 增加验证
		fileName := strings.TrimSuffix(fileHeader.Filename, fileExt)
		fmt.Println(fileName)
		//newFilename = fileName + fileExt
	default:
		global.Logger.Error().Msgf("文件扩展名错误")
		return result, errors.New(lang.Translate("cn", "file ext error"))
	}

	data := new(response.DataResponse)
	hotAddr := global.Viper.GetString("system.file_link")
	fileServiceUrl := hotAddr + "/upload"
	var pathUpload request.PathUpload
	pathUpload.FileContent = base64.StdEncoding.EncodeToString(fileBytes.Bytes())
	pathUpload.DirName = param.FileType
	pathUpload.NewFilename = newFilename
	buildParam, _ := json.Marshal(pathUpload)
	err = sendRequest("POST", fileServiceUrl, bytes.NewReader(buildParam), data)
	if err != nil {
		global.Logger.Err(err).Msgf("post请求接口错误")
		return result, errors.New(lang.Translate("cn", "fail"))
	}
	if data.Code != response.Success {
		global.Logger.Err(err).Msgf("upload code错误")
		return result, errors.New(lang.Translate("cn", "fail"))
	}
	result = data.Data.Url
	return result, nil
}

func sendRequest(method, url string, body io.Reader, response interface{}) error {
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		global.Logger.Err(err).Msg("创建http请求失败")
		return err
	}
	if method == "POST" {
		//httpReq.Header.Add("Content-Type", "multipart/form-data")
		httpReq.Header.Add("Content-Type", "application/json")
	}
	httpReq.Close = true
	respData, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		global.Logger.Err(err).Msg("发送http请求失败")
		return err
	}
	respBytes, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		global.Logger.Err(err).Msg("读取响应数据失败")
		return err
	}
	defer respData.Body.Close()
	if err = json.Unmarshal(respBytes, response); err != nil {
		global.Logger.Err(err).Msg("解析数据")
		return err
	}
	return nil
}
