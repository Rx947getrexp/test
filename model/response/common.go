package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"go-speed/i18n"
	"net/http"
	"strings"
)

type Empty struct {
}

type Response struct {
	Code int             `json:"code"`
	Msg  string          `json:"message"`
	Data json.RawMessage `json:"data"`
}

func (s Response) Print() {
	fmt.Println(fmt.Sprintf(`Response: {"code":%d,"message":"%s","data":"%s"}`, s.Code, s.Msg, string(s.Data)))
}

func (s Response) Dump() string {
	return fmt.Sprintf(`Response: {"code":%d,"message":"%s","data":"%s"}`, s.Code, s.Msg, string(s.Data))
}

type PageRes struct {
	List     interface{} `json:"list"`
	Summary  float64     `json:"summary"`
	Subtotal float64     `json:"subtotal"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Size     int         `json:"size"`
}

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

type ResultList struct {
	List interface{} `json:"list"`
}

const (
	Success = 200
	Fail    = 100

	CodeTokenExpired = 301 // 登录态已失效，请重新登陆
)

func RespOk(c *gin.Context, msg string, data interface{}) {
	dataBytes, _ := json.Marshal(data)
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Msg:  msg,
		Data: dataBytes,
	})
}

func RespFail(c *gin.Context, msg string, data interface{}, code ...int) {
	dataBytes, _ := json.Marshal(data)

	var _code int
	_code, ok := i18n.ReturnCodeMap[msg]
	if !ok {
		_code = 0
	}
	if strings.Contains(msg, "pq") || strings.Contains(msg, "column") {
		msg = "error"
	} else {
		msg = i18n.I18nTrans(c, msg)
	}
	retCode := Fail
	if len(code) > 0 {
		retCode = code[0]
	} else if _code != 0 && _code != 200 {
		retCode = _code
	}
	global.MyLogger(c).Err(fmt.Errorf(">>>>>>>>> code: %d, msg: %s", retCode, msg)).Msgf("response fail msg to client")
	c.JSON(http.StatusOK, Response{
		Code: retCode,
		Msg:  msg,
		Data: dataBytes,
	})
}

func ResOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Msg:  msg,
	})
}

func ResFail(c *gin.Context, msg string, code ...int) {
	if len(code) > 0 {
		c.JSON(http.StatusOK, Response{
			Code: code[0],
			Msg:  msg,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code: Fail,
			Msg:  msg,
		})
	}
}
