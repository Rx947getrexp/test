package global

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-speed/constant"
	"runtime/debug"
	"strings"
	"time"
)

func GetLang(c *gin.Context) string {
	return c.GetHeader(constant.HeaderKeyLang)
}

func GetLangKV(c *gin.Context) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, constant.HeaderKeyLang, c.GetHeader(constant.HeaderKeyLang))
}

func GetClientId(c *gin.Context) string {
	return strings.TrimSpace(c.GetHeader(constant.HeaderKeyClientId))
}

func GetAppVersion(c *gin.Context) string {
	return strings.TrimSpace(c.GetHeader(constant.HeaderKeyAppVersion))
}

func GetClientIdKV(c *gin.Context) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, constant.HeaderKeyClientId, c.GetHeader(constant.HeaderKeyClientId))
}

func GetClientVersion(c *gin.Context) string {
	return strings.TrimSpace(c.GetHeader(constant.HeaderKeyClientVersion))
}

func GetUserAgent(c *gin.Context) string {
	return c.GetHeader(constant.HeaderKeyUserAgent)
}

func GetUserAgentKV(c *gin.Context) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, constant.HeaderKeyUserAgent, c.GetHeader(constant.HeaderKeyUserAgent))
}

func GetChannel(c *gin.Context) string {
	return c.GetHeader(constant.HeaderKeyChannel)
}

func GetChannelKV(c *gin.Context) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, constant.HeaderKeyChannel, c.GetHeader(constant.HeaderKeyChannel))
}

func GetClaims(c *gin.Context) string {
	return c.GetHeader(constant.HeaderKeyClaims)
}

func GetClaimsKV(c *gin.Context) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, constant.HeaderKeyClaims, c.GetHeader(constant.HeaderKeyClaims))
}

func GetTraceId(c *gin.Context) string {
	return c.GetHeader(constant.HeaderKeyTraceId)
}

func GetTraceIdKV(c *gin.Context) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, constant.HeaderKeyTraceId, c.GetHeader(constant.HeaderKeyTraceId))
}

func GetHeaderKV(c *gin.Context, key string) string {
	return fmt.Sprintf(`{"%s":"%s"}"`, key, c.GetHeader(key))
}

func SprintAllHeader(c *gin.Context) string {
	var items []string
	for _, key := range constant.HeaderKeys {
		items = append(items, GetHeaderKV(c, key))
	}
	items = append(items, fmt.Sprintf(`{"URL":"%s"}"`, c.Request.URL.String()))
	items = append(items, fmt.Sprintf(`{"Method":"%s"}"`, c.Request.Method))
	items = append(items, fmt.Sprintf(`{"ClientIP":"%s"}"`, c.ClientIP()))
	return fmt.Sprintf("[%s]", strings.Join(items, ","))
}

func PrintAllHeader(c *gin.Context, err ...error) {
	if err != nil {
		Logger.Err(err[0]).Msgf(SprintAllHeader(c))
	} else {
		Logger.Info().Msgf(SprintAllHeader(c))
	}
}

func TraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 插入请求开始时间
		defer func() {
			if r := recover(); r != nil {
				// 同时打印到日志文件和标准输出中
				MyLogger(c).Err(errors.New("panic-error-dump")).Msgf("%v\n\n%v", r, string(debug.Stack()))
			}
		}()

		startTime := time.Now()
		traceId := c.GetHeader(constant.HeaderKeyTraceId)
		if traceId == "" {
			_traceId := uuid.New().String()
			c.Set(LoggerTraceIdKey, _traceId)
			MyLogger(c).Warn().Msgf("____START____ <Trace-Id: empty> <newTraceId: %s> <Headers: %s>", _traceId, SprintAllHeader(c))
			c.Header(constant.HeaderKeyTraceId, _traceId)
		} else {
			c.Set(LoggerTraceIdKey, traceId)
			MyLogger(c).Info().Msgf("____START____ <Trace-Id: %s> <Headers: %s>", GetTraceId(c), SprintAllHeader(c))
			c.Header(constant.HeaderKeyTraceId, traceId)
		}

		c.Next()
		endTime := time.Now()
		MyLogger(c).Info().Msgf("____END____ <ClientIP: %s> <API_URL: %+v> <Method: %+v> <start-time: %+v> <end-time: %+v> <耗时：%d 毫秒>",
			c.ClientIP(), c.Request.URL.String(), c.Request.Method, startTime, endTime, endTime.Sub(startTime).Milliseconds())
	}
}
