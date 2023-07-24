package task

import (
	"fmt"
	"go-speed/global"
	"go-speed/model"
	"time"
)

// MemberLevelExpiredTask 会员等级到期任务（变更会员等级）
func MemberLevelExpiredTask() {
	global.Recovery()
	global.Logger.Info().Msg("OrderExpiredTask start...")
	var err error
	for {
		nowTime := time.Now().Unix()
		afterDayTime := nowTime + 24*3600
		var list []*model.TUser
		err = global.Db.Where("v2ray_tag = 1 and expired_time <= ?", afterDayTime).OrderBy("expired_time asc").Find(&list)
		if err != nil {
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
	return false
}

// MemberSpeedExpiredTask 会员加速服务到期任务(删除节点UUID等)
func MemberSpeedExpiredTask() {
	global.Recovery()
	global.Logger.Info().Msg("MemberSpeedExpiredTask start...")
	for {
		nowTime := time.Now().Unix()
		fmt.Println(nowTime)
		time.Sleep(time.Second * 30)
	}
}
