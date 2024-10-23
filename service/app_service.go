package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/do"
	"go-speed/model/entity"
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
		return nil, fmt.Errorf("user-not-found, uid:%d", uid)
	}
	return user, nil
}

func GetUserByUserName(userName string) (*model.TUser, error) {
	user := new(model.TUser)
	has, err := global.Db.Where("uname = ?", userName).Get(user)
	if err != nil {
		global.Logger.Err(err).Msg("查询用户出错！")
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func CountUserByUserName(uname string) (int64, error) {
	var counts int64
	_, err := global.Db.SQL("select count(*) from t_user where uname = ?", uname).Get(&counts)
	if err != nil {
		global.Logger.Err(err).Msg("查询用户出错！")
		return 0, err
	}
	return counts, nil
}

func CountUserByUserNameAndClientId(uname, clientId string) (int64, error) {
	var counts int64
	_, err := global.Db.SQL("select count(*) from t_user where uname = ?", uname).Get(&counts)
	if err != nil {
		global.Logger.Err(err).Msg("查询用户出错！")
		return 0, err
	}
	return counts, nil
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
			if useCount >= limit {
				global.Logger.Err(err).Msg("设备数超限制")
				return errors.New("设备数超限制")
			}
		}
		/*
			if useCount >= limit {
				global.Logger.Err(err).Msg("设备数超限制")
				return errors.New("设备数超限制")
			}*/
		//更新
		userDev.UpdatedAt = time.Now()
		userDev.UserId = user.Id
		userDev.Status = constant.UserDevNormalStatus
		rows, err = global.Db.Cols("updated_at", "user_id", "status").Where("id = ?", userDev.Id).Update(userDev)
	} else {
		//根据等级判断，vip最多允许登录几台设备
		if useCount >= limit {
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

// 检查登录设备是否到达上限
func CheckDevNumLimits(ctx *gin.Context, devId int64, user *model.TUser) (bool, error) {
	// 登录设备数量上限，根据用户等级来定义
	// TODO：先简单处理，目前还没有产品形态定义
	limits := 6
	switch user.Level {
	case 0:
		limits = 6
	case 1:
		limits = 8
	case 2:
		limits = 10
	case 3:
		limits = 20
	default:
		limits = 50
	}

	// 30天内登录过的设备数量
	startTime := time.Now().Add(-1 * time.Hour * 30 * 24).Format("2006-01-02 15:04:05")

	var list []map[string]interface{}
	cols := "id,user_id,dev_id,status,created_at,updated_at,comment"
	err := global.Db.Where("user_id = ? and status = 1 and updated_at >= ?", user.Id, startTime).
		Table("t_user_dev").
		Cols(cols).
		OrderBy("id asc").
		Find(&list)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("数据库链接出错")
		return false, err
	}
	devCount := 0
	for _, item := range list {
		dId := item["dev_id"].(int64)
		if devId != dId {
			devCount = devCount + 1
		}
	}
	global.MyLogger(ctx).Info().Msgf("account: %s, level: %d, devCount: %d, limits: %d", user.Email, user.Level, devCount, limits)
	if devCount >= limits {
		global.MyLogger(ctx).Info().Msgf("dev limits error, account: %s, level: %d, devCount: %d, limits: %d",
			user.Email, user.Level, devCount, limits)
		return true, nil
	}

	var record *entity.TUserDev
	err = dao.TUserDev.Ctx(ctx).Where(do.TUserDev{
		UserId: user.Id,
		DevId:  devId,
	}).Scan(&record)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get TUserDev failed, email: %s", user.Email)
		return false, err
	}
	global.MyLogger(ctx).Debug().Msgf("TUserDev: %+v", *record)

	if record == nil {
		lastId, err := dao.TUserDev.Ctx(ctx).Data(do.TUserDev{
			UserId:    user.Id,
			DevId:     devId,
			Status:    constant.UserDevNormalStatus,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
			Comment:   "check&ban",
		}).InsertAndGetId()
		if err != nil {
			return false, gerror.Wrap(err, "insert TUserDev failed")
		}
		global.MyLogger(ctx).Debug().Msgf("insert TUserDev lastId: %d", lastId)
	} else {
		affect, err := dao.TUserDev.Ctx(ctx).Data(do.TUserDev{
			Status:    constant.UserDevBanStatus,
			UpdatedAt: gtime.Now(),
		}).Where(do.TUserDev{
			UserId: user.Id,
			DevId:  devId,
		}).UpdateAndGetAffected()
		if err != nil {
			return false, gerror.Wrap(err, "update TUserDev failed")
		}
		global.MyLogger(ctx).Debug().Msgf("update TUserDev affect: %d", affect)
	}
	return false, nil

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
	/*
		if user.Level < 2 {
			return limit, nil
		}
	*/
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

func GetUserConfig(userId int64) (*model.TUserConfig, error) {
	userConfig := new(model.TUserConfig)
	has, err := global.Db.Where("user_id = ?", userId).Get(userConfig)
	if err != nil {
		global.Logger.Err(err).Msg("查询数据库出错！")
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userConfig, nil
}

func CreateUserConfig(userId, nodeId int64) error {
	config := &model.TUserConfig{
		UserId:    userId,
		NodeId:    nodeId,
		Status:    constant.UserConfigStatusNormal,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rows, err := global.Db.Insert(config)
	if err != nil {
		global.Logger.Err(err).Msg("创建用户配置记录失败")
		return err
	}
	if rows != 1 {
		return fmt.Errorf("创建用户配置记录失败，rows:%d", rows)
	}
	return nil
}

func UpdateUserConfig(userId, nodeId int64) error {
	config := &model.TUserConfig{
		NodeId:    nodeId,
		Status:    constant.UserConfigStatusNormal,
		UpdatedAt: time.Now(),
	}
	rows, err := global.Db.Where("user_id = ?", userId).Update(config)
	if err != nil {
		global.Logger.Err(err).Msg("更新用户配置记录失败")
		return err
	}
	if rows != 1 {
		return fmt.Errorf("更新用户配置记录失败，rows:%d", rows)
	}
	return nil
}

//func GetLatestUserConfig(userId, devId int64) (*model.TUserConfig, error) {
//	userConfig := new(model.TUserConfig)
//	var ss *xorm.Session
//	if devId == 0 {
//		ss = global.Db.Where("user_id = ? and status = ?", userId, constant.UserConfigStatusNormal)
//	} else {
//		ss = global.Db.Where("user_id = ? and dev_id = ? and status = ?", userId, devId, constant.UserConfigStatusNormal)
//	}
//	has, err := ss.OrderBy("id desc").Get(userConfig)
//	if err != nil {
//		global.Logger.Err(err).Msg("查询数据库出错！")
//		return nil, err
//	}
//	if !has {
//		return nil, nil
//	}
//	return userConfig, nil
//}

func DeleteUserConfig(userId int64) error {
	config := &model.TUserConfig{
		Status:    constant.UserConfigStatusDeleted,
		UpdatedAt: time.Now(),
	}
	rows, err := global.Db.Where("user_id = ?", userId).Update(config)
	if err != nil {
		global.Logger.Err(err).Msg("删除用户配置记录失败")
		return err
	}
	if rows != 1 {
		return fmt.Errorf("删除用户配置记录失败，rows:%d", rows)
	}
	return nil
}
