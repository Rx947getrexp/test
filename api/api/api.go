package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go-speed/global"
	"go-speed/lang"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"net/http"
	"time"
)

// GenerateDevId C端获取DEV_ID，并保存在本地全局存储
func GenerateDevId(c *gin.Context) {
	id, _ := service.GenSnowflake()
	userAgent := c.GetHeader("User-Agent")
	ua := user_agent.New(userAgent)
	dev := &model.TDev{
		Id:        id,
		Os:        ua.OS(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err := global.Db.Insert(dev)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("db连接出错")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	result := make(map[string]interface{})
	result["dev_id"] = id
	response.RespOk(c, "成功", result)
}

func SendEmail(c *gin.Context) {
	param := new(request.SendEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	user := new(model.TUser)
	has, err := global.Db.Where("uname = ?", param.Email).Get(user)
	if err != nil {
		global.Logger.Err(err).Msg("db连接出错")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if !has {
		global.Logger.Err(err).Msg("邮箱地址未注册")
		response.RespFail(c, "邮箱地址未注册", nil)
		return
	}
	clientIp := c.ClientIP()
	err = service.SendTelSms(param.Email, clientIp)
	if err != nil {
		global.Logger.Err(err).Msg("发送短信失败！")
		response.ResFail(c, "发送验证码失败,请稍后再试！")
		return
	}
	response.ResOk(c, "发送成功")
}

func Reg(c *gin.Context) {
	param := new(request.RegRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if param.Account == "" || param.Passwd == "" || param.EnterPasswd == "" {
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}

	id, _ := service.GenSnowflake()
	uuid := util.GetUuid()
	pwdDecode := util.AesDecrypt(param.Passwd)

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	sess.Begin()

	user := &model.TUser{
		Id:          id,
		Uname:       param.Account,
		Passwd:      util.MD5(pwdDecode),
		Email:       param.Account,
		Phone:       "",
		Level:       0,
		ExpiredTime: 0,
		V2rayUuid:   uuid,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Comment:     "",
	}
	rows, err := sess.Insert(user)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("添加user出错")
		sess.Rollback()
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}

	sess.Commit()
	response.ResOk(c, "注册成功")
}

func Login(c *gin.Context) {
	param := new(request.LoginRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if param.Account == "" || param.Passwd == "" {
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)
	user := new(model.TUser)
	has, err := global.Db.Where("uname = ? and passwd = ? and status = 0", param.Account, pwdMd5).Get(user)
	if err != nil {
		global.Logger.Err(err).Msgf("登录出错！%s", param.Account)
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
	}
	if !has {
		global.Logger.Err(err).Msgf("用户名或密码不正确！%s", param.Account)
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
	}
	dataParam := response.LoginClientParam{
		UserId: user.Id,
		Token:  service.GenerateTokenByUser(user.Id, service.CommonUserType),
	}
	response.RespOk(c, "登录成功", dataParam)
}

func ChangePasswd(c *gin.Context) {
	param := new(request.ChangePasswdRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("修改密码失败")
		response.ResFail(c, "修改密码失败！")
		return
	}
	oldPwdDecode := util.AesDecrypt(param.OldPasswd)
	oldPwdMd5 := util.MD5(oldPwdDecode)
	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)

	user.Passwd = pwdMd5
	user.UpdatedAt = time.Now()
	rows, err := global.Db.Cols("passwd", "updated_at").Where("id = ? and passwd = ?", user.Id, oldPwdMd5).Update(user)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("修改密码失败")
		response.ResFail(c, "修改密码失败！")
		return
	}
	response.ResOk(c, "成功")
}

func ForgetPasswd(c *gin.Context) {
	param := new(request.ForgetRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	err := service.VerifyMsg(param.Account, param.VerifyCode)
	if err != nil {
		global.Logger.Err(err).Msg("验证码错误")
		response.ResFail(c, "验证码错误！")
		return
	}
	if param.Passwd != param.EnterPasswd {
		response.ResFail(c, "两次密码不一致")
		return
	}
	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)
	user := new(model.TUser)
	user.Passwd = pwdMd5
	user.UpdatedAt = time.Now()
	rows, err := global.Db.Cols("passwd", "updated_at").Where("uname = ? ", param.Account).Update(user)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("修改密码失败")
		response.ResFail(c, "修改密码失败！")
		return
	}
	response.ResOk(c, "成功")
}

// JWTAuth 验证token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization-Token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": "授权已过期",
			})
			c.Abort()
			return
		}

		err = service.AddLog(c, claims.UserId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    100,
				"message": "网络错误",
			})
			c.Abort()
			return
		}
		//isBool := RedisInsert(strconv.FormatInt(claims.UserId, 10))
		//if !isBool {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 301,
		//		"msg":  "授权已过期",
		//	})
		//	c.Abort()
		//	return
		//}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		//uu := c.MustGet("claims").(*service.CustomClaims)
		//fmt.Println("claims...", uu.UserId)
	}
}
