package remote

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/types/api"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service/api/speed_api"
)

func GetUserEmailByUserId(ctx *gin.Context, userId uint64) (email string, err error) {
	var userEntity *entity.TUser
	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: userId}).Scan(&userEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("userId %d 查询db失败", userId)
		return
	}

	if userEntity != nil {
		return userEntity.Email, nil
	}

	resp, err := speed_api.DescribeUserInfo(ctx, &api.DescribeUserInfoReq{UserId: userId})
	if err != nil {
		return
	}
	if resp != nil {
		return resp.Email, nil
	}
	return "", nil
}
