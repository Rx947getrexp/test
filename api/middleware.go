package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go-speed/global"
	"go-speed/model/response"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type resourceHosts struct {
	File    string `mapstructure:"file"`
	Backend string `mapstructure:"backend"`
	Website string `mapstructure:"website"`
}

var (
	resourceHostMap map[string]resourceHosts
	once            sync.Once
)

// LoadResourceHost 根据客户端使用的域名，返回相应的资源域名
func LoadResourceHost(c *gin.Context) {
	once.Do(func() {
		err := global.Viper.UnmarshalKey("file-host", &resourceHostMap)
		if err != nil {
			global.Logger.Error().Msgf("加载配置文件失败")
			return
		}
	})
	clientDomain := c.GetHeader("XXX-DOMAIN")
	global.Logger.Debug().Msgf("clientDomain=%s,resourceHostMap=%v", clientDomain, resourceHostMap)
	for clientHost, hosts := range resourceHostMap {
		if strings.Contains(clientDomain, clientHost) {
			c.Set("h-file", hosts.File)
			c.Set("h-backend", hosts.Backend)
			c.Set("h-website", hosts.Website)
			return
		}
	}
}

// PrintRequest 打印请求内容
func PrintRequest(c *gin.Context) {
	if global.Logger.GetLevel() > zerolog.DebugLevel {
		c.Next()
		return
	}
	bodyBytes := ReadBodyToCache(c)
	global.Logger.Debug().Msgf("request={%s}, data={%v}", c.Request.RequestURI, string(bodyBytes))
}

func ReadBodyToCache(c *gin.Context) []byte {
	if bodeBytes, ok := c.Get("body-bytes"); ok {
		return bodeBytes.([]byte)
	}
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	c.Set("body-bytes", bodyBytes)
	return bodyBytes
}

var whitePath = map[string]bool{
	"/created_article": true,
	"/update_article":  true,
}

func FilteredSQLInject(c *gin.Context) {
	if whitePath[c.Request.URL.Path] {
		c.Next()
		return
	}
	queryForm, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		global.Logger.Err(err).Msg("参数解析失败")
		response.RespFail(c, "param err", nil)
		c.Abort()
		return
	}
	for _, params := range queryForm {
		for _, param := range params {
			if !isLegalParam(param) {
				global.Logger.Err(err).Msgf("参数%s包含非法字符串", param)
				response.RespFail(c, "param err", nil)
				c.Abort()
				return
			}
		}
	}
	headerParam := c.Request.Header
	for _, value := range headerParam {
		for _, v := range value {
			if !isLegalParam(v) {
				global.Logger.Err(err).Msgf("参数%s包含非法字符串", v)
				response.RespFail(c, "param err", nil)
				c.Abort()
				return
			}
		}
	}
	bodyBytes := ReadBodyToCache(c)
	switch c.ContentType() {
	case "application/json":
		data := make(map[string]interface{})
		_ = json.Unmarshal(bodyBytes, &data)
		if !validMap(data) {
			global.Logger.Err(err).Msgf("参数%v包含非法字符串", data)
			response.RespFail(c, "param err", nil)
			c.Abort()
			return
		}
	default:
		// 触发gin内部把参数绑定到PostForm里面
		_, _ = c.GetPostFormMap("")
		for _, params := range c.Request.PostForm {
			for _, param := range params {
				if !isLegalParam(param) {
					global.Logger.Err(err).Msgf("参数%s包含非法字符串", param)
					response.RespFail(c, "param err", nil)
					c.Abort()
					return
				}
			}
		}
	}
	c.Next()
}

func validMap(data map[string]interface{}) bool {
	for _, value := range data {
		switch value.(type) {
		case map[string]interface{}:
			legal := validMap(value.(map[string]interface{}))
			if !legal {
				return false
			}
		case string:
			legal := isLegalParam(value.(string))
			if !legal {
				return false
			}
		}
	}
	return true
}

// 非法参数集合
var _illegalSlice = []string{
	"select",
	"update",
	"delete",
	"insert",
	"truncate",
	"declare",
	"exec",
	"drop",
	"execute",
}

// 判断参数是否合法
func isLegalParam(param string) bool {
	if param == "" {
		return true
	}
	param = strings.ToLower(param)
	for _, key := range _illegalSlice {
		if strings.Contains(param, key) {
			return false
		}
	}
	return true
}

// RateMiddleware 限制访问
func RateMiddleware(limiter *Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果ip请求连接数在两秒内超过5次，返回429并抛出error
		if !limiter.Allow(c.ClientIP(), 5, 2*time.Second) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    100,
				"message": "too many requests",
			})
			log.Println("too many requests")
			return
		}
		c.Next()
	}
}
