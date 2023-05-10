package log_adaptor

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"time"
	"xorm.io/xorm/log"
)

type XormLogger struct {
	showSQL bool
	zlog    zerolog.Logger
}

func NewXormLogger(w io.Writer) *XormLogger {
	innerLogger := zerolog.New(w).With().Timestamp().Logger()
	return &XormLogger{zlog: innerLogger}
}

func (x *XormLogger) BeforeSQL(log.LogContext) {
}

func (x *XormLogger) AfterSQL(ctx log.LogContext) {
	var sessionID string
	v := ctx.Ctx.Value(log.SessionIDKey)
	if key, ok := v.(string); ok {
		sessionID = fmt.Sprintf(" [%s]", key)
	}
	var event *zerolog.Event
	if ctx.ExecuteTime >= time.Second {
		// 大于1s打印警告信息
		event = x.zlog.Warn()
	} else {
		event = x.zlog.Info()
	}
	msg := fmt.Sprintf("执行sql：sessionID：%s, sql：%s, args：%v, executeTime：%v",
		sessionID, ctx.SQL, ctx.Args, ctx.ExecuteTime)
	event.Msg(msg)
}

func (x *XormLogger) Debugf(format string, v ...interface{}) {
	x.zlog.Debug().Msgf(format, v)
}

func (x *XormLogger) Errorf(format string, v ...interface{}) {
	x.zlog.Error().Msgf(format, v)
}

func (x *XormLogger) Infof(format string, v ...interface{}) {
	x.zlog.Info().Msgf(format, v)
}

func (x *XormLogger) Warnf(format string, v ...interface{}) {
	x.zlog.Warn().Msgf(format, v)
}

func (x *XormLogger) Level() log.LogLevel {
	zeroLevel := x.zlog.GetLevel()
	switch zeroLevel {
	case zerolog.DebugLevel:
		return log.LOG_DEBUG
	case zerolog.InfoLevel:
		return log.LOG_INFO
	case zerolog.WarnLevel:
		return log.LOG_WARNING
	case zerolog.ErrorLevel:
		return log.LOG_ERR
	case zerolog.Disabled:
		return log.LOG_OFF
	}
	return log.LOG_UNKNOWN
}

func (x *XormLogger) SetLevel2(levelStr string) {
	level, _ := zerolog.ParseLevel(levelStr)
	x.zlog = x.zlog.Level(level)
}

func (x *XormLogger) SetLevel(l log.LogLevel) {
	var zeroLevel zerolog.Level
	switch l {
	case log.LOG_DEBUG:
		zeroLevel = zerolog.DebugLevel
	case log.LOG_INFO:
		zeroLevel = zerolog.InfoLevel
	case log.LOG_WARNING:
		zeroLevel = zerolog.WarnLevel
	case log.LOG_ERR:
		zeroLevel = zerolog.ErrorLevel
	case log.LOG_OFF:
		fallthrough
	case log.LOG_UNKNOWN:
		zeroLevel = zerolog.Disabled
	}
	x.zlog = x.zlog.Level(zeroLevel)
}

func (x *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		x.showSQL = true
		return
	}
	x.showSQL = show[0]
}

func (x *XormLogger) IsShowSQL() bool {
	return x.showSQL
}
