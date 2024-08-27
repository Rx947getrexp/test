package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-speed/api"
	"go-speed/api/api/common"
	v2rayConfig "go-speed/api/api/config"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"go-speed/util/geo"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/user_agent"
	"github.com/shopspring/decimal"
	"xorm.io/xorm"
)

var gConnectNum int

var configs = "{\"log\":{\"level\":\"{{logLevel}}\",\"output\":\"{{leafLogFile}}\"},\"dns\":{\"servers\":[\"1.1.1.1\",\"8.8.8.8\"],\"hosts\":{\"node2.wuwuwu360.xyz\":[\"107.148.239.239\"]}},\"inbounds\":[{\"protocol\":\"tun\",\"settings\":{\"fd\":\"{{tunFd}}\"},\"tag\":\"tun_in\"}],\"outbounds\":[{\"protocol\":\"failover\",\"tag\":\"failover_out\",\"settings\":{\"actors\":[\"proxy1\",\"proxy2\"],\"failTimeout\":4,\"healthCheck\":true,\"checkInterval\":300,\"failover\":true,\"fallbackCache\":false,\"cacheSize\":256,\"cacheTimeout\":60}},{\"tag\":\"proxy1\",\"protocol\":\"chain\",\"settings\":{\"actors\":[\"tls\",\"ws\",\"trojan\"]}},{\"tag\":\"proxy2\",\"protocol\":\"chain\",\"settings\":{\"actors\":[\"tls\",\"ws\",\"trojan\"]}},{\"protocol\":\"tls\",\"tag\":\"tls\",\"settings\":{\"alpn\":[\"http/1.1\"],\"insecure\":true}},{\"protocol\":\"ws\",\"tag\":\"ws\",\"settings\":{\"path\":\"/work\"}},%s,{\"protocol\":\"direct\",\"tag\":\"direct_out\"},{\"protocol\":\"drop\",\"tag\":\"reject_out\"}],\"router\":{\"domainResolve\":true,\"rules\":[{\"external\":[\"site:{{dlcFile}}:cn\"],\"target\":\"direct_out\"},{\"external\":[\"mmdb:{{geoFile}}:cn\"],\"target\":\"direct_out\"},{\"domainKeyword\":[\"apple\",\"icloud\"],\"target\":\"direct_out\"}]}}"

// GenerateDevId C端获取DEV_ID，并保存在本地全局存储
func GenerateDevId(c *gin.Context) {
	result := make(map[string]interface{})
	//查询库中是否有Client-Id
	clientId := getClientId(c)
	if clientId != "" {
		var bean model.TDev
		has, err := global.Db.Where("client_id = ?", clientId).Get(&bean)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("db连接出错, clientId: %s", clientId)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
		if has {
			result["dev_id"] = fmt.Sprint(bean.Id)
			response.RespOk(c, i18n.RetMsgSuccess, result)
			return
		}
	}

	id, _ := service.GenSnowflake()
	userAgent := c.GetHeader("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	if os == "" {
		os = userAgent
	}
	dev := &model.TDev{
		Id:        id,
		Os:        os,
		ClientId:  clientId,
		Network:   constant.NetworkAutoMode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsSend:    2,
		Comment:   "",
	}
	rows, err := global.Db.Insert(dev)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("db连接出错, clientId: %s", clientId)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	result["dev_id"] = fmt.Sprint(id)
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func SendEmail(c *gin.Context) {
	param := new(request.SendEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}

	if service.CheckEmailSendFlag(c, param.Email) {
		global.MyLogger(c).Err(fmt.Errorf("发送限制")).Msgf("邮件发送频率限制！email:%s", param.Email)
		response.RespFail(c, i18n.RetMesEmailSendLimit, nil)
		return
	}
	user := new(model.TUser)
	has, err := global.Db.Where("uname = ?", param.Email).Get(user)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("db连接出错, email:%s", param.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if !has {
		global.MyLogger(c).Error().Msgf("邮箱地址未注册, email:%s", param.Email)
		response.RespFail(c, i18n.RetMsgEmailNotReg, nil)
		return
	}
	clientIp := c.ClientIP()
	err = service.SendTelSms(c, param.Email, clientIp)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("邮件发送失败！email:%s", param.Email)
		response.RespFail(c, i18n.RetMsgVerifyCodeSendFail, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSendSuccess)
}

func Reg(c *gin.Context) {
	param := new(request.RegRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	// 参数检查
	if param.Account == "" || param.Passwd == "" || param.EnterPasswd == "" {
		global.MyLogger(c).Error().Msgf("参数无效！param: %+v", *param)
		response.RespFail(c, i18n.RetMsgParamInputInvalid, nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		global.MyLogger(c).Error().Msgf("密码错误！param: %+v", *param)
		response.RespFail(c, i18n.RetMsgTwoPasswordNotMatch, nil)
		return
	}
	var counts int64
	_, err := global.Db.SQL("select count(*) from t_user where uname = ?", param.Account).Get(&counts)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("db连接出错, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if counts > 0 {
		global.MyLogger(c).Error().Msgf("账号已注册, account: %s, param: %+v", param.Account, *param)
		response.RespFail(c, i18n.RetMsgEmailHasRegErr, nil)
		return
	}

	// 通过以上检查后，说明账号没有注册过，或者已经注销了。
	//// 先看用户是否有注册过
	//userInfo, err := service.GetUserByUserName(param.Account)
	//if err != nil {
	//	response.RespFail(c, i18n.RetMsgDBErr, nil)
	//	return
	//}
	//// 注册过，且非已注销，直接报错返回
	//if userInfo != nil && userInfo.Status != constant.UserStatusCancelled {
	//	fmt.Printf("zzzzzzz用户邮箱：%s", param.Account)
	//	response.RespFail(c, i18n.RetMsgEmailHasRegErr, nil)
	//	return
	//}
	/*
		//渠道来源
		var channel int = 1 //默认大陆区域 1-中国；2-俄罗斯；3-其它(英语系)
		var err error
		sourceChannel := c.GetHeader("Channel")
		if sourceChannel != "" {
			channel, err = sourceChannel
			if err != nil {
				response.RespFail(c, lang.Translate("cn", "fail"), nil)
				return
			}
		}
	*/
	channel := c.GetHeader("Channel")
	var sendSec int64 = 0
	//查询库中是否有Client-Id
	clientId := getClientId(c)
	if clientId != "" {
		//查询
		var (
			userFlag          int64
			userCancelledFlag int64
		)
		_, rx := global.Db.SQL("select count(*) as total from t_user where client_id = ?", clientId).Get(&userFlag)
		if rx != nil {
			global.MyLogger(c).Err(rx).Msgf("db连接出错, clientId: %s, param: %+v", clientId, *param)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}

		_, rx = global.Db.SQL("select count(*) as total from t_user_cancelled where client_id = ?", clientId).Get(&userCancelledFlag)
		if rx != nil {
			global.MyLogger(c).Err(rx).Msgf("db连接出错, clientId: %s, param: %+v", clientId, *param)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}

		global.MyLogger(c).Info().Msgf("userFlag:%d, userCancelledFlag: %d", userFlag, userCancelledFlag)
		if userFlag == 0 && userCancelledFlag == 0 {
			//sendSec = 3600
			sendSec = 2 * 60 * 60 // 统一赠送15天 (之前没有送过的)
		}
		//sendSec = 31536000 // TODO: 统一赠送一年

		/*
			var bean model.TDev
				if !has { //送过一次的不再送了
					sendSec += 3600 //此种情况才赠送时间
				}

				has, err := global.Db.Where("client_id = ? and is_send = 2", clientId).Get(&bean)

				if err != nil {
					global.MyLogger(c).Err(err).Msg("db连接出错")
					response.RespFail(c, lang.Translate("cn", "fail"), nil)
					return
				}
				if !has { //送过一次的不再送了
					sendSec += 3600 //此种情况才赠送时间
				}
		*/
	} else if channel != "" {
		sendSec = 2 * 60 * 60 // 统一赠送15天 (通过渠道推广来的)，TODO: 目前没办法校验渠道的有效性
	}
	level := 0
	disablePayment := geo.IsNeedDisablePaymentFeature(c, param.Account)
	if disablePayment {
		sendSec = 24 * 60 * 60 * 365 * 10 // 英国、美国 ip赠送 10年时长
		level = 2
	}

	global.MyLogger(c).Info().Msgf("Email(%s) gifted time(%d)", param.Account, sendSec)

	pwdDecode := util.AesDecrypt(param.Passwd)

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	sess.Begin()

	nowTime := time.Now()

	user := &model.TUser{
		Uname:       param.Account,
		Passwd:      util.MD5(pwdDecode),
		Email:       param.Account,
		Phone:       "",
		Level:       level,
		ExpiredTime: nowTime.Unix() + sendSec,
		V2rayUuid:   "c541b521-17dd-11ee-bc4e-0c9d92c013fb", //暂时写配置文件的UUID
		Status:      0,
		Channel:     channel,
		ChannelId:   1,
		CreatedAt:   nowTime,
		UpdatedAt:   nowTime,
		ClientId:    clientId,
		Comment:     "",
	}
	rows, err := sess.Insert(user)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("添加user出错, clientId: %s, param: %+v", clientId, *param)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	//更新Uuid
	rnd := rand.New(rand.NewSource(user.Id))
	uuid.SetRand(rnd)
	nonce, _ := uuid.NewRandomFromReader(rnd)
	user.V2rayUuid = nonce.String() //正式注册生成uuid
	//user.V2rayUuid = "bf268a88-318f-d58f-0e9f-66d6f066be31" //需要注释
	rows, err = sess.Cols("v2ray_uuid").Where("id = ?", user.Id).Update(user)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("添加user-uuid出错, clientId: %s, param: %+v", clientId, *param)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgRegFailed, nil)
		return
	}

	//赠送表
	if sendSec > 0 {
		gift := &model.TGift{
			UserId:    user.Id,
			OpId:      fmt.Sprint(nowTime.Unix()),
			OpUid:     user.Id,
			Title:     "注册赠送",
			GiftSec:   int(sendSec),
			GType:     1,
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
			Comment:   "",
		}
		rows, err = sess.Insert(gift)
		if err != nil || rows != 1 {
			global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("添加赠送记录出错, clientId: %s, param: %+v", clientId, *param)
			sess.Rollback()
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
	}

	//推荐关系
	team := &model.TUserTeam{
		UserId:     user.Id,
		DirectId:   0,
		DirectTree: "",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Comment:    "",
	}
	if param.InviteCode != "" {
		directTeam := new(model.TUserTeam)
		has, err := global.Db.Where("user_id = ?", param.InviteCode).Get(directTeam)
		if err != nil || !has {
			global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("查询上线用户出错, clientId: %s, param: %+v", clientId, *param)
			sess.Rollback()
			response.RespFail(c, i18n.RetMsgReferrerIDIncorrect, nil)
			return
		}
		team.DirectId = directTeam.UserId
		if directTeam.DirectTree == "" {
			team.DirectTree = fmt.Sprint(directTeam.UserId)
		} else {
			team.DirectTree = fmt.Sprint(directTeam.DirectTree, ",", directTeam.UserId)
		}
	}
	rows, err = sess.Insert(team)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("添加team出错, clientId: %s, param: %+v", clientId, *param)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	sess.Commit()
	response.ResOk(c, i18n.RetMsgRegSuccess)
}

func Login(c *gin.Context) {
	param := new(request.LoginRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	param.Account = strings.TrimSpace(param.Account)
	param.Passwd = strings.TrimSpace(param.Passwd)
	if param.Account == "" || param.Passwd == "" {
		global.MyLogger(c).Error().Msgf("参数无效, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgAccountPasswordEmptyErr, nil)
		return
	}

	// 先看用户是否有注册过
	var userInfo *entity.TUser
	err := dao.TUser.Ctx(c).Where(do.TUser{Uname: param.Account}).Scan(&userInfo)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("检查是否有注册失败, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if userInfo == nil {
		global.MyLogger(c).Error().Msgf("账号不存在！%s， param: %+v", param.Account, *param)
		response.RespFail(c, i18n.RetMsgAccountNotExist, nil)
		return
	}

	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)
	user := new(model.TUser)
	has, err := global.Db.Where("uname = ? and passwd = ? and status = 0", param.Account, pwdMd5).Get(user)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("登录出错！%s， param: %+v", param.Account, *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if !has {
		global.MyLogger(c).Error().Msgf("密码不正确！%s， param: %+v", param.Account, *param)
		response.RespFail(c, i18n.RetMsgPasswordIncorrect, nil)
		return
	}
	devId := c.Request.Header.Get("Dev-Id")
	if devId != "" && user.Email != "zzz@qq.com" {
		dId, err := strconv.ParseInt(devId, 10, 64)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("DevID atoi failed！%s, devId: %s, param: %+v", param.Account, devId, *param)
			response.RespFail(c, i18n.RetMsgDevIdParseErr, nil)
			return
		}

		if !service.HasDev(int64(dId)) {
			global.MyLogger(c).Err(fmt.Errorf("%s", i18n.RetMsgDevIdNotExitsErr)).Msgf("param: %+v", *param)
			response.RespFail(c, i18n.RetMsgDevIdNotExitsErr, nil)
			return
		}

		limits, err := service.CheckDevNumLimits(int64(dId), user)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("登录出错！%s， param: %+v", param.Account, *param)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
		if limits {
			global.MyLogger(c).Error().Msgf("登录出错，设备数量超过限制！%s， param: %+v", param.Account, *param)
			response.RespFail(c, i18n.RetMsgReachedDevicesLimit, nil)
			return
		}
	}
	dataParam := response.LoginClientParam{
		UserId: user.Id,
		Token:  service.GenerateTokenByUser(user.Id, service.CommonUserType),
	}
	response.RespOk(c, i18n.RetMsgSuccess, dataParam)
}

func ChangePasswd(c *gin.Context) {
	param := new(request.ChangePasswdRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		global.MyLogger(c).Error().Msgf("参数Passwd无效, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgTwoPasswordNotMatch, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	oldPwdDecode := util.AesDecrypt(param.OldPasswd)
	oldPwdMd5 := util.MD5(oldPwdDecode)
	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)

	user.Passwd = pwdMd5
	user.UpdatedAt = time.Now()
	rows, err := global.Db.Cols("passwd", "updated_at").Where("id = ? and passwd = ?", user.Id, oldPwdMd5).Update(user)
	if err != nil {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("修改密码失败, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	global.MyLogger(c).Info().Msgf("rows: %d", rows)
	response.ResOk(c, i18n.RetMsgSuccess)
}

func ForgetPasswd(c *gin.Context) {
	param := new(request.ForgetRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		global.MyLogger(c).Err(fmt.Errorf("密码输入不一致")).Msgf("两次密码不一致, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgTwoPasswordNotMatch, nil)
		return
	}
	err := service.VerifyMsg(c, param.Account, param.VerifyCode)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("验证码错误, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgVerificationCodeErr, nil)
		return
	}

	_, err = common.GetUserByEmail(c, param.Account)
	if err != nil {
		return
	}
	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)
	user := new(model.TUser)
	user.Passwd = pwdMd5
	user.UpdatedAt = time.Now()
	_, err = global.Db.Cols("passwd", "updated_at").Where("uname = ? ", param.Account).Update(user)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("修改密码失败, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSuccess)
}

func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	res := response.UserInfoResponse{
		Id:          user.Id,
		Uname:       user.Uname,
		Uuid:        user.V2rayUuid,
		MemberType:  user.Level,
		ExpiredTime: user.ExpiredTime,
		SurplusFlow: 0,
	}
	response.RespOk(c, i18n.RetMsgSuccess, res)
}

func TeamList(c *gin.Context) {
	param := new(request.TeamListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	session := service.TeamList(param, user)
	count, err := service.TeamList(param, user).Count()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("查询出错！email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	cols := "u.uname,u.level,t.created_at"
	session.Cols(cols)
	session.OrderBy("t.id desc")
	dataList, _ := commonPageListV2(c, param.Page, param.Size, count, session)
	var list []response.TeamListResponse
	for _, item := range dataList.List.([]map[string]interface{}) {
		team := response.TeamListResponse{
			Uname:       item["uname"].(string),
			MemberType:  item["level"].(int),
			CreatedTime: item["created_at"].(string),
		}
		list = append(list, team)
	}
	dataList.List = list
	response.RespOk(c, i18n.RetMsgSuccess, dataList)
}

func TeamInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	var fans int64
	var hours float64
	var list []map[string]interface{}
	sql1 := "select count(*) as fans from t_user_team where direct_id = ?"
	sql2 := "select round(sum(gift_sec)/3600,2) as hours from t_gift where user_id = ? and g_type = 2"
	_, err = global.Db.SQL(sql1, user.Id).Get(&fans)
	_, err = global.Db.SQL(sql2, user.Id).Get(&hours)
	err = global.Db.
		Cols("u.uname,g.created_at,g.gift_sec,g.title").
		Table("t_gift as g").
		Where("g.g_type = 2 and g.user_id = ?", user.Id).
		Join("LEFT", "t_user as u", "u.id = g.op_uid").
		OrderBy("g.id desc").
		Limit(10).
		Find(&list)
	var dataList []response.AwardInfo
	if len(list) > 0 {
		for _, item := range list {
			awardInfo := response.AwardInfo{
				Uname:   item["uname"].(string),
				Title:   item["title"].(string),
				GiftSec: item["gift_sec"].(int),
				TimeStr: item["created_at"].(string),
			}
			dataList = append(dataList, awardInfo)
		}
	} else {
		dataList = []response.AwardInfo{}
	}
	res := response.TeamInfoResponse{
		Fans:       fans,
		AwardHour:  decimal.NewFromFloat(hours).Truncate(2).String(),
		AwardMoney: "0",
		AwardList:  dataList,
	}
	response.RespOk(c, i18n.RetMsgSuccess, res)
}

// call
func GetConfig(c *gin.Context) {
	param := new(request.BanDevRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(c).Info().Msgf(">>> param: %+v", *param)
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	// 当参数node_id设置时，表示指定node_id来获取配置；
	// 当参数node_id没有设置时，表示获取全部配置
	var (
		nodeEntity   *entity.TNode
		nodeEntities []entity.TNode
	)
	if param.NodeId > 0 {
		err = dao.TNode.Ctx(c).Where(do.TNode{Id: param.NodeId, Status: 1}).Scan(&nodeEntity)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("数据库链接出错, email: %s", user.Email)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
	}
	if nodeEntity == nil {
		err = dao.TNode.Ctx(c).Where(do.TNode{IsRecommend: 1, Status: 1}).Scan(&nodeEntity)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("数据库链接出错, email: %s", user.Email)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
	}
	if nodeEntity == nil {
		global.MyLogger(c).Err(fmt.Errorf("nodeEntity is null")).Msgf("email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	err = dao.TNode.Ctx(c).Where(do.TNode{
		CountryEn: nodeEntity.CountryEn,
		Status:    1,
	}).Scan(&nodeEntities)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("数据库链接出错, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	//var sqlWhere string
	//if param.NodeId > 0 {
	//	sqlWhere = fmt.Sprintf("id = %d and status = 1", param.NodeId)
	//} else {
	//	sqlWhere = "is_recommend = 1 and status = 1"
	//}
	//
	////else if user.Email == "ru100@qq.com" {
	////	sqlWhere = fmt.Sprintf("status = 1")
	////} /*else {
	////	sqlWhere = fmt.Sprintf("id not in (100003) and status = 1")
	////}*/
	//
	uuid := user.V2rayUuid
	//var list []map[string]interface{}
	//cols := "id,name,title,title_en,country,country_en,server,port," +
	//	"min_port as min,max_port as max,path,is_recommend"
	//errs := global.Db.Where(sqlWhere).
	//	Table("t_node").
	//	Cols(cols).
	//	OrderBy("id desc").
	//	Find(&list)
	//if errs != nil {
	//	global.MyLogger(c).Err(errs).Msgf("数据库链接出错, email: %s", user.Email)
	//	response.RespFail(c, i18n.RetMsgDBErr, nil)
	//	return
	//}
	//var dnsArray = []string{}
	var d_proxy = []string{}
	//var d_data = []string{}
	//var d_proto = []string{}
	//i := 0
	v2rayServs := make([]v2rayConfig.Server, 0)
	for _, item := range nodeEntities {
		nodeId := item.Id
		nodePorts := []int{item.Port}
		for x := item.MinPort; x <= item.MaxPort; x++ {
			nodePorts = append(nodePorts, x)
		}
		global.MyLogger(c).Info().Msgf(">>>>> nodePorts: %+v", nodePorts)

		dnsList, _ := service.FindNodeDnsByNodeId(nodeId, user.Level+1) // user_level+1等于服务器域名的等级

		for _, dns := range dnsList {
			for _, nodePort := range nodePorts {
				//i = i + 1
				//mproxy := "\"proxy" + strconv.Itoa(i) + "\""

				//d_proxy = append(d_proxy, mproxy)
				//m := fmt.Sprintf("{\"tag\": %s,\"protocol\": \"chain\",\"settings\": {\"actors\": [\"tls\",\"ws\",\"trojan%d\"]}}", mproxy, i)
				//d_data = append(d_data, m)
				//np := fmt.Sprintf("{\"protocol\": \"trojan\",\"settings\": {\"address\": \"%s\",\"port\": %d,\"password\": \"%s\"},\"tag\": \"trojan%d\"}", dns.Dns, nodePort, uuid, i)
				//d_proto = append(d_proto, np)
				//name := fmt.Sprintf("trojan%d", i)
				//name := "trojan"
				//retstr := fmt.Sprintf("{\"protocol\": \"trojan\",\"settings\": {\"address\": \"%s\",\"port\": 443,\"password\": \"%s\"},\"tag\": \"%s\"}", dns.Dns, uuid, name)

				//dnsArray = append(dnsArray, retstr)
				v2rayServs = append(v2rayServs, v2rayConfig.Server{Password: uuid, Port: nodePort, Address: dns.Dns})

				mproxy := fmt.Sprintf("{\"password\": \"%s\",\"port\": %d,\"email\": \"\",\"level\": 0,\"flow\": \"\",\"address\": \"%s\"}", uuid, nodePort, dns.Dns)
				d_proxy = append(d_proxy, mproxy)
			}
		}
	}
	global.MyLogger(c).Info().Msgf(">>>>> d_proxy: %+v", d_proxy)
	//mystring := "{\"log\":{\"level\":\"{{logLevel}}\",\"output\":\"{{leafLogFile}}\"},\"dns\":{\"servers\":[\"1.1.1.1\",\"8.8.8.8\"],\"hosts\":{\"node2.wuwuwu360.xyz\":[\"107.148.239.239\"]}},\"inbounds\":[{\"protocol\":\"tun\",\"settings\":{\"fd\":\"{{tunFd}}\"},\"tag\":\"tun_in\"}],\"outbounds\":[{\"protocol\":\"failover\",\"tag\":\"failover_out\",\"settings\":{\"actors\":[%s],\"failTimeout\":4,\"healthCheck\":true,\"checkInterval\":300,\"failover\":true,\"fallbackCache\":false,\"cacheSize\":256,\"cacheTimeout\":60}},%s,{\"protocol\":\"tls\",\"tag\":\"tls\",\"settings\":{\"alpn\":[\"http/1.1\"],\"insecure\":true}},{\"protocol\":\"ws\",\"tag\":\"ws\",\"settings\":{\"path\":\"/work\"}},%s,{\"protocol\":\"direct\",\"tag\":\"direct_out\"},{\"protocol\":\"drop\",\"tag\":\"reject_out\"}],\"router\":{\"domainResolve\":true,\"rules\":[{\"external\":[\"site:{{dlcFile}}:cn\"],\"target\":\"direct_out\"},{\"external\":[\"mmdb:{{geoFile}}:cn\"],\"target\":\"direct_out\"},{\"domainKeyword\":[\"apple\",\"icloud\"],\"target\":\"direct_out\"}]}}"
	//mystring := "{\"routing\":{\"rules\":[{\"type\":\"field\",\"outboundTag\":\"direct\",\"domain\":[\"icloud\",\"apple\",\"geosite:private\",\"geosite:cn\"]},{\"ip\":[\"geoip:private\",\"geoip:cn\"],\"outboundTag\":\"direct\",\"type\":\"field\"}],\"domainMatcher\":\"hybrid\",\"domainStrategy\":\"AsIs\",\"balancers\":[]},\"log\":{\"loglevel\":\"warning\",\"dnsLog\":false},\"outbounds\":[{\"tag\":\"proxy\",\"mux\":{\"enabled\":false,\"concurrency\":50},\"protocol\":\"trojan\",\"streamSettings\":{\"wsSettings\":{\"path\":\"/work\",\"headers\":{\"host\":\"\"}},\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"allowInsecure\":true,\"fingerprint\":\"\"},\"security\":\"tls\",\"network\":\"ws\"},\"settings\":{\"servers\":[%s]}},{\"tag\":\"direct\",\"protocol\":\"freedom\"},{\"tag\":\"reject\",\"protocol\":\"blackhole\"}]}"
	//mystring := "{\"routing\":{\"rules\":[{\"type\":\"field\",\"outboundTag\":\"direct\",\"domain\":[\"icloud\",\"apple\",\"geosite:private\"]},{\"ip\":[\"geoip:private\"],\"outboundTag\":\"direct\",\"type\":\"field\"}],\"domainMatcher\":\"hybrid\",\"domainStrategy\":\"AsIs\",\"balancers\":[]},\"log\":{\"loglevel\":\"warning\",\"dnsLog\":false},\"outbounds\":[{\"tag\":\"proxy\",\"mux\":{\"enabled\":false,\"concurrency\":50},\"protocol\":\"trojan\",\"streamSettings\":{\"wsSettings\":{\"path\":\"/work\",\"headers\":{\"host\":\"\"}},\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"allowInsecure\":true,\"fingerprint\":\"\"},\"security\":\"tls\",\"network\":\"ws\"},\"settings\":{\"servers\":[%s]}},{\"tag\":\"direct\",\"protocol\":\"freedom\"},{\"tag\":\"reject\",\"protocol\":\"blackhole\"}]}"
	//if user.Email == "ru101@qq.com" {
	//	mystring = "{\"routing\":{\"rules\":[{\"type\":\"field\",\"outboundTag\":\"proxy\",\"domain\":[\"regexp:.*\"]}],\"domainMatcher\":\"hybrid\",\"domainStrategy\":\"AsIs\",\"balancers\":[]},\"log\":{\"loglevel\":\"warning\",\"dnsLog\":false},\"outbounds\":[{\"tag\":\"proxy\",\"mux\":{\"enabled\":false,\"concurrency\":50},\"protocol\":\"trojan\",\"streamSettings\":{\"wsSettings\":{\"path\":\"/work\",\"headers\":{\"host\":\"\"}},\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"allowInsecure\":true,\"fingerprint\":\"\"},\"security\":\"tls\",\"network\":\"ws\"},\"settings\":{\"servers\":[%s]}},{\"tag\":\"direct\",\"protocol\":\"freedom\"},{\"tag\":\"reject\",\"protocol\":\"blackhole\"}]}"
	//}
	//if user.Email == "test12345@qq.com" {
	v, err := json.Marshal(v2rayConfig.GenV2rayConfig(c, v2rayServs, nodeEntity.CountryEn, false))
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GenV2rayConfig failed, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	//global.MyLogger(c).Info().Msgf(">>> get_conf 1 >>> user.Email: %s, config: %+v", user.Email, string(v))
	//global.MyLogger(c).Info().Msgf(">>> get_conf 2 >>> user.Email: %s, config: %+v", user.Email, fmt.Sprintf(mystring, strings.Join(d_proxy, ",")))
	c.String(http.StatusOK, fmt.Sprintf(string(v)))
	//} else {
	//	global.MyLogger(c).Info().Msgf(">>> get_conf >>> user.Email: %s, mystring: %+v", user.Email, mystring)
	//	c.String(http.StatusOK, fmt.Sprintf(mystring, strings.Join(d_proxy, ",")))
	//}

	//	d_proxy,d_data,d_proto)
	//c.String(http.StatusOK, fmt.Sprintf(configs, strings.Join(dnsArray, ",")))
	/*
		param := new(request.NoticeListRequest)
		if err := c.ShouldBind(param); err != nil {
			global.MyLogger(c).Err(err).Msg("绑定参数")
			response.RespFail(c, lang.Translate("cn", "fail"), nil)
			return
		}
		session := service.NoticeList(param)
		count, err := service.NoticeList(param).Count()
		if err != nil {
			global.MyLogger(c).Err(err).Msg(i18n.RetMsgDBErr)
			response.RespFail(c, i18n.RetMsgDBErr)
			return
		}
		cols := "n.id,n.title,n.tag,n.created_at"
		session.Cols(cols)
		session.OrderBy("n.id desc")
		dataList, _ := commonPageListV2(c, param.Page, param.Size, count, session)
		response.RespOk(c, "成功", dataList)
		response.RespFail(c, "推荐人ID不正确", nil)
	*/

}
func AppInfo(c *gin.Context) {
	/*
		host := "http://" + c.Request.Host
		gateWay := host + "/app-upload"
		var version model.TAppVersion
		has, err := global.Db.Where("status = 1 and app_type = 3").OrderBy("id desc").Limit(1).Get(&version)
		if err != nil || !has {
			global.MyLogger(c).Err(err).Msg("key不存在！")
			response.RespFail(c, "失败！")
			return
		}
		type Result struct {
			Code int    `json:"code"`
			Msg  string `json:"message"`
		}
		c.JSON(http.StatusOK, Result{
			Code: -1,
			Msg:  gateWay + version.Link,
		})
		return
	*/
	disablePayment := geo.IsNeedDisablePaymentFeature(c, "")
	if disablePayment {
		response.RespOk(c, "", "")
		return
	}
	host := "http://" + c.Request.Host
	gateWay := host + "/app-upload"
	var list []*model.TDict
	err := global.Db.Where("key_id = ?", "app_link").
		//Or("key_id = ?", "app_js_zip").
		//Or("key_id = ?", "app_version").
		Find(&list)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("key不存在！clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	var result = make(map[string]interface{})
	for _, item := range list {
		result[item.KeyId] = item.Value
	}
	var version model.TAppVersion
	has, err := global.Db.Where("status = 1 and app_type = 3").OrderBy("id desc").Limit(1).Get(&version)
	if err != nil || !has {
		global.MyLogger(c).Err(fmt.Errorf("err: %+v", err)).Msgf("key不存在！clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	result["app_version"] = version.Version
	result["app_js_zip"] = gateWay + version.Link
	result["app_zip_hash"] = "xxx"
	response.RespOk(c, gateWay+version.Link, result)

}

func PCAppInfo(c *gin.Context) {
	/*
		host := "http://" + c.Request.Host
		gateWay := host + "/app-upload"
		var version model.TAppVersion
		has, err := global.Db.Where("status = 1 and app_type = 3").OrderBy("id desc").Limit(1).Get(&version)
		if err != nil || !has {
			global.MyLogger(c).Err(err).Msg("key不存在！")
			response.RespFail(c, "失败！")
			return
		}
		type Result struct {
			Code int    `json:"code"`
			Msg  string `json:"message"`
		}
		c.JSON(http.StatusOK, Result{
			Code: -1,
			Msg:  gateWay + version.Link,
		})
		return
	*/
	disablePayment := geo.IsNeedDisablePaymentFeature(c, "")
	if disablePayment {
		response.RespOk(c, "", "")
		return
	}
	host := "http://" + c.Request.Host
	gateWay := host + "/app-upload"
	var list []*model.TDict
	err := global.Db.Where("key_id = ?", "app_link").
		//Or("key_id = ?", "app_js_zip").
		//Or("key_id = ?", "app_version").
		Find(&list)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("key不存在！clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	var result = make(map[string]interface{})
	for _, item := range list {
		result[item.KeyId] = item.Value
	}
	var version model.TAppVersion
	has, err := global.Db.Where("status = 1 and app_type = 4").OrderBy("id desc").Limit(1).Get(&version)
	if err != nil || !has {
		global.MyLogger(c).Err(fmt.Errorf("err: %+v", err)).Msgf("key不存在！clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	result["app_version"] = version.Version
	result["app_js_zip"] = gateWay + version.Link
	result["app_zip_hash"] = "xxx"
	response.RespOk(c, gateWay+version.Link, result)

}

func NoticeList(c *gin.Context) {
	param := new(request.NoticeListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	session := service.NoticeList(param)
	count, err := service.NoticeList(param).Count()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("NoticeList failed, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	cols := "n.id,n.title,n.tag,n.created_at"
	session.Cols(cols)
	session.OrderBy("n.id desc")
	dataList, _ := commonPageListV2(c, param.Page, param.Size, count, session)
	response.RespOk(c, i18n.RetMsgSuccess, dataList)
}

func NoticeDetail(c *gin.Context) {
	param := new(request.NoticeDetailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	var notice model.TNotice
	has, err := global.Db.Where("id = ?", param.Id).Get(&notice)
	if err != nil || !has {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("notice不存在, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgQueryResultIsEmpty, nil)
		return
	}
	var result = make(map[string]interface{})
	result["id"] = notice.Id
	result["title"] = notice.Title
	result["content"] = notice.Content
	result["tag"] = notice.Tag
	result["created_at"] = notice.CreatedAt
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func ReceiveFree(c *gin.Context) {
	var rows int64
	var err error
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	//1天只能领取3次
	var counts int64
	todayStr := time.Now().Format("2006-01-02")
	_, err = global.Db.SQL("select count(*) from t_activity where user_id = ? and created_at >= ?", user.Id, todayStr).Get(&counts)
	if counts >= 3 {
		global.MyLogger(c).Error().Msgf("领取次数超过限制，email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgActivity3TimesLimits, nil)
		return
	}

	var status = 2
	var giftSec = 0
	rand.Seed(time.Now().UnixNano())
	var randNum = rand.Intn(100)
	if randNum >= 60 {
		status = 1
		giftSec = 3600
	}

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	sess.Begin()

	id, _ := service.GenSnowflake()
	activity := &model.TActivity{
		Id:        id,
		UserId:    user.Id,
		Status:    status,
		GiftSec:   giftSec,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}

	rows, err = sess.Insert(activity)
	if err != nil || rows < 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("领取失败, email: %s", user.Email)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}

	if status == 1 {
		gift := &model.TGift{
			UserId:    user.Id,
			OpId:      fmt.Sprint(id),
			OpUid:     user.Id,
			Title:     "免费领会员",
			GiftSec:   giftSec,
			GType:     3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Comment:   "",
		}
		rows, err = sess.Insert(gift)
		if err != nil || rows < 1 {
			global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("领取失败, email: %s", user.Email)
			sess.Rollback()
			response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			return
		}

		nowTimeUnix := time.Now().Unix()
		var newExpiredTime int64
		if user.ExpiredTime < nowTimeUnix {
			newExpiredTime = nowTimeUnix + int64(giftSec)
		} else {
			newExpiredTime = user.ExpiredTime + int64(giftSec)
		}
		user.ExpiredTime = newExpiredTime
		user.UpdatedAt = time.Now()
		rows, err = sess.Cols("expired_time", "updated_at").Where("id = ?", user.Id).Update(user)
		if err != nil || rows < 1 {
			global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("更新用户状态失败, email: %s", user.Email)
			sess.Rollback()
			response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			return
		}
	}

	sess.Commit()
	//下发服务器配置给客户端
	result := make(map[string]interface{})
	result["status"] = status
	result["hours"] = giftSec / 3600
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func ReceiveFreeSummary(c *gin.Context) {
	result := make(map[string]interface{})
	dateStr := time.Now().Format("2006-01-02")
	_, err := global.Db.SQL("select count(id) as nums,ROUND(IFNULL(sum(gift_sec)/3600,0),2) as hours from t_activity where created_at > ?", dateStr).Get(&result)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("ReceiveFreeSummary failed, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

// call
func NodeList(c *gin.Context) {
	//la := c.GetHeader("Lang")
	//用户评级
	level := 1 //默认1
	token := c.Request.Header.Get("Authorization-Token")
	if token != "" {
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("ParseTokenByUser failed, clientId: %s", getClientId(c))
			response.RespFail(c, i18n.RetMsgAuthExpired, nil, response.CodeTokenExpired)
			return
		}
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
			response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
			return
		}
		level = service.RatingMemberLevel(user)
	}
	var list []map[string]interface{}
	cols := "id,name,title,title_en,country,country_en,server,port," +
		"min_port as min,max_port as max,path,is_recommend"
	err := global.Db.Where("status = 1").
		Table("t_node").
		Cols(cols).
		OrderBy("id desc").
		Find(&list)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("数据库链接出错, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	for _, item := range list {
		var dnsArray []map[string]interface{}
		nodeId := item["id"].(int64)
		dnsList, _ := service.FindNodeDnsByNodeId(nodeId, level)
		for _, dns := range dnsList {
			var dnsItem = make(map[string]interface{})
			dnsItem["id"] = dns.Id
			dnsItem["node_id"] = dns.NodeId
			dnsItem["dns"] = util.AesEncrypt(dns.Dns)
			//dnsItem["ip"] = dns.Ip
			dnsItem["level"] = dns.Level
			dnsArray = append(dnsArray, dnsItem)
		}
		item["dns_list"] = dnsArray
	}
	var result = make(map[string]interface{})
	result["list"] = list
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func DnsList(c *gin.Context) {
	//用户评级
	level := 1 //默认1
	status := 1
	token := c.Request.Header.Get("Authorization-Token")
	if token != "" {
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("ParseTokenByUser failed, clientId: %s", getClientId(c))
			response.RespFail(c, i18n.RetMsgAuthExpired, nil, response.CodeTokenExpired)
			return
		}
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
			response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
			return
		}
		level = service.RatingMemberLevel(user)
	}
	if global.GetDevId(c) == "1733336209297510400" || global.GetClientId(c) == "9782DC21-7EC9-46FA-A70F-3D37FBF5AED0" {
		level = 100
		status = 100
	}

	var list []map[string]interface{}
	cols := "id,site_type,dns"
	err := global.Db.Where("status = ? and level = ?", status, level).
		Table("t_app_dns").
		Cols(cols).
		OrderBy("id desc").
		Find(&list)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("数据库链接出错, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	for _, item := range list {
		item["dns"] = util.AesEncrypt(item["dns"].(string))
	}
	var result = make(map[string]interface{})
	result["list"] = list
	global.MyLogger(c).Info().Msgf("list: %+v", result)
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func ComboList(c *gin.Context) {
	var result = make(map[string]interface{})
	var list []map[string]interface{}
	cols := "g.id,g.m_type,g.title,g.price,g.period,g.dev_limit,g.flow_limit,g.is_discount,g.low,g.high"
	err := global.Db.Table("t_goods as g").Cols(cols).Where("status = 1").OrderBy("id desc").Find(&list)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("数据库链接出错, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	//获取USD价格
	usdCny := 7 //实际从redis中获取
	baseGoods := make(map[string]interface{})
	for _, item := range list {
		mType := item["m_type"].(int32)
		period := item["period"].(int32)
		if period == 7 {
			baseGoods[fmt.Sprint(mType, "_", period)] = item["price"]
		}
	}
	for _, item := range list {
		mType := item["m_type"].(int32)
		period := item["period"].(int32)
		price, _ := decimal.NewFromString(item["price"].(string))
		basePrice, _ := decimal.NewFromString(baseGoods[fmt.Sprint(mType, "_", 7)].(string))
		baseConvertPrice := basePrice.Div(decimal.NewFromInt(7))
		convertPrice := price.Div(decimal.NewFromInt(int64(period)))
		cheapPercent := (baseConvertPrice.Sub(convertPrice)).Div(baseConvertPrice).Mul(decimal.NewFromInt(100))
		item["price"] = price.StringFixed(2)
		item["coin"] = "USD"
		item["convert_price"] = convertPrice.StringFixed(2)
		item["cheap_percent"] = cheapPercent.StringFixed(0)
		item["cny_price"] = price.Mul(decimal.NewFromInt(int64(usdCny))).StringFixed(2)
		isDiscount := item["is_discount"].(int32)
		if isDiscount == 1 {
			item["discount_title"] = fmt.Sprintf("随机送%v-%v天", item["low"], item["high"])
		} else {
			item["discount_title"] = ""
		}
	}

	//if len(list) == 0 {
	//	result["list"] = []map[string]interface{}{}
	//	response.RespOk(c, i18n.RetMsgSuccess, result)
	//}
	//var tmpMap = make(map[int]interface{})
	//for _, item := range list {
	//	tmpMap[item["m_type"].(int)] = true
	//}
	//var rList []map[string]interface{}
	//for k, _ := range tmpMap {
	//	var tMap = make(map[string]interface{})
	//	var tmpList []map[string]interface{}
	//	for _, item := range list {
	//		if item["m_type"].(int) == k {
	//			tmpList = append(tmpList, item)
	//		}
	//	}
	//	tMap["m_type"] = k
	//	tMap["arr"] = tmpList
	//	rList = append(rList, tMap)
	//}
	result["list"] = list
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func ExpireUserList(c *gin.Context) {
	var list []map[string]interface{}
	ex_time := time.Now().Unix()
	global.Db.Table("t_user").Where("expired_time < ? and expired_time+1200>?", ex_time, ex_time).Find(&list)
	result := make(map[string]interface{})
	result["list"] = list
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

// call
func AdList(c *gin.Context) {
	var list []map[string]interface{}
	global.Db.Table("t_ad").Where("status = 1").Find(&list)
	result := make(map[string]interface{})
	result["list"] = list
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func UploadLog(c *gin.Context) {
	param := new(request.UploadLogRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	devId := c.Request.Header.Get("Dev-Id")
	var dId int64
	if devId != "" {
		dId, err = strconv.ParseInt(devId, 10, 64)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("设备鉴权失败, email: %s", user.Email)
			response.RespFail(c, i18n.RetMsgDeviceAuthFailed, nil)
			return
		}
		if !service.CheckUserDev(dId, user) {
			global.MyLogger(c).Err(fmt.Errorf("设备鉴权失败 err:%+v", err)).Msgf("设备鉴权失败, email: %s", user.Email)
			response.RespFail(c, i18n.RetMsgDeviceAuthFailed, nil)
			return
		}
	}

	log := &model.TUploadLog{
		UserId:    user.Id,
		DevId:     dId,
		Content:   param.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err := global.Db.Insert(log)
	if err != nil || rows < 1 {
		global.MyLogger(c).Err(fmt.Errorf("上传日志失败 err:%+v", err)).Msgf("上传日志失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgUploadLogFailed, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSuccess)
}

func CreateOrder(c *gin.Context) {
	param := new(request.CreateOrderRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	goods := new(model.TGoods)
	has, err := global.Db.Where("id = ? and status = 1", param.GoodsId).Get(goods)
	if err != nil || !has {
		global.MyLogger(c).Err(fmt.Errorf("创建订单失败 err:%+v", err)).Msgf("创建订单失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDealCreateFailed, nil)
		return
	}
	id, _ := service.GenSnowflake()
	order := &model.TOrder{
		Id:        id,
		UserId:    user.Id,
		GoodsId:   param.GoodsId,
		Title:     goods.Title,
		Price:     goods.Price,
		PriceCny:  goods.Price,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err := global.Db.Insert(order)
	if err != nil || rows < 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("创建订单失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDealCreateFailed, nil)
		return
	}
	result := make(map[string]interface{})
	result["oid"] = id
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func OrderList(c *gin.Context) {
	param := new(request.OrderListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	session := service.OrderList(param, user)
	count, err := service.OrderList(param, user).Count()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("OrderList failed, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	cols := "*"
	session.Cols(cols)
	session.OrderBy("o.id desc")
	dataList, _ := commonPageListV2(c, param.Page, param.Size, count, session)
	response.RespOk(c, i18n.RetMsgSuccess, dataList)
}

func DevList(c *gin.Context) {
	param := new(request.DevListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	session := service.UserDevList(param, user)
	count, err := service.UserDevList(param, user).Count()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("UserDevList failed, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	cols := "d.os,ud.*"
	session.Cols(cols)
	session.OrderBy("ud.id desc")
	dataList, _ := commonPageListV2(c, param.Page, param.Size, count, session)
	response.RespOk(c, i18n.RetMsgSuccess, dataList)
}

func BanDev(c *gin.Context) {
	param := new(request.BanDevRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, param: %+v, claims: %+v, clientId: %s",
			*param, *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	userDev := new(model.TUserDev)
	userDev.Status = constant.UserDevBanStatus
	userDev.UpdatedAt = time.Now()
	rows, err := global.Db.Cols("status", "updated_at").
		Where("user_id = ? and dev_id = ? and status = ?", user.Id, param.DevId, constant.UserDevNormalStatus).
		Update(userDev)
	if err != nil || rows < 1 {
		global.MyLogger(c).Err(fmt.Errorf("err: %+v", err)).Msgf("踢除设备失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgRemoveDevFailed, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSuccess)
}

// 连接
func Connect(c *gin.Context) {
	/*
		global.MyLogger(c).Info().Msgf("11This is info log")

		param := new(request.ConnectDevRequest)

		global.MyLogger(c).Info().Msgf(" is info log")
		if err := c.ShouldBind(param); err != nil {
			global.MyLogger(c).Err(err).Msg("绑定参数")
			global.MyLogger(c).Info().Msgf("111mmmThis is info log")
			response.RespFail(c, lang.Translate("cn", "fail"), nil)
			return
		}

		claims := c.MustGet("claims").(*service.CustomClaims)
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.MyLogger(c).Err(err).Msg("用户token鉴权失败")
			response.RespFail(c, i18n.RetMsgAuthFailed)
			return
		}
	*/
	param := new(request.BanDevRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败，clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	global.MyLogger(c).Info().Msgf(">>>>>>>>> user: %s, ExpiredTime: %d", user.Uname, user.ExpiredTime)

	req := &request.NodeAddSubRequest{}
	if user.ExpiredTime > time.Now().Unix() {
		//发送请求：
		req.Tag = "1"
		global.MyLogger(c).Info().Msgf(">>>>>>>>> user: %s, ExpiredTime: %d, Tag: %s", user.Uname, user.ExpiredTime, req.Tag)
	} else {
		req.Tag = "2"
		global.MyLogger(c).Error().Msgf(">>>>>>>>> user: %s, ExpiredTime: %d, Tag: %s", user.Uname, user.ExpiredTime, req.Tag)
	}
	if user.V2rayUuid == "c541b521-17dd-11ee-bc4e-0c9d92c013fb" || user.V2rayUuid == "bf268a88-318f-d58f-0e9f-66d6f066be31" {
		//fmt.Printf("connect ok %s", req.Uuid)
		//response.ResOk(c, i18n.RetMsgSuccess)
		//return
		req.Tag = "1"
	}
	req.Uuid = user.V2rayUuid
	req.Email = user.Email
	req.Level = fmt.Sprintf("%d", user.Level)

	service.UpdateLoginInfo(c, user.Id)

	//url := "https://node2.wuwuwu360.xyz/node/add_sub"
	// 当参数node_id设置时，表示指定node_id来获取配置；
	// 当参数node_id没有设置时，表示获取全部配置
	dnsList, _ := service.FindNodes(user.Level + 1)
	index := 0
	if len(dnsList) > 0 {
		index = gConnectNum % len(dnsList)
	}
	var (
		nodeName          string
		chosenName        string
		recommendNodeName string
		randNodeName      string
	)
	for i, item := range dnsList {
		//mType := item.Server // item["id"].(int32)
		//period := item["period"].(int32)
		//if period == 7 {
		//	baseGoods[fmt.Sprint(mType, "_", period)] = item["price"]
		//}
		//if user.Email != "ru100@qq.com" {
		//	global.MyLogger(c).Info().Msgf(">>>>>>>>> user: %s is not whitelist, skip %s", user.Email, item.Server)
		//	continue
		//}

		url := fmt.Sprintf("https://%s/site-api/node/add_sub", item.Server)
		if strings.Contains(item.Server, "http") {
			url = fmt.Sprintf("%s/node/add_sub", item.Server)
		}
		timestamp := fmt.Sprint(time.Now().Unix())
		headerParam := make(map[string]string)
		res := new(response.Response)
		headerParam["timestamp"] = timestamp
		headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
		global.MyLogger(c).Info().Msgf("33333:level:%d,req.Tag:%s,udid:%s,email:%s,url:%s,level:%s", user.Level, req.Tag, req.Uuid, req.Email, url, req.Level)
		err = util.HttpClientPostV2(url, headerParam, req, res)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("email: %s, 发送失败 %s", user.Email, err.Error())
			//response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			//return
			continue
		}
		if param.NodeId > 0 && item.Id == param.NodeId {
			chosenName = item.Name
		}
		if i == index {
			randNodeName = item.Name
		}
		if item.IsRecommend == constant.NodeRecommendFlag {
			recommendNodeName = item.Name
		}
	}
	if chosenName != "" {
		nodeName = chosenName
	} else if recommendNodeName != "" {
		nodeName = recommendNodeName
	} else if randNodeName != "" {
		nodeName = randNodeName
	}
	if req.Tag == "1" {
		gConnectNum = gConnectNum + 1
		if gConnectNum >= 1000000 {
			gConnectNum = 0
		}
		var result = make(map[string]interface{})
		result["node_name"] = nodeName
		response.RespOk(c, i18n.RetMsgSuccess, result)
		global.MyLogger(c).Info().Msgf(">>>>>>>>> user: %s, result: %v", user.Uname, result)
	} else {
		response.RespFail(c, i18n.RetMsgAccountExpired, nil)
		global.MyLogger(c).Error().Msgf(">>>>>>>>> 过期 email: %s, user: %s, result: nil", user.Email, user.Uname)
	}
	return
}

func CLoseConnect(c *gin.Context) {

}

func getDevIdFromHeader(c *gin.Context) int64 {
	devId := c.Request.Header.Get("Dev-Id")
	devIdInt, err := strconv.ParseInt(devId, 10, 64)
	if err != nil {
		global.MyLogger(c).Err(err).Msg("参数DevId错误")
		return 0
	}
	return devIdInt
}

// ChangeNetwork 切换节点工作
func ChangeNetwork(c *gin.Context) {
	param := new(request.ChangeNetworkRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	if param.WorkMode != 1 && param.WorkMode != 2 {
		global.MyLogger(c).Error().Msgf("参数错误, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgParamInvalid, nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}
	devId := c.Request.Header.Get("Dev-Id")
	dId, err := strconv.ParseInt(devId, 10, 64)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("设备鉴权失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDeviceAuthFailed, nil)
		return
	}
	if !service.CheckUserDev(int64(dId), user) {
		global.MyLogger(c).Error().Msgf("设备鉴权失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDeviceAuthFailed, nil)
		return
	}

	//判断VIP级别是否满足，并远程添加UUID

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	sess.Begin()

	var rows int64
	workMode := new(model.TWorkMode)
	workMode.UpdatedAt = time.Now()
	workMode.ModeType = param.WorkMode
	rows, err = sess.Cols("updated_at", "mode_type").Where("user_id = ? and dev_id = ?", user.Id, devId).Update(workMode)
	if err != nil || rows < 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("workMode出错, email: %s", user.Email)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	workLog := &model.TWorkLog{
		UserId:    user.Id,
		DevId:     int64(dId),
		ModeType:  param.WorkMode,
		NodeId:    0,
		Flow:      0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err = sess.Insert(workLog)
	if err != nil || rows < 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf("workLog出错, email: %s", user.Email)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	sess.Commit()
	//下发服务器配置给客户端
	result := make(map[string]interface{})
	result["node_id"] = 1
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func SwitchButtonStatus(c *gin.Context) {
	//param := new(request.SwitchButtonStatusRequest)
	//if err := c.ShouldBind(param); err != nil {
	//	global.MyLogger(c).Err(err).Msg("绑定参数")
	//	response.RespFail(c, lang.Translate("cn", "fail"), nil)
	//	return
	//}
	//claims := c.MustGet("claims").(*service.CustomClaims)
	//user, err := service.GetUserByClaims(claims)
	//if err != nil {
	//	global.MyLogger(c).Err(err).Msg("用户token鉴权失败")
	//	response.RespFail(c, i18n.RetMsgAuthFailed)
	//	return
	//}
}

func AppFilter(c *gin.Context) {
	var list []*model.TDict
	global.Db.Where("key_id = ?", "filter_pac").
		Or("key_id = ?", "filter_refuse").
		Find(&list)

	var result = make(map[string]interface{})
	for _, item := range list {
		if item.KeyId == "filter_pac" {
			result["poc_filter"] = item.Value
		} else if item.KeyId == "filter_refuse" {
			result["refuse_filter"] = item.Value
		}

	}
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func PrintParam() gin.HandlerFunc {
	/*
		return func(c *gin.Context) {
			fmt.Println(666, c.Request.Header)

			bodyBytes := api.ReadBodyToCache(c)
			fmt.Printf("request={%s}, data={%v}", c.Request.RequestURI, string(bodyBytes))
		}
	*/
	return func(c *gin.Context) {
		fmt.Println(666, c.Request.Header)
		fmt.Println(777, c.Request.Body)
		fmt.Println(999, c.Request.Form)

		bodyBytes := api.ReadBodyToCache(c)
		fmt.Printf("request={%s}, data={%v}", c.Request.RequestURI, string(bodyBytes))
	}

}

// JWTAuth 验证token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.Config.TestEnv == "mac-local-test-env-KJHKJhkiuyqwe(*&12" {
			claims := &service.CustomClaims{UserId: 10123}
			c.Set("claims", claims)
			global.MyLogger(c).Info().Msgf("local-test skip auth")
			return
		}
		token := c.Request.Header.Get("Authorization-Token")
		if token == "" {
			global.MyLogger(c).Err(fmt.Errorf("token is nil")).Msgf("token is nil, clientId: %s", getClientId(c))
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": i18n.I18nTrans(c, i18n.RetMsgAuthorizationTokenInvalid),
			})
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("ParseTokenByUser failed, clientId: %s", getClientId(c))
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": i18n.I18nTrans(c, i18n.RetMsgAuthExpired),
			})
			c.Abort()
			return
		}
		service.UpdateLoginInfo(c, claims.UserId)
		err = service.AddLog(c, claims.UserId)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("AddLog failed, clientId: %s, userId: %d", getClientId(c), claims.UserId)
			c.JSON(http.StatusOK, gin.H{
				"code":    100,
				"message": i18n.I18nTrans(c, i18n.RetMsgDevIdInvalid),
			})
			c.Abort()
			return
		}

		devId := c.Request.Header.Get("Dev-Id")
		if devId != "" {
			var userDev model.TUserDev
			has, err := global.Db.Where("user_id = ? and dev_id = ? and status = 2 ", claims.UserId, devId).Get(&userDev)
			if err != nil {
				global.MyLogger(c).Err(err).Msgf("get TUserDev failed, clientId: %s, userId: %d", getClientId(c), claims.UserId)
				c.JSON(http.StatusOK, gin.H{
					"code":    100,
					"message": i18n.I18nTrans(c, i18n.RetMsgDevIdInvalid),
				})
				c.Abort()
				return
			}
			if has {
				global.MyLogger(c).Err(err).Msgf("授权已过期, clientId: %s, userId: %d, devId: %s", getClientId(c), claims.UserId, devId)
				c.JSON(http.StatusOK, gin.H{
					"code":    301,
					"message": i18n.I18nTrans(c, i18n.RetMsgAuthExpired),
				})
				c.Abort()
				return
			}

		}
		common.SaveDeviceID(c, claims.UserId)
		c.Set("claims", claims)
		//uu := c.MustGet("claims").(*service.CustomClaims)
		//fmt.Println("claims...", uu.UserId)
	}
}

func commonPageListV2(c *gin.Context, page, size int, total int64, session *xorm.Session) (response.PageResult, error) {
	if size >= constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	offset := 0
	if (page - 1) > 0 {
		offset = (page - 1) * size
	}
	var list []map[string]interface{}
	err := session.
		Limit(size, offset).
		Find(&list)
	if err != nil {
		global.MyLogger(c).Err(err).Msg(i18n.RetMsgDBErr)
		return response.PageResult{}, err
	}
	if len(list) == 0 {
		list = []map[string]interface{}{}
	}
	var dataList response.PageResult
	dataList.Total = total
	dataList.Page = page
	dataList.Size = size
	dataList.List = list
	return dataList, nil
}

// 注销账户
func CancelAccount(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	err = sess.Begin()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("注销失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
		return
	}

	userCancelled := &model.TUserCancelled{
		Id:          user.Id,
		Uname:       user.Uname,
		Passwd:      user.Passwd,
		Email:       user.Email,
		Phone:       user.Phone,
		Level:       user.Level,
		ExpiredTime: user.ExpiredTime,
		V2rayUuid:   user.V2rayUuid,
		V2rayTag:    user.V2rayTag,
		ChannelId:   user.ChannelId,
		Channel:     user.Channel,
		Status:      constant.UserStatusCancelled,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   time.Now(),
		Comment:     user.Comment,
		ClientId:    user.ClientId,
	}
	rows, err := sess.Insert(userCancelled)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).
			Msgf("注销账号失败, Insert cancelled user, rows: %d, err: %v, email: %s", rows, err, user.Email)
		_ = sess.Rollback()
		response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
		return
	}

	rows, err = sess.Delete(user)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).
			Msgf("注销账号失败, Delete cancelled user, rows: %d, err: %v, email: %s", rows, err, user.Email)
		_ = sess.Rollback()
		response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
		return
	}

	err = sess.Commit()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("注销账号Commit失败, Commit err: %v, email: %s", err, user.Email)
		_ = sess.Rollback()
		response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSuccess)
	return
}

// SaveUserConfig 存储用户配置，例如：选择的节点ID等
func SaveUserConfig(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	param := new(request.SaveUserConfigRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数, claims: %+v, clientId: %s, email: %s", *claims, getClientId(c), user.Email)
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}

	// 查询用户配置表，是否存在配置项
	userConfig, err := service.GetUserConfig(user.Id)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetUserConfig failed, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if userConfig == nil {
		err = service.CreateUserConfig(user.Id, param.NodeId)
	} else {
		err = service.UpdateUserConfig(user.Id, param.NodeId)
	}
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("保存用户配置失败, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSuccess)
	return
}

func GetUserConfig(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	// 查询用户配置表，是否存在配置项
	userConfig, err := service.GetUserConfig(user.Id)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetUserConfig failed, claims: %+v, Email: %s", *claims, user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if userConfig != nil && userConfig.Status == constant.UserConfigStatusNormal {
		resp := request.GetUserConfigResponse{
			UserId:    userConfig.UserId,
			NodeId:    userConfig.NodeId,
			CreatedAt: userConfig.CreatedAt.Format(constant.TimeFormat),
			UpdatedAt: userConfig.UpdatedAt.Format(constant.TimeFormat),
		}
		response.RespOk(c, i18n.RetMsgSuccess, resp)
	} else {
		response.RespOk(c, i18n.RetMsgSuccess, nil)
	}
	return
}

// DeleteUserConfig 删除用户配置
func DeleteUserConfig(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v, clientId: %s", *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	// 查询用户配置表，是否存在配置项
	userConfig, err := service.GetUserConfig(user.Id)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetUserConfig failed, Email: %+v, clientId: %s", user.Email, getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if userConfig == nil || userConfig.Status == constant.UserConfigStatusDeleted {
		response.ResOk(c, i18n.RetMsgSuccess)
		return
	}
	err = service.DeleteUserConfig(user.Id)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("删除用户配置失败, Email: %+v, clientId: %s", user.Email, getClientId(c))
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	response.ResOk(c, i18n.RetMsgSuccess)
	return
}

// call
func TrafficList(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetUserByClaims failed")
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	_, err = service.GetUserTrafficCurrentMonth(user.Email)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetUserTrafficCurrentMonth failed")
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	//for _, item := range items {
	//	// 汇总信息
	//	// 单条信息
	//
	//}
	//result["list"] = lis
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}

func getClientId(c *gin.Context) string {
	return global.GetClientId(c)
}

func isValidSignature(signature string) bool {
	keyBytes := []byte(constant.SecretKey)
	hmacKey := hmac.New(sha256.New, keyBytes)
	hmacKey.Write(keyBytes)
	Signature := hex.EncodeToString(hmacKey.Sum(nil))
	fmt.Println("Signature", Signature)
	return signature == Signature
}

// 签名验证
func Verify(c *gin.Context) {
	signature := c.Request.Header.Get("X-Signature")
	if !isValidSignature(signature) {
		global.MyLogger(c).Err(fmt.Errorf("Invalid signature")).Msgf("Invalid signature: %s", signature)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		c.Abort()
		return
	}
	var isAllowed bool
	AllowedIP := []string{"127.0.0.1", "185.22.154.46", "45.251.243.140"} // ip池白名单
	// 获取请求IP
	ip := c.ClientIP()
	fmt.Printf("获取请求的客户端IP: %s", ip)
	// 检查 IP 地址是否在白名单中
	for _, allowedIP := range AllowedIP {
		if ip == allowedIP {
			isAllowed = true
			break
		}
	}
	// 如果 IP 地址不在白名单中，返回未授权的错误响应
	if !isAllowed {
		global.MyLogger(c).Err(fmt.Errorf("IP is Unauthorized")).Msgf("未知的IP没有加入白名单 IP: %s", ip) // 插入日志
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized IP"})
		return
	}
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
func ServerStateSwitching(c *gin.Context) {
	// 获取请求参数
	param := new(request.ServerStateSwitchingRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	switch param.Status {
	case constant.Healthy:
		//上架
		_, err := dao.TNode.Ctx(c).Where(do.TNode{Ip: param.Ip}).Update(do.TNode{Status: 1})
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("数据库查询出错, ip: %s", param.Ip)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
	case constant.Unhealthy:
		//下架
		_, err := dao.TNode.Ctx(c).Where(do.TNode{Ip: param.Ip}).Update(do.TNode{Status: 2})
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("数据库查询出错, ip: %s", param.Ip)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}
	default:
		response.RespFail(c, i18n.RetMsgParamInvalid, nil)
	}
}
