package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"go-speed/api"
	"go-speed/api/api/common"
	v2rayConfig "go-speed/api/api/config"
	"go-speed/api/api/internal"
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
	"go-speed/task"
	"go-speed/util"
	"go-speed/util/geo"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"xorm.io/xorm"
)

var gConnectNum int

// GenerateDevId C端获取DEV_ID，并保存在本地全局存储
func GenerateDevId(c *gin.Context) {
	result := make(map[string]interface{})
	result["dev_id"] = "123123"
	response.RespOk(c, i18n.RetMsgSuccess, result)
}

func SendEmail(c *gin.Context) {
	param := new(request.SendEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("%s 绑定参数失败, clientId: %s", i18n.ErrLabelParams, getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}

	if service.CheckEmailSendFlag(c, param.Email) {
		global.MyLogger(c).Warn().Msgf("邮件发送频率限制！email:%s", param.Email)
		response.RespFail(c, i18n.RetMesEmailSendLimit, nil)
		return
	}
	user := new(model.TUser)
	has, err := global.Db.Where("uname = ?", param.Email).Get(user)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("%s query t_user failed, param: %+v", i18n.ErrLabelDB, *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if !has {
		global.MyLogger(c).Warn().Msgf("%s 邮箱地址未注册！param: %+v", i18n.ErrLabelParams, *param)
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
		global.MyLogger(c).Err(err).Msgf("%s 绑定参数失败, clientId: %s", i18n.ErrLabelParams, getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	// 参数检查
	if param.Account == "" || param.Passwd == "" || param.EnterPasswd == "" {
		global.MyLogger(c).Warn().Msgf("%s 参数无效！param: %+v", i18n.ErrLabelParams, *param)
		response.RespFail(c, i18n.RetMsgParamInputInvalid, nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		global.MyLogger(c).Warn().Msgf("%s 密码错误！param: %+v", i18n.ErrLabelParams, *param)
		response.RespFail(c, i18n.RetMsgTwoPasswordNotMatch, nil)
		return
	}

	// 检查推荐人ID是否有效
	inviteUserId, inviteUserInfo, inviteUserTeam, err := internal.GetInviteUserWithErrRsp(c, param.InviteCode)
	if err != nil {
		return
	}
	if inviteUserId > 0 && inviteUserInfo == nil {
		global.MyLogger(c).Warn().Msgf("%s 推荐人Code无效, param: %+v", i18n.ErrLabelParams, *param)
		response.RespFail(c, i18n.RetMsgReferrerIDIncorrect, nil)
		return
	}

	// 检查账号是否已经注册
	var counts int64
	_, err = global.Db.SQL("select count(*) from t_user where uname = ?", param.Account).Get(&counts)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("%s query t_user failed, param: %+v", i18n.ErrLabelDB, *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if counts > 0 {
		global.MyLogger(c).Warn().Msgf("%s 账号已注册, account: %s, param: %+v", i18n.ErrLabelParams, param.Account, *param)
		response.RespFail(c, i18n.RetMsgEmailHasRegErr, nil)
		return
	}

	channel := c.GetHeader("Channel")
	var sendSec int64 = 0
	level := constant.UserLevelNormal
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
			global.MyLogger(c).Err(rx).Msgf("%s query t_user failed, clientId: %s, param: %+v", i18n.ErrLabelDB, clientId, *param)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}

		_, rx = global.Db.SQL("select count(*) as total from t_user_cancelled where client_id = ?", clientId).Get(&userCancelledFlag)
		if rx != nil {
			global.MyLogger(c).Err(rx).Msgf("%s query t_user_cancelled failed, clientId: %s, param: %+v", i18n.ErrLabelDB, clientId, *param)
			response.RespFail(c, i18n.RetMsgDBErr, nil)
			return
		}

		global.MyLogger(c).Info().Msgf("userFlag:%d, userCancelledFlag: %d", userFlag, userCancelledFlag)
		if userFlag == 0 && userCancelledFlag == 0 {
			sendSec = 2 * 60 * 60 // 统一赠送15天 (之前没有送过的)
			level = constant.UserLevelVIP1
		}
	} else if channel != "" {
		sendSec = 2 * 60 * 60 // 统一赠送15天 (通过渠道推广来的)，TODO: 目前没办法校验渠道的有效性
		level = constant.UserLevelVIP1
	}

	disablePayment := geo.IsNeedDisablePaymentFeature(c, param.Account)
	if disablePayment {
		sendSec = 24 * 60 * 60 * 365 * 10 // 英国、美国 ip赠送 10年时长
		level = constant.UserLevelVIP2
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
		global.MyLogger(c).Err(fmt.Errorf("err:%+v, rows: %d", err, rows)).Msgf("添加user出错, clientId: %s, param: %+v", clientId, *param)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	//更新Uuid
	rnd := rand.New(rand.NewSource(user.Id))
	uuid.SetRand(rnd)
	nonce, _ := uuid.NewRandomFromReader(rnd)
	user.V2rayUuid = nonce.String() //正式注册生成uuid
	rows, err = sess.Cols("v2ray_uuid").Where("id = ?", user.Id).Update(user)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v, rows: %d", err, rows)).Msgf("添加user-uuid出错, clientId: %s, param: %+v", clientId, *param)
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
			global.MyLogger(c).Err(fmt.Errorf("err:%+v, rows: %d", err, rows)).Msgf("添加赠送记录出错, clientId: %s, param: %+v", clientId, *param)
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
	if inviteUserId > 0 {
		if inviteUserTeam != nil {
			team.DirectId = inviteUserTeam.UserId
			if inviteUserTeam.DirectTree == "" {
				team.DirectTree = fmt.Sprint(inviteUserTeam.UserId)
			} else {
				team.DirectTree = fmt.Sprint(inviteUserTeam.DirectTree, ",", inviteUserTeam.UserId)
			}
		} else {
			team.DirectTree = fmt.Sprint(inviteUserId)
		}
	}
	rows, err = sess.Insert(team)
	if err != nil || rows != 1 {
		global.MyLogger(c).Err(fmt.Errorf("err:%+v, rows:%d", err, rows)).Msgf("添加team出错, clientId: %s, param: %+v", clientId, *param)
		sess.Rollback()
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	sess.Commit()
	_ = internal.UpdateUserDeviceByClientId(c, user.Id, user.Email)

	response.ResOk(c, i18n.RetMsgRegSuccess)
}

func Login(c *gin.Context) {
	param := new(request.LoginRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("%s 绑定参数失败, clientId: %s", i18n.ErrLabelParams, getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	param.Account = strings.TrimSpace(param.Account)
	param.Passwd = strings.TrimSpace(param.Passwd)

	if param.Account == "" || param.Passwd == "" {
		global.MyLogger(c).Warn().Msgf("%s 账号或密码为空, param: %+v", i18n.ErrLabelParams, *param)
		response.RespFail(c, i18n.RetMsgAccountPasswordEmptyErr, nil)
		return
	}

	// 先看用户是否有注册过
	var userInfo *entity.TUser
	err := dao.TUser.Ctx(c).Where(do.TUser{Uname: param.Account}).Scan(&userInfo)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("%s 检查是否有注册失败, param: %+v", i18n.ErrLabelDB, *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if userInfo == nil {
		global.MyLogger(c).Warn().Msgf("%s 账号不存在，param: %+v", i18n.ErrLabelParams, *param)
		response.RespFail(c, i18n.RetMsgAccountNotExist, nil)
		return
	}

	pwdDecode := util.AesDecrypt(param.Passwd)
	pwdMd5 := util.MD5(pwdDecode)

	if userInfo.Passwd != pwdMd5 {
		global.MyLogger(c).Warn().Msgf("%s 密码不正确, param: %+v", i18n.ErrLabelParams, *param)
		response.RespFail(c, i18n.RetMsgPasswordIncorrect, nil)
		return
	}

	err = internal.CheckClientIdNumLimitsWithErrRsp(c, userInfo)
	if err != nil {
		return
	}
	// 登陆成功后，老用户挽回赠送免费时长活动
	internal.UserRecovery(c, userInfo)

	// 更新失败时，不返回报错，自己内部评估如何优化
	_ = internal.UpdateUserDeviceByClientId(c, userInfo.Id, userInfo.Email)

	dataParam := response.LoginClientParam{
		UserId: userInfo.Id,
		Token:  service.GenerateTokenByUser(userInfo.Id, service.CommonUserType),
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
		global.MyLogger(c).Warn().Msgf("两次密码不一致, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgTwoPasswordNotMatch, nil)
		return
	}
	err := service.VerifyMsg(c, param.Account, param.VerifyCode)
	if err != nil {
		global.MyLogger(c).Warn().Msgf("验证码校验失败, param: %+v", *param)
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

	var items []entity.TAppDns
	e := dao.TAppDns.Ctx(c).
		Where(do.TAppDns{Status: 1, Level: 1}).
		Cache(gdb.CacheOption{Duration: 10 * time.Minute}).Scan(&items)
	if e != nil {
		global.MyLogger(c).Err(e).Msgf(">>>>> Get TAppDns failed: %+v", e.Error())
	}

	var dnsList []string
	for _, i := range items {
		dnsList = append(dnsList, i.Dns)
	}
	var dns string
	if len(dnsList) > 0 {
		rand.Seed(time.Now().UnixNano())
		dns = dnsList[rand.Intn(len(dnsList))]
	}

	res := response.UserInfoResponse{
		Id:          user.Id,
		Uname:       user.Uname,
		Uuid:        user.V2rayUuid,
		MemberType:  user.Level,
		ExpiredTime: user.ExpiredTime,
		SurplusFlow: 0,
		SpecialFlag: geo.IsNeedDisablePaymentFeature(c, user.Email),
		DNS:         dns,
		Timestamp:   time.Now().Unix(),
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

	uuid := user.V2rayUuid
	var d_proxy = []string{}
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
				v2rayServs = append(v2rayServs, v2rayConfig.Server{Password: uuid, Port: nodePort, Address: dns.Dns})

				mproxy := fmt.Sprintf("{\"password\": \"%s\",\"port\": %d,\"email\": \"\",\"level\": 0,\"flow\": \"\",\"address\": \"%s\"}", uuid, nodePort, dns.Dns)
				d_proxy = append(d_proxy, mproxy)
			}
		}
	}
	global.MyLogger(c).Info().Msgf(">>>>> d_proxy: %+v", d_proxy)
	v, err := json.Marshal(v2rayConfig.GenV2rayConfig(c, v2rayServs, nodeEntity.CountryEn, false))
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GenV2rayConfig failed, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf(string(v)))
}
func AppInfo(c *gin.Context) {
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
		global.MyLogger(c).Err(err).Msg("查询app_link失败")
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
		global.MyLogger(c).Warn().Msgf("领取次数超过限制，email: %s", user.Email)
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
		var affected int64
		affected, err = dao.TUser.Ctx(c).Where(do.TUser{Id: user.Id, Email: user.Email}).
			Data(do.TUser{Kicked: 0}).
			UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("update user Kicked flag failed, email: %s", user.Email)
			response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			return
		}
		global.MyLogger(c).Debug().Msgf("update user Kicked flag, affected: %d", affected)

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
		user.Kicked = 0
		rows, err = sess.Cols("expired_time", "updated_at").Where("id = ?", user.Id).Update(user)
		if err != nil || rows > 1 {
			global.MyLogger(c).Err(fmt.Errorf("err:%+v", err)).Msgf(
				"更新用户状态失败, email: %s, rows: %d", user.Email, rows)
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

	log := &model.TUploadLog{
		UserId:    user.Id,
		DevId:     0,
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
	_, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, param: %+v, claims: %+v, clientId: %s",
			*param, *claims, getClientId(c))
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	//var record *entity.TUserDev
	//err = dao.TUserDev.Ctx(c).Where(do.TUserDev{
	//	UserId: user.Id,
	//	DevId:  param.DevId,
	//}).Scan(&record)
	//if err != nil {
	//	global.MyLogger(c).Err(err).Msgf("get TUserDev failed, email: %s", user.Email)
	//	response.RespFail(c, i18n.RetMsgDBErr, nil)
	//	return
	//}
	//if record == nil || record.Status == constant.UserDevBanStatus {
	//	response.ResOk(c, i18n.RetMsgSuccess)
	//	return
	//}
	//global.MyLogger(c).Debug().Msgf("TUserDev: %+v", *record)
	//
	//affect, err := dao.TUserDev.Ctx(c).Data(do.TUserDev{
	//	Status:    constant.UserDevBanStatus,
	//	UpdatedAt: gtime.Now(),
	//}).Where(do.TUserDev{
	//	UserId: user.Id,
	//	DevId:  param.DevId,
	//}).UpdateAndGetAffected()
	//if err != nil {
	//	global.MyLogger(c).Err(err).Msgf("update userDev failed, email: %s", user.Email)
	//	response.RespFail(c, i18n.RetMsgRemoveDevFailed, nil)
	//	return
	//}
	//global.MyLogger(c).Debug().Msgf("update TUserDev status affect: %d", affect)
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
	req.Uuid = util.GetUserV2rayConfigUUID(user.V2rayUuid)
	req.Email = util.GetUserV2rayConfigEmail(user.Email)
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

		//url := fmt.Sprintf("https://%s/site-api/node/add_sub", item.Server)
		//if strings.Contains(item.Server, "http") {
		//	url = fmt.Sprintf("%s/node/add_sub", item.Server)
		//}
		url := fmt.Sprintf("http://%s:15003/node/add_sub", item.Ip)
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
	devId := ""
	dId := 0

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
			global.MyLogger(c).Warn().Msgf("token is nil, clientId: %s", getClientId(c))
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
			if strings.Contains(err.Error(), "token is expired by ") {
				global.MyLogger(c).Warn().Msgf("token is expired. %s", getClientId(c))
			} else {
				global.MyLogger(c).Err(err).Msgf("ParseTokenByUser failed, clientId: %s", getClientId(c))
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": i18n.I18nTrans(c, i18n.RetMsgAuthExpired),
			})
			c.Abort()
			return
		}
		service.UpdateLoginInfo(c, claims.UserId)
		common.SaveDeviceID(c, claims.UserId)
		c.Set("claims", claims)
		//uu := c.MustGet("claims").(*service.CustomClaims)
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

	//affect, err := dao.TUser.Ctx(c).Data(do.TUser{
	//	Status:    constant.UserStatusCancelled,
	//	UpdatedAt: gtime.Now(),
	//}).Where(do.TUser{
	//	Id:        user.Id,
	//	Email:     user.Email,
	//	V2RayUuid: user.V2rayUuid,
	//}).UpdateAndGetAffected()
	//if err != nil {
	//	global.MyLogger(c).Err(err).Msgf("update user status = 10 failed, email: %s", user.Email)
	//	response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
	//	return
	//}
	//if affect != 1 {
	//	global.MyLogger(c).Err(errors.New("update rows failed")).Msgf("affect: %d, email: %s", affect, user.Email)
	//	response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
	//	return
	//}

	// 删除所有节点上的配置
	err = task.DeleteUserV2rayConfig(c, user.Email, user.V2rayUuid)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("DeleteUser failed, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgLogoutFailed, nil)
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

func InternalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if global.Config.TestEnv == "mac-local-test-env-KJHKJhkiuyqwe(*&12" {
		//	claims := &service.CustomClaims{UserId: 10123}
		//	c.Set("claims", claims)
		//	global.MyLogger(c).Info().Msgf("local-test skip auth")
		//	return
		//}
		token := c.Request.Header.Get("Authorization-Token")
		if token != "J7RtY3DvV2pK0fM5rW4aU1cL8yB9eQ6sI8gH2kZ5xT7uF1oP6vN8jA4lR9mG3bE0wH7nY6tS5zC8iQ1fX9rV6hO5lJ4dU3pV8aB2e(*&12" {
			global.MyLogger(c).Err(fmt.Errorf("token is invalid")).Msgf("token is: %s", token)
			//c.JSON(http.StatusOK, gin.H{
			//	"code":    301,
			//	"message": i18n.I18nTrans(c, i18n.RetMsgAuthorizationTokenInvalid),
			//})
			c.Abort()
			return
		}

		if !util.IsInArrayIgnoreCase(c.ClientIP(), []string{"127.0.0.1", "31.128.41.86", "45.251.243.140", "185.22.152.47", "185.22.154.21"}) {
			global.MyLogger(c).Err(fmt.Errorf("cleint IP is not allow")).Msgf(c.ClientIP())
			c.Abort()
			return
		}
	}
}

// 官网接口，获取后台配置的推广人员与渠道映射关系
func GetPromotionDnsMapping(c *gin.Context) {
	var (
		err      error
		entities []entity.TPromotionDns
	)
	global.MyLogger(c).Info().Msg("GetPromotionDnsMapping start.")
	defer global.MyLogger(c).Info().Msg("GetPromotionDnsMapping end.")
	// 获取请求参数
	req := new(request.GetPromotionDnsMappingRequest)
	if err := c.ShouldBind(req); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败, clientId: %s", getClientId(c))
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}

	whereDo := do.TPromotionDns{
		Status: 1,
	}
	// 获取请求域名的主域名
	// 如 www.yyy360.xyz -> yyy360.xyz
	if req.Dns != "" {
		whereDo.Dns = util.GetDomain(req.Dns)
		global.MyLogger(c).Info().Msgf("GetPromotionDnsMapping host is %s.", whereDo.Dns)
	}

	// 查询数据库
	err = dao.TPromotionDns.Ctx(c).
		Where(whereDo).
		Scan(&entities)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetPromotionDnsMapping DB failed. Error: %v", err)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	// 如果数据为空，则返回错误
	if entities == nil {
		global.MyLogger(c).Warn().Msg("GetPromotionDnsMapping result is empty.")
		response.RespOk(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	items := make([]response.PromotionDnsRes, 0)
	for _, entity := range entities {
		items = append(items, response.PromotionDnsRes{
			AndroidChannel: entity.AndroidChannel,
			WinChannel:     entity.WinChannel,
			MacChannel:     entity.MacChannel,
		})
	}
	// 返回数据
	response.RespOk(c, i18n.RetMsgSuccess, response.PromotionDnsResponse{
		List: items,
	})
}

// 官网接口，下载页面的各个商店的推广链接
func GetPromotionShopMapping(c *gin.Context) {
	var (
		err      error
		entities []entity.TAppStore
	)

	// 查询数据库
	err = dao.TAppStore.Ctx(c).Where(do.TAppStore{Status: 1}).Scan(&entities)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("GetPromotionShopMapping DB failed. Error: %v", err)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	// 组装返回数据
	items := make([]response.PromotionShopRes, 0)
	for _, entity := range entities {
		items = append(items, response.PromotionShopRes{
			Type:    entity.Type,
			Url:     entity.Url,
			TitleCn: entity.TitleCn,
			TitleEn: entity.TitleEn,
			TitleRu: entity.TitleRu,
			Cover:   entity.Cover,
		})
	}

	// 返回数据
	response.RespOk(c, i18n.RetMsgSuccess, response.PromotionShopResponse{
		List: items,
	})
}
