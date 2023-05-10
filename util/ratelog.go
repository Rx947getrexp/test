package util

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"go-speed/config"
	"io"
	"os"
	"time"
)

func GetLogWriter(logConfig config.Log) (out LevelWriter) {
	switch logConfig.Adapter {
	case "file":
		basicPath := logConfig.Path
		allPath := basicPath + ".all"
		allWriter, err := rotatelogs.New(
			allPath+".%Y%m%d",
			rotatelogs.WithLinkName(allPath),
			rotatelogs.WithRotationCount(logConfig.ReverseDays),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			panic(err)
		}
		errPath := basicPath + ".error"
		errWriter, err := rotatelogs.New(
			errPath+".%Y%m%d",
			rotatelogs.WithLinkName(errPath),
			rotatelogs.WithRotationCount(logConfig.ReverseDays),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			panic(err)
		}
		return LevelWriter{
			allWriter: allWriter,
			errWriter: errWriter,
		}
	}
	return LevelWriter{
		allWriter: os.Stdout,
	}
}

type LevelWriter struct {
	allWriter io.Writer
	errWriter io.Writer
}

func (lw LevelWriter) Write(p []byte) (n int, err error) {
	return lw.allWriter.Write(p)
}

func (lw LevelWriter) WriteLevel(l zerolog.Level, p []byte) (n int, err error) {
	// 没有分文件，或者错误等级比较低，写到常规日志中就行了
	if lw.errWriter == nil || l < zerolog.ErrorLevel {
		return lw.allWriter.Write(p)
	}
	// 错误日志，打印到两个文件中
	_, _ = lw.allWriter.Write(p)
	return lw.errWriter.Write(p)
}
