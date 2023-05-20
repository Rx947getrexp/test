package service

import (
	"errors"
	"go-speed/global"
	"go-speed/model"
)

func GetUserByClaims(claims *CustomClaims) (*model.TUser, error) {
	uid := claims.UserId
	user := new(model.TUser)
	has, err := global.Db.Where("id = ? and status = 0", uid).Get(user)
	if err != nil {
		global.Logger.Err(err).Msg("查询用户出错！")
		return nil, err
	}
	if !has {
		return nil, errors.New("您已被风控，请联系客服！")
	}
	return user, nil
}
