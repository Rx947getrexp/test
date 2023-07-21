package task

import (
	"fmt"
	"go-speed/global"
	"time"
)

// MemberLevelExpiredTask 会员等级到期任务（变更会员等级）
func MemberLevelExpiredTask() {
	global.Recovery()
	global.Logger.Info().Msg("OrderExpiredTask start...")
	for {
		nowTime := time.Now().Unix()
		fmt.Println(nowTime)
		time.Sleep(time.Second * 30)
	}
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
