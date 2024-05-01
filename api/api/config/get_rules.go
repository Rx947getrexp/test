package config

import (
	"github.com/gin-gonic/gin"
	"go-speed/i18n"
	"go-speed/model/response"
)

type GetRulesReq struct {
	//UserId uint64 `form:"user_id" json:"user_id"`
}

type GetRulesRes struct {
	Ips     []string `json:"ips"`
	Domains []string `json:"domains"`
}

func GetRules(ctx *gin.Context) {
	var (
		//err error
		//req = new(GetRulesReq)
		res GetRulesRes
	)
	//if err = ctx.ShouldBind(req); err != nil {
	//	global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
	//	response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
	//	return
	//}
	//global.MyLogger(ctx).Info().Msgf(">>> req: %+v", *req)

	//_, err = common.CheckUserByUserId(ctx, req.UserId)
	//if err != nil {
	//	return
	//}
	res = GetRulesRes{
		Ips:     GenRuleIp(ctx, "", false),
		Domains: GenRuleDomain(ctx, "", false),
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, res)
}
