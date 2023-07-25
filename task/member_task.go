package task

import (
	"go-speed/global"
	"go-speed/model"
	"time"
)

// MemberLevelExpiredTask 会员等级到期任务（变更会员等级）
func MemberLevelExpiredTask() {
	global.Recovery()
	global.Logger.Info().Msg("MemberLevelExpiredTask start...")
	var err error
	for {
		nowTime := time.Now().Unix()
		afterDayTime := nowTime + 24*3600
		var list []*model.TUser
		err = global.Db.Where("v2ray_tag = 1 and expired_time <= ?", afterDayTime).OrderBy("expired_time asc").Find(&list)
		if err != nil {
			global.Logger.Err(err).Msg("MemberLevelExpiredTask err")
			time.Sleep(time.Second * 10)
			continue
		}
		if len(list) == 0 {
			time.Sleep(time.Minute)
			continue
		}
		for _, item := range list {
			for {
				if needUpdateUser(item) {
					break
				}
				time.Sleep(time.Second)
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func needUpdateUser(user *model.TUser) bool {
	currTimeSec := time.Now().Unix()
	if user.ExpiredTime <= currTimeSec {
		global.Logger.Info().Msgf("用户[%s]会员过期:", user.Uname)
		//远程调用执行v2ray服务器删除UUID；得到回执后，并修改用户的v2ray_tag字段为2
		return true
	}
	return false
}

// MemberSpeedExpiredTask 会员加速服务到期任务(删除节点UUID等)
func MemberSpeedExpiredTask() {
	global.Recovery()
	global.Logger.Info().Msg("MemberSpeedExpiredTask start...")
	var err error
	for {
		nowTime := time.Now().Unix()
		afterDayTime := nowTime + 24*3600
		var list []*model.TSuccessRecord
		err = global.Db.Where("end_time <= ? and status = 1", afterDayTime).OrderBy("end_time asc").Find(&list)
		if err != nil {
			global.Logger.Err(err).Msg("MemberSpeedExpiredTask err")
			time.Sleep(time.Second * 10)
			continue
		}
		if len(list) == 0 {
			time.Sleep(time.Minute)
			continue
		}
		for _, item := range list {
			for {
				if needUpdateLevelRecord(item) {
					break
				}
				time.Sleep(time.Second)
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func needUpdateLevelRecord(record *model.TSuccessRecord) bool {
	currTimeSec := time.Now().Unix()
	if record.EndTime <= currTimeSec {
		global.Logger.Info().Msgf("套餐记录[%d]会员过期:", record.Id)
		//执行过期流程
		return true
	}
	return false
}
