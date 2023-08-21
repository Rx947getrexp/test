package service

import (
	"errors"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/request"
	"math/rand"
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
	useCount, err := getUserDevs(user)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return errors.New("查询设备数出错")
	}
	limit, _ := getUserLimitDevs(user)
	//if err != nil {
	//	global.Logger.Err(err).Msg("数据库链接出错")
	//	return errors.New("查询设备限制数出错")
	//}
	userDev := new(model.TUserDev)
	has, err := global.Db.Where("dev_id = ?", devId).Get(userDev)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return errors.New("查询设备出错")
	}
	var rows int64
	if has {
		if userDev.Status == 2 {
			//如果是已踢的设备，需判断目前总活跃总数是否受限
			if useCount >= limit && user.V2rayUuid != "c541b521-17dd-11ee-bc4e-0c9d92c013fb" {
				global.Logger.Err(err).Msg("设备数超限制")
				return errors.New("设备数超限制")
			}
		}
		//更新
		userDev.UpdatedAt = time.Now()
		userDev.UserId = user.Id
		userDev.Status = constant.UserDevNormalStatus
		rows, err = global.Db.Cols("updated_at", "user_id", "status").Where("id = ?", userDev.Id).Update(userDev)
	} else {
		//根据等级判断，vip最多允许登录几台设备
		if useCount >= limit && user.V2rayUuid != "c541b521-17dd-11ee-bc4e-0c9d92c013fb" {
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

func getUserDevs(user *model.TUser) (int64, error) {
	var useCount int64
	_, err := global.Db.SQL("select count(id) from t_user_dev where user_id = ? and status = 1", user.Id).Get(&useCount)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错")
		return useCount, errors.New("查询出错")
	}
	return useCount, err
}

func getUserLimitDevs(user *model.TUser) (int64, error) {
	var err error
	var limit int64 = 2 //默认2
	if user.Level == 0 {
		return limit, nil
	}
	//取当前使用套餐的等级
	var result = make(map[string]interface{})
	has, err := global.Db.Table("t_success_record as r").
		Cols("g.id,g.dev_limit").
		Where("r.user_id = ? and r.status = 1", user.Id).
		Join("LEFT", "t_order as o", "r.order_id = o.id").
		Join("LEFT", "t_goods as g", "o.goods_id = g.id").
		Get(&result)
	if err != nil || !has {
		global.Logger.Err(err).Msg("查询设备数限制出错")
		return limit, err
	}
	limit = int64(result["dev_limit"].(int))
	return limit, err
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

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}
