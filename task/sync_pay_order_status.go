package task

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service"
	"net/http/httptest"
	"time"
)

const (
	leaderLockKeySyncOrderStatus       = "hs-fly-SyncPayOrderStatus-leader-lock"
	electionIntervalSyncPayOrderStatus = 30 * time.Second
	lockTimeoutSyncPayOrderStatus      = electionIntervalSyncPayOrderStatus + 20*time.Second
)

// SyncPayOrderStatus 同步订单支付状态
func SyncPayOrderStatus() {
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

func doSyncPayOrderStatus() {
	var (
		ctx   = gctx.New()
		err   error
		items []entity.TPayOrder
	)
	err = dao.TPayOrder.Ctx(ctx).
		Where(do.TPayOrder{
			PaymentChannelId: []string{
				constant.PayChannelUPay, constant.PayChannelWebMoneyPay,
				constant.PayChannelFreekassa_7, constant.PayChannelFreekassa_12,
				constant.PayChannelFreekassa_36, constant.PayChannelFreekassa_43,
				constant.PayChannelFreekassa_44,
			},
			Status: []string{constant.ParOrderStatusInit, constant.ParOrderStatusUnpaid},
		}).
		//Where(do.TPayOrder{OrderNo: "100701092254247"}).
		WhereGTE(dao.TPayOrder.Columns().CreatedAt, gtime.Now().Add(-30*time.Minute)).
		Order(dao.TPayOrder.Columns().Id, "DESC").Limit(5000).
		Scan(&items)
	if err != nil {
		global.Logger.Err(err).Msg("query TPayOrder failed")
		return
	}
	global.Logger.Info().Msgf("len(items): %d", len(items))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer c.Done()
	for _, item := range items {
		_, err = service.SyncOrderStatus(c, item.OrderNo, nil)
		if err != nil {
			global.Logger.Err(err).Msgf("SyncOrderStatus failed, orderNo: %s", item.OrderNo)
			continue
		}
	}
}
