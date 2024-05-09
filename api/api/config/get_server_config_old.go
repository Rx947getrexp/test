package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"net/http"
	"strconv"
	"strings"
)

func GetConfigOld(c *gin.Context) {
	param := new(request.BanDevRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败")
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(c).Info().Msgf(">>> param: %+v", *param)
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetUserByClaims(claims)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v", *claims)
		response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
		return
	}

	// 当参数node_id设置时，表示指定node_id来获取配置；
	// 当参数node_id没有设置时，表示获取全部配置
	sqlWhere := "status = 1"
	if param.NodeId > 0 {
		sqlWhere = fmt.Sprintf("id = %d and status = 1", param.NodeId)
	} else if user.Email == "ru100@qq.com" {
		sqlWhere = fmt.Sprintf("status = 1")
	} /*else {
		sqlWhere = fmt.Sprintf("id not in (100003) and status = 1")
	}*/

	uuid := user.V2rayUuid
	var list []map[string]interface{}
	cols := "id,name,title,title_en,country,country_en,server,port," +
		"min_port as min,max_port as max,path,is_recommend"
	errs := global.Db.Where(sqlWhere).
		Table("t_node").
		Cols(cols).
		OrderBy("id desc").
		Find(&list)
	if errs != nil {
		global.MyLogger(c).Err(errs).Msgf("数据库链接出错, email: %s", user.Email)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	//var dnsArray = []string{}
	var d_proxy = []string{}
	var d_data = []string{}
	var d_proto = []string{}
	i := 0
	for _, item := range list {

		nodeId := item["id"].(int64)
		nodePorts := []int{443, 13001, 13002, 13003, 13004, 13005}
		//nodePorts := make([]int64, 0)
		//nodePorts = append(nodePorts, item["port"].(int64))
		//for _no := item["min"].(int64); _no <= item["max"].(int64); _no++ {
		//	nodePorts = append(nodePorts, _no)
		//}
		global.MyLogger(c).Info().Msgf(">>>>> nodePorts: %+v", nodePorts)

		dnsList, _ := service.FindNodeDnsByNodeId(nodeId, user.Level+1) // user_level+1等于服务器域名的等级

		for _, dns := range dnsList {
			for _, nodePort := range nodePorts {
				i = i + 1
				mproxy := "\"proxy" + strconv.Itoa(i) + "\""

				//d_proxy = append(d_proxy, mproxy)
				m := fmt.Sprintf("{\"tag\": %s,\"protocol\": \"chain\",\"settings\": {\"actors\": [\"tls\",\"ws\",\"trojan%d\"]}}", mproxy, i)
				d_data = append(d_data, m)
				np := fmt.Sprintf("{\"protocol\": \"trojan\",\"settings\": {\"address\": \"%s\",\"port\": %d,\"password\": \"%s\"},\"tag\": \"trojan%d\"}", dns.Dns, nodePort, uuid, i)
				d_proto = append(d_proto, np)
				//name := fmt.Sprintf("trojan%d", i)
				//name := "trojan"
				//retstr := fmt.Sprintf("{\"protocol\": \"trojan\",\"settings\": {\"address\": \"%s\",\"port\": 443,\"password\": \"%s\"},\"tag\": \"%s\"}", dns.Dns, uuid, name)

				//dnsArray = append(dnsArray, retstr)

				mproxy = fmt.Sprintf("{\"password\": \"%s\",\"port\": %d,\"email\": \"\",\"level\": 0,\"flow\": \"\",\"address\": \"%s\"}", uuid, nodePort, dns.Dns)
				d_proxy = append(d_proxy, mproxy)
			}
		}
	}
	global.MyLogger(c).Info().Msgf(">>>>> d_proxy: %+v", d_proxy)
	//mystring := "{\"log\":{\"level\":\"{{logLevel}}\",\"output\":\"{{leafLogFile}}\"},\"dns\":{\"servers\":[\"1.1.1.1\",\"8.8.8.8\"],\"hosts\":{\"node2.wuwuwu360.xyz\":[\"107.148.239.239\"]}},\"inbounds\":[{\"protocol\":\"tun\",\"settings\":{\"fd\":\"{{tunFd}}\"},\"tag\":\"tun_in\"}],\"outbounds\":[{\"protocol\":\"failover\",\"tag\":\"failover_out\",\"settings\":{\"actors\":[%s],\"failTimeout\":4,\"healthCheck\":true,\"checkInterval\":300,\"failover\":true,\"fallbackCache\":false,\"cacheSize\":256,\"cacheTimeout\":60}},%s,{\"protocol\":\"tls\",\"tag\":\"tls\",\"settings\":{\"alpn\":[\"http/1.1\"],\"insecure\":true}},{\"protocol\":\"ws\",\"tag\":\"ws\",\"settings\":{\"path\":\"/work\"}},%s,{\"protocol\":\"direct\",\"tag\":\"direct_out\"},{\"protocol\":\"drop\",\"tag\":\"reject_out\"}],\"router\":{\"domainResolve\":true,\"rules\":[{\"external\":[\"site:{{dlcFile}}:cn\"],\"target\":\"direct_out\"},{\"external\":[\"mmdb:{{geoFile}}:cn\"],\"target\":\"direct_out\"},{\"domainKeyword\":[\"apple\",\"icloud\"],\"target\":\"direct_out\"}]}}"
	//mystring := "{\"routing\":{\"rules\":[{\"type\":\"field\",\"outboundTag\":\"direct\",\"domain\":[\"icloud\",\"apple\",\"geosite:private\",\"geosite:cn\"]},{\"ip\":[\"geoip:private\",\"geoip:cn\"],\"outboundTag\":\"direct\",\"type\":\"field\"}],\"domainMatcher\":\"hybrid\",\"domainStrategy\":\"AsIs\",\"balancers\":[]},\"log\":{\"loglevel\":\"warning\",\"dnsLog\":false},\"outbounds\":[{\"tag\":\"proxy\",\"mux\":{\"enabled\":false,\"concurrency\":50},\"protocol\":\"trojan\",\"streamSettings\":{\"wsSettings\":{\"path\":\"/work\",\"headers\":{\"host\":\"\"}},\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"allowInsecure\":true,\"fingerprint\":\"\"},\"security\":\"tls\",\"network\":\"ws\"},\"settings\":{\"servers\":[%s]}},{\"tag\":\"direct\",\"protocol\":\"freedom\"},{\"tag\":\"reject\",\"protocol\":\"blackhole\"}]}"
	mystring := "{\"routing\":{\"rules\":[{\"type\":\"field\",\"outboundTag\":\"direct\",\"domain\":[\"icloud\",\"apple\",\"geosite:private\"]},{\"ip\":[\"geoip:private\"],\"outboundTag\":\"direct\",\"type\":\"field\"}],\"domainMatcher\":\"hybrid\",\"domainStrategy\":\"AsIs\",\"balancers\":[]},\"log\":{\"loglevel\":\"warning\",\"dnsLog\":false},\"outbounds\":[{\"tag\":\"proxy\",\"mux\":{\"enabled\":false,\"concurrency\":50},\"protocol\":\"trojan\",\"streamSettings\":{\"wsSettings\":{\"path\":\"/work\",\"headers\":{\"host\":\"\"}},\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"allowInsecure\":true,\"fingerprint\":\"\"},\"security\":\"tls\",\"network\":\"ws\"},\"settings\":{\"servers\":[%s]}},{\"tag\":\"direct\",\"protocol\":\"freedom\"},{\"tag\":\"reject\",\"protocol\":\"blackhole\"}]}"
	if user.Email == "ru101@qq.com" {
		mystring = "{\"routing\":{\"rules\":[{\"type\":\"field\",\"outboundTag\":\"proxy\",\"domain\":[\"regexp:.*\"]}],\"domainMatcher\":\"hybrid\",\"domainStrategy\":\"AsIs\",\"balancers\":[]},\"log\":{\"loglevel\":\"warning\",\"dnsLog\":false},\"outbounds\":[{\"tag\":\"proxy\",\"mux\":{\"enabled\":false,\"concurrency\":50},\"protocol\":\"trojan\",\"streamSettings\":{\"wsSettings\":{\"path\":\"/work\",\"headers\":{\"host\":\"\"}},\"tlsSettings\":{\"alpn\":[\"http/1.1\"],\"allowInsecure\":true,\"fingerprint\":\"\"},\"security\":\"tls\",\"network\":\"ws\"},\"settings\":{\"servers\":[%s]}},{\"tag\":\"direct\",\"protocol\":\"freedom\"},{\"tag\":\"reject\",\"protocol\":\"blackhole\"}]}"
	}
	global.MyLogger(c).Info().Msgf(">>> get_conf >>> user.Email: %s, mystring: %+v", user.Email, mystring)
	global.MyLogger(c).Info().Msgf(">>> get_conf >>> user.Email: %s, return config: %+v", user.Email, fmt.Sprintf(mystring, strings.Join(d_proxy, ",")))
	c.String(http.StatusOK, fmt.Sprintf(mystring, strings.Join(d_proxy, ",")))

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
