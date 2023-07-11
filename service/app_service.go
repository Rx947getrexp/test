package service

import (
	"errors"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/request"
	"time"
	"xorm.io/xorm"
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

func TeamList(param *request.TeamListRequest, user *model.TUser) *xorm.Session {
	session := global.Db.Table("t_user_team as t")
	session.Where("t.direct_id = ?", user.Id)
	session.Join("LEFT", "t_user as u", "u.id = t.user_id")
	return session
}

func OrderList(param *request.OrderListRequest, user *model.TUser) *xorm.Session {
	session := global.Db.Table("t_order as o")
	session.Where("o.user_id = ?", user.Id)
	return session
}

func NoticeList(param *request.NoticeListRequest) *xorm.Session {
	session := global.Db.Table("t_notice as n")
	session.Where("n.status = 1")
	return session
}

func UserDevList(param *request.DevListRequest, user *model.TUser) *xorm.Session {
	session := global.Db.Table("t_user_dev as ud")
	session.Where("user_id = ? and status = 1", user.Id)
	session.Join("LEFT", "t_dev as d", "d.id = ud.dev_id")
	return session
}

func HasDev(devId int64) bool {
	dev := new(model.TDev)
	has, err := global.Db.Where("id = ?", devId).Get(dev)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return false
	}
	if !has {
		return false
	}
	return true
}

func CheckUserDev(devId int64, user *model.TUser) bool {
	if !HasDev(devId) {
		return false
	}
	userDev := new(model.TUserDev)
	has, err := global.Db.Where("dev_id = ? and user_id = ? and status = 1", devId, user.Id).Get(userDev)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return false
	}
	return has
}

func UpdateUserDev(devId int64, user *model.TUser) error {
	if !HasDev(devId) {
		return errors.New("该设备号不存在")
	}
	userDev := new(model.TUserDev)
	has, err := global.Db.Where("dev_id = ?", devId).Get(userDev)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return errors.New("数据库链接出错")
	}
	var rows int64
	if has {
		//更新
		userDev.UpdatedAt = time.Now()
		userDev.UserId = user.Id
		userDev.Status = constant.UserDevNormalStatus
		rows, err = global.Db.Cols("updated_at", "user_id", "status").Where("id = ?", userDev.Id).Update(userDev)
	} else {
		var useCount int64
		_, err := global.Db.SQL("select count(id) from t_user_dev where user_id = ? and status = 1", user.Id).Get(&useCount)
		if err != nil {
			global.Logger.Err(err).Msg("数据库链接出错")
			return errors.New("数据库链接出错")
		}
		//根据等级判断，vip最多允许登录几台设备
		if useCount >= 2 {
			global.Logger.Err(err).Msg("设备数超限制")
			return errors.New("设备数超限制")
		}
		//添加
		userDev.CreatedAt = time.Now()
		userDev.UpdatedAt = time.Now()
		userDev.UserId = user.Id
		userDev.DevId = devId
		userDev.Status = constant.UserDevNormalStatus
		rows, err = global.Db.Insert(userDev)
	}
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return errors.New("数据库链接出错")
	}
	if rows < 1 {
		global.Logger.Err(err).Msg("数据库操作出错")
		return errors.New("数据库操作出错")
	}
	return nil
}

func UpdateUserWorkMode(devId int64, user *model.TUser) error {
	workMode := new(model.TWorkMode)
	has, err := global.Db.Where("user_id = ? and dev_id = ?", user.Id, devId).Get(workMode)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return errors.New("数据库链接出错")
	}
	var rows int64
	if has {
		workMode.UserId = user.Id
		workMode.UpdatedAt = time.Now()
		rows, err = global.Db.Cols("user_id", "updated_at").Where("dev_id = ?", devId).Update(workMode)
	} else {
		workMode.UserId = user.Id
		workMode.DevId = devId
		workMode.CreatedAt = time.Now()
		workMode.UpdatedAt = time.Now()
		workMode.ModeType = 1
		rows, err = global.Db.Insert(workMode)
	}
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return errors.New("数据库链接出错")
	}
	if rows < 1 {
		global.Logger.Err(err).Msg("数据库操作出错")
		return errors.New("数据库操作出错")
	}
	return nil
}
