package task

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/entity"
	"go-speed/service"
	"time"
)

const (
	leaderLockKeySyncOrderStatus       = "hs-fly-SyncPayOrderStatus-leader-lock"
	electionIntervalSyncPayOrderStatus = 1 * time.Minute
	lockTimeoutSyncPayOrderStatus      = electionIntervalSyncPayOrderStatus + 10*time.Minute
)

// SyncPayOrderStatus 同步订单支付状态
func SyncPayOrderStatus() {
	global.Recovery()
	global.Logger.Info().Msg("SyncPayOrderStatus start...")
	ctx := context.Background()
	for {
		isLeader, err := tryAcquireLock(ctx, leaderLockKeySyncOrderStatus, lockTimeoutSyncPayOrderStatus)
		if err != nil {
			global.Logger.Err(err).Msg("tryAcquireLock failed")
		} else if isLeader {
			global.Logger.Info().Msg("I am the leader")
			// 在这里执行主进程的逻辑
			doSyncPayOrderStatus()
			releaseLock(ctx, leaderLockKeySyncOrderStatus)
		} else {
			global.Logger.Info().Msg("I am a follower")
			// 在这里执行从进程的逻辑
		}
		time.Sleep(electionIntervalSyncPayOrderStatus)
	}
}

func doSyncPayOrderStatus() {
	var (
		ctx   = gctx.New()
		err   error
		items []entity.TPayOrder
	)
	err = dao.TPayOrder.Ctx(ctx).
		WhereNot(dao.TPayOrder.Columns().Status, constant.ParOrderStatusPaid).
		WhereGTE(dao.TPayOrder.Columns().CreatedAt, gtime.Now().Add(-7*time.Hour*24)).
		Order(dao.TPayOrder.Columns().Id, "DESC").
		Scan(&items)
	if err != nil {
		global.Logger.Err(err).Msg("query TPayOrder failed")
		return
	}

	for _, item := range items {
		_, err = service.SyncOrderStatus(&gin.Context{}, item.OrderNo)
		if err != nil {
			global.Logger.Err(err).Msgf("SyncOrderStatus failed, orderNo: %s", item.OrderNo)
			continue
		}
	}
}
