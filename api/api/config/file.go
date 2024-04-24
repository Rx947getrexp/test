package config

import (
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"io/ioutil"
	"os"
	"strings"
)

func readFileLines(ctx *gin.Context, filePath, country string) (lines []string, err error) {
	var (
		file    *os.File
		content []byte
	)
	lines = make([]string, 0)

	// 检查文件是否存在
	_, err = os.Stat(filePath)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("file Stat failed, filePath: %s", filePath)
		return
	}
	// 文件存在，读取文件内容
	file, err = os.Open(filePath)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("file Open failed, %s", filePath)
		return
	}
	defer file.Close()

	// 读取整个文件内容到一个字符串
	content, err = ioutil.ReadFile(filePath)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("file ReadFile failed, %s", filePath)
		return
	}

	// 使用strings.SplitN按行分割字符串，每次分割后保留一个空白字符
	items := strings.SplitAfter(string(content), "\n")

	// 遍历每一行
	for _, item := range items {
		// 去除每行的首尾空白字符
		cleanLine := strings.TrimSpace(item)
		// 如果处理后的行不为空，则打印
		if cleanLine != "" {
			lines = append(lines, cleanLine)
		}
	}
	return
}
