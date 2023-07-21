package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/user_agent"
	"github.com/shopspring/decimal"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/lang"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"xorm.io/xorm"
)

// GenerateDevId C端获取DEV_ID，并保存在本地全局存储
func GenerateDevId(c *gin.Context) {
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
		Network:   constant.NetworkAutoMode,
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
	result["dev_id"] = fmt.Sprint(id)
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
		response.RespFail(c, "参数错误，请检查", nil)
		return
	}
	if param.Account == "" || param.Passwd == "" || param.EnterPasswd == "" {
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if param.Passwd != param.EnterPasswd {
		response.RespFail(c, "两次密码不一致，请检查", nil)
		return
	}

	uuid, _ := uuid.NewUUID()
	pwdDecode := util.AesDecrypt(param.Passwd)

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	sess.Begin()

	user := &model.TUser{
		Uname:       param.Account,
		Passwd:      util.MD5(pwdDecode),
		Email:       param.Account,
		Phone:       "",
		Level:       0,
		ExpiredTime: 0,
		V2rayUuid:   uuid.String(),
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Comment:     "",
	}
	rows, err := sess.Insert(user)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("添加user出错")
		sess.Rollback()
		response.RespFail(c, "用户名重复请检查", nil)
		return
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
			global.Logger.Err(err).Msg("查询上线用户出错")
			sess.Rollback()
			response.RespFail(c, "推荐人ID不正确", nil)
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
		global.Logger.Err(err).Msg("添加team出错")
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
		return
	}
	if !has {
		global.Logger.Err(err).Msgf("用户名或密码不正确！%s", param.Account)
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	devId := c.Request.Header.Get("Dev-Id")
	if devId != "" {
		dId, err := strconv.Atoi(devId)
		if err != nil {
			global.Logger.Err(err).Msgf("登录出错！%s", param.Account)
			response.RespFail(c, lang.Translate("cn", "fail"), nil)
			return
		}
		err = service.UpdateUserDev(int64(dId), user)
		if err != nil {
			global.Logger.Err(err).Msgf("登录出错！%s", param.Account)
			response.RespFail(c, "设备数限制", nil)
			return
		}
		//err = service.UpdateUserWorkMode(int64(dId), user)
		//if err != nil {
		//	global.Logger.Err(err).Msgf("登录出错！%s", param.Account)
		//	response.RespFail(c, lang.Translate("cn", "fail"), nil)
		//}
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
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
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

func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
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
	response.RespOk(c, "成功", res)
}

func TeamList(c *gin.Context) {
	param := new(request.TeamListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	session := service.TeamList(param, user)
	count, err := service.TeamList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "u.uname,u.level,t.created_at"
	session.Cols(cols)
	session.OrderBy("t.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
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
	response.RespOk(c, "成功", dataList)
}

func TeamInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
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
	response.RespOk(c, "成功", res)
}

func AppInfo(c *gin.Context) {
	host := "http://" + c.Request.Host
	gateWay := host + "/app-upload"
	var list []*model.TDict
	err := global.Db.Where("key_id = ?", "app_link").
		//Or("key_id = ?", "app_js_zip").
		//Or("key_id = ?", "app_version").
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("key不存在！")
		response.ResFail(c, "失败！")
		return
	}
	var result = make(map[string]interface{})
	for _, item := range list {
		result[item.KeyId] = item.Value
	}
	var version model.TAppVersion
	has, err := global.Db.Where("status = 1 and app_type = 3").OrderBy("id desc").Limit(1).Get(&version)
	if err != nil || !has {
		global.Logger.Err(err).Msg("key不存在！")
		response.ResFail(c, "失败！")
		return
	}
	result["app_version"] = version.Version
	result["app_js_zip"] = gateWay + version.Link
	result["app_zip_hash"] = "xxx"
	response.RespOk(c, "成功", result)
}

func NoticeList(c *gin.Context) {
	param := new(request.NoticeListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	session := service.NoticeList(param)
	count, err := service.NoticeList(param).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "n.id,n.title,n.tag,n.created_at"
	session.Cols(cols)
	session.OrderBy("n.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func NoticeDetail(c *gin.Context) {
	param := new(request.NoticeDetailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	var notice model.TNotice
	has, err := global.Db.Where("id = ?", param.Id).Get(&notice)
	if err != nil || !has {
		global.Logger.Err(err).Msg("notice不存在")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	var result = make(map[string]interface{})
	result["id"] = notice.Id
	result["title"] = notice.Title
	result["content"] = notice.Content
	result["tag"] = notice.Tag
	result["created_at"] = notice.CreatedAt
	response.RespOk(c, "成功", result)
}

func ReceiveFree(c *gin.Context) {
	var rows int64
	var err error
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
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
		global.Logger.Err(err).Msg("领取失败")
		sess.Rollback()
		response.ResFail(c, "失败！")
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
			global.Logger.Err(err).Msg("领取失败")
			sess.Rollback()
			response.ResFail(c, "失败！")
			return
		}
	}

	sess.Commit()
	//下发服务器配置给客户端
	result := make(map[string]interface{})
	result["status"] = status
	result["hours"] = giftSec / 3600
	response.RespOk(c, "成功", result)
}

func ReceiveFreeSummary(c *gin.Context) {
	result := make(map[string]interface{})
	dateStr := time.Now().Format("2006-01-02")
	_, err := global.Db.SQL("select count(id) as nums,ROUND(IFNULL(sum(gift_sec)/3600,0),2) as hours from t_activity where created_at > ?", dateStr).Get(&result)
	if err != nil {
		fmt.Println(err)
		response.ResFail(c, "查询出错")
		return
	}
	response.RespOk(c, "成功", result)
}

func NodeList(c *gin.Context) {
	//用户评级
	level := 1 //默认1
	token := c.Request.Header.Get("Authorization-Token")
	if token != "" {
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			global.Logger.Err(err).Msg("token出错")
			response.RespFail(c, lang.Translate("cn", "fail"), nil)
			return
		}
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.Logger.Err(err).Msg("用户token鉴权失败")
			response.ResFail(c, "用户鉴权失败！")
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
		global.Logger.Err(err).Msg("数据库链接出错")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
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
	response.RespOk(c, "成功", result)
}

func DnsList(c *gin.Context) {
	//用户评级
	level := 1 //默认1
	token := c.Request.Header.Get("Authorization-Token")
	if token != "" {
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			global.Logger.Err(err).Msg("token出错")
			response.RespFail(c, lang.Translate("cn", "fail"), nil)
			return
		}
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.Logger.Err(err).Msg("用户token鉴权失败")
			response.ResFail(c, "用户鉴权失败！")
			return
		}
		level = service.RatingMemberLevel(user)
	}
	var list []map[string]interface{}
	cols := "id,site_type,dns"
	err := global.Db.Where("status = 1 and level = ?", level).
		Table("t_app_dns").
		Cols(cols).
		OrderBy("id desc").
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	for _, item := range list {
		item["dns"] = util.AesEncrypt(item["dns"].(string))
	}
	var result = make(map[string]interface{})
	result["list"] = list
	response.RespOk(c, "成功", result)
}

func ComboList(c *gin.Context) {
	var result = make(map[string]interface{})
	var list []map[string]interface{}
	cols := "g.id,g.m_type,g.title,g.price,g.period,g.dev_limit,g.flow_limit,g.is_discount,g.low,g.high"
	err := global.Db.Table("t_goods as g").Cols(cols).Where("status = 1").OrderBy("id desc").Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
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
	//	response.RespOk(c, "成功", result)
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
	response.RespOk(c, "成功", result)
}

func AdList(c *gin.Context) {
	var list []map[string]interface{}
	global.Db.Table("t_ad").Where("status = 1").Find(&list)
	result := make(map[string]interface{})
	result["list"] = list
	response.RespOk(c, "成功", result)
}

func UploadLog(c *gin.Context) {
	param := new(request.UploadLogRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	devId := c.Request.Header.Get("Dev-Id")
	dId, err := strconv.Atoi(devId)
	if err != nil {
		global.Logger.Err(err).Msg("设备鉴权失败")
		response.ResFail(c, "设备鉴权失败！")
		return
	}
	if !service.CheckUserDev(int64(dId), user) {
		global.Logger.Err(err).Msg("设备鉴权失败")
		response.ResFail(c, "设备鉴权失败！")
		return
	}
	log := &model.TUploadLog{
		UserId:    user.Id,
		DevId:     int64(dId),
		Content:   param.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err := global.Db.Insert(log)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("上传日志失败")
		response.ResFail(c, "上传日志失败！")
		return
	}
	response.ResOk(c, "成功")
}

func CreateOrder(c *gin.Context) {
	param := new(request.CreateOrderRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	goods := new(model.TGoods)
	has, err := global.Db.Where("id = ? and status = 1", param.GoodsId).Get(goods)
	if err != nil || !has {
		global.Logger.Err(err).Msg("创建订单失败")
		response.ResFail(c, "创建订单失败！")
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
		global.Logger.Err(err).Msg("创建订单失败")
		response.ResFail(c, "创建订单失败！")
		return
	}
	result := make(map[string]interface{})
	result["oid"] = id
	response.RespOk(c, "成功", result)
}

func OrderList(c *gin.Context) {
	param := new(request.OrderListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	session := service.OrderList(param, user)
	count, err := service.OrderList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "*"
	session.Cols(cols)
	session.OrderBy("o.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func DevList(c *gin.Context) {
	param := new(request.DevListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	session := service.UserDevList(param, user)
	count, err := service.UserDevList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "d.os,ud.*"
	session.Cols(cols)
	session.OrderBy("ud.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func BanDev(c *gin.Context) {
	param := new(request.BanDevRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	userDev := new(model.TUserDev)
	userDev.Status = constant.UserDevBanStatus
	userDev.UpdatedAt = time.Now()
	rows, err := global.Db.Cols("status", "updated_at").
		Where("user_id = ? and dev_id = ? and status = ?", user.Id, param.DevId, constant.UserDevNormalStatus).
		Update(userDev)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("踢除设备失败")
		response.ResFail(c, "踢除设备失败！")
		return
	}
	response.ResOk(c, "成功")
}

// ChangeNetwork 切换节点工作
func ChangeNetwork(c *gin.Context) {
	param := new(request.ChangeNetworkRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if param.WorkMode != 1 || param.WorkMode != 2 {
		global.Logger.Err(nil).Msg("参数错误")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token鉴权失败")
		response.ResFail(c, "用户鉴权失败！")
		return
	}
	devId := c.Request.Header.Get("Dev-Id")
	dId, err := strconv.Atoi(devId)
	if err != nil {
		global.Logger.Err(err).Msg("设备鉴权失败")
		response.ResFail(c, "设备鉴权失败！")
		return
	}
	if !service.CheckUserDev(int64(dId), user) {
		global.Logger.Err(err).Msg("设备鉴权失败")
		response.ResFail(c, "设备鉴权失败！")
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
		global.Logger.Err(err).Msg("workMode出错")
		sess.Rollback()
		response.RespFail(c, "失败", nil)
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
		global.Logger.Err(err).Msg("workLog出错")
		sess.Rollback()
		response.RespFail(c, "失败", nil)
		return
	}

	sess.Commit()
	//下发服务器配置给客户端
	result := make(map[string]interface{})
	result["node_id"] = 1
	response.RespOk(c, "成功", result)
}

func SwitchButtonStatus(c *gin.Context) {
	//param := new(request.SwitchButtonStatusRequest)
	//if err := c.ShouldBind(param); err != nil {
	//	global.Logger.Err(err).Msg("绑定参数")
	//	response.RespFail(c, lang.Translate("cn", "fail"), nil)
	//	return
	//}
	//claims := c.MustGet("claims").(*service.CustomClaims)
	//user, err := service.GetUserByClaims(claims)
	//if err != nil {
	//	global.Logger.Err(err).Msg("用户token鉴权失败")
	//	response.ResFail(c, "用户鉴权失败！")
	//	return
	//}
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

func commonPageListV2(page, size int, total int64, session *xorm.Session) (response.PageResult, error) {
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
		global.Logger.Err(err).Msg("查询出错！")
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
