package node

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
)

func ListNodeForReport(c *gin.Context) {
	token := c.Request.Header.Get("Authorization-Token")
	if token != "" {
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("token出错")
			response.RespFail(c, i18n.RetMsgParamParseErr, nil)
			return
		}
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("用户token鉴权失败, claims: %+v", *claims)
			response.RespFail(c, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
			return
		}
		global.MyLogger(c).Info().Msgf("user: %s", user.Email)
	}

	var (
		err   error
		nodes []entity.TNode
	)

	err = dao.TNode.Ctx(c).Where(do.TNode{Status: 1}).Scan(&nodes)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("query nodes failed.")
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}

	items := make([]response.NodeItem, 0)
	for _, n := range nodes {
		items = append(items, response.NodeItem{Ip: n.Ip})
	}

	response.RespOk(c, i18n.RetMsgSuccess, response.ListNodeForReport{
		Items: items,
	})
}
