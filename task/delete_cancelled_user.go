package task

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"time"
)

func DeleteCancelledUser() {
	global.Recovery()
	global.Logger.Info().Msg("SyncPayOrderStatus start...")
	//ctx := context.Background()
	for {
		//isLeader, err := tryAcquireLock(ctx, leaderLockKeySyncOrderStatus, lockTimeoutSyncPayOrderStatus)
		//if err != nil {
		//	global.Logger.Err(err).Msg("tryAcquireLock failed")
		//} else if isLeader {
		//	global.Logger.Info().Msg("I am the leader")
		//	// 在这里执行主进程的逻辑
		//	doSyncPayOrderStatus()
		//	releaseLock(ctx, leaderLockKeySyncOrderStatus)
		//} else {
		//	global.Logger.Info().Msg("I am a follower")
		//	// 在这里执行从进程的逻辑
		//}
		doSyncPayOrderStatus()
		time.Sleep(electionIntervalSyncPayOrderStatus)
	}
}

func doDeleteCancelledUser() {
	//items, err := GetLatestCancelledUsers()
	//if err != nil {
	//	return
	//}

	//for _, item := range items {
	//	dao.TUser
	//}
}

func GetLatestCancelledUsers() (items []entity.TUserCancelled, err error) {
	err = dao.TUserCancelled.Ctx(context.Background()).
		WhereGTE(
			dao.TUserCancelled.Columns().UpdatedAt,
			do.TUserCancelled{
				UpdatedAt: gtime.Now().AddDate(0, 0, -7),
			}).Scan(&items)
	if err != nil {
		global.Logger.Err(err).Msgf("get cancelled user failed")
		return
	}
	return items, nil
}
