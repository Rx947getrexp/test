package task

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"net/http/httptest"
	"os"
	"time"
)

const (
	leaderLockKey    = "hs-fly-DeleteExpiredUser-leader-lock"
	electionInterval = 15 * time.Minute
	lockTimeout      = electionInterval + 10*time.Minute
)

// 踢掉已经过期的账号
func DeleteExpiredUser() {
	global.Recovery()
	global.Logger.Info().Msg("DeleteExpiredUser start...")
	//ctx := context.Background()
	for {
		//isLeader, err := tryAcquireLock(ctx, leaderLockKey, lockTimeout)
		//if err != nil {
		//	global.Logger.Err(err).Msg("tryAcquireLock failed")
		//} else if isLeader {
		//	global.Logger.Info().Msg("I am the leader")
		//	// 在这里执行主进程的逻辑
		//	DoDeleteExpiredUser()
		//	releaseLock(ctx, leaderLockKey)
		//} else {
		//	global.Logger.Info().Msg("I am a follower")
		//	// 在这里执行从进程的逻辑
		//}
		DoDeleteExpiredUser()
		time.Sleep(1 * time.Second)
	}
}

func DoDeleteExpiredUser() {
	var (
		err   error
		items []entity.TUser
	)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	err = dao.TUser.Ctx(ctx).
		WhereLTE(dao.TUser.Columns().ExpiredTime, time.Now().Add(time.Minute*1).Unix()).
		Where(do.TUser{Kicked: 0}).Limit(100).
		Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("get expired users failed")
		return
	}
	if len(items) == 0 {
		global.MyLogger(ctx).Info().Msg("no expired user")
		return
	}

	for _, item := range items {
		err = kickUser(ctx, &item)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("kickUser failed, user: %d/%s", item.Id, item.Email)
		}
	}

	global.MyLogger(ctx).Info().Msg("DoDeleteExpiredUser finished this time")
}

func kickUser(ctx *gin.Context, user *entity.TUser) (err error) {
	var userInfoNow *entity.TUser
	err = dao.TUser.Ctx(ctx).Where(do.TUser{
		Id:    user.Id,
		Email: user.Email,
	}).Scan(&userInfoNow)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query user now info failed, email: %s", user.Email)
		return
	}

	if !service.IsVIPExpired(userInfoNow) {
		global.MyLogger(ctx).Info().Msgf("user is not expired, (%d/%s) ", user.Id, user.Email)
		return
	}

	var affected int64
	affected, err = dao.TUser.Ctx(ctx).
		Where(do.TUser{
			Id:      user.Id,
			Email:   user.Email,
			Version: user.Version,
		}).
		Data(do.TUser{Level: 0}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Kick user update Kicked flag failed, email: %s", user.Email)
		return
	}
	if affected == 1 {
		global.MyLogger(ctx).Info().Msgf("userLevel update success, (%d/%s) ", user.Id, user.Email)
	} else {
		global.MyLogger(ctx).Info().Msgf("userLevel update not success, (%d/%s), affected=%d, but not 1", user.Id, user.Email, affected)
	}

	// 删除所有节点上的配置
	err = DeleteUserV2rayConfig(ctx, user.Email, user.V2RayUuid)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("DeleteUser failed, email: %s", user.Email)
		return
	}
	affected, err = dao.TUser.Ctx(ctx).
		Where(do.TUser{
			Id:      user.Id,
			Email:   user.Email,
			Version: user.Version,
		}).
		Data(do.TUser{
			Level:        0,
			Kicked:       1,
			LastKickedAt: gtime.Now(),
			Version:      user.Version + 1,
		}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Kick user update Kicked flag failed, email: %s", user.Email)
		return
	}
	if affected == 1 {
		global.MyLogger(ctx).Info().Msgf("Kick user success, (%d/%s) ", user.Id, user.Email)
	} else {
		global.MyLogger(ctx).Info().Msgf("Kick user not success, (%d/%s), affected=%d, but not 1", user.Id, user.Email, affected)
	}
	return nil
}

func DeleteUser(ctx *gin.Context, user *model.TUser) error {
	req := &request.NodeAddSubRequest{
		Email: user.Email,
		Uuid:  user.V2rayUuid,
		Level: fmt.Sprintf("%d", user.Level),
		Tag:   "2", // TODO：删除用户
	}

	// TODO：白名单逻辑
	var (
		countryEntities []entity.TServingCountry
		nodeEntities    []entity.TNode
		countryNames    []string
	)
	err := dao.TServingCountry.Ctx(ctx).Where(do.TServingCountry{Status: 1}).Scan(&countryEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("get TServingCountry failed.")
		return err
	}
	for _, s := range countryEntities {
		if s.IsFree == constant.IsFreeSiteNo {
			countryNames = append(countryNames)
		}
	}
	err = dao.TNode.Ctx(ctx).
		Where(do.TNode{Status: 1}).
		WhereIn(dao.TNode.Columns().CountryEn, countryNames).
		Scan(&nodeEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("get TNode failed.")
		return err
	}

	//nodeList, _ := service.FindNodes(user.Level + 1)
	for _, node := range nodeEntities {
		//url := fmt.Sprintf("https://%s/site-api/node/add_sub", node.Server)
		//if strings.Contains(node.Server, "http") {
		//	url = fmt.Sprintf("%s/node/add_sub", node.Server)
		//}
		url := fmt.Sprintf("http://%s:15003/node/add_sub", node.Ip)

		timestamp := fmt.Sprint(time.Now().Unix())
		headerParam := make(map[string]string)
		res := new(response.Response)
		headerParam["timestamp"] = timestamp
		headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))

		global.Logger.Info().Msgf("delete-user-req, url: %s, req: %s", url, gjson.MustEncode(req))
		err = util.HttpClientPostV2(url, headerParam, req, res)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("delete-user-failed. email: %s, uuid: %s, ip: %s", req.Email, req.Uuid, node.Ip)
			return err
		}
		global.MyLogger(ctx).Info().Msgf("delete-user-success. email: %s, uuid: %s, ip: %s", req.Email, req.Uuid, node.Ip)
	}
	return nil
}

func DeleteUserV2rayConfig(ctx *gin.Context, email, uuid string) error {
	var nodeEntities []entity.TNode
	err := dao.TNode.Ctx(ctx).
		Where(do.TNode{Status: 1}).
		Scan(&nodeEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("get TNode failed.")
		return err
	}

	for _, node := range nodeEntities {
		err = service.DeleteUserConfigForNode(ctx, email, uuid, node.Ip)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("delete-user-failed. email: %s, uuid: %s, ip: %s", email, uuid, node.Ip)
			return err
		}
		global.MyLogger(ctx).Info().Msgf("delete-user-success. email: %s, uuid: %s, ip: %s", email, uuid, node.Ip)
	}
	return nil
}

func UpdateUserStatus(user *model.TUser) error {
	userUpdate := &model.TUser{
		Status: 1,
	}
	rows, err := global.Db.Cols("status").Where("id = ?", user.Id).Update(userUpdate)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("添加user-uuid出错")
		return fmt.Errorf("update t_user表，先别说了")
	}
	return nil
}

func tryAcquireLock(ctx context.Context, leaderLockKey string, lockTimeout time.Duration) (bool, error) {
	// 尝试获取锁
	ok, err := global.Redis.SetNX(ctx, leaderLockKey, os.Getpid(), lockTimeout).Result()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	// 成功获取锁，更新锁的超时时间
	go func() {
		for {
			time.Sleep(lockTimeout / 2)
			global.Redis.Expire(ctx, leaderLockKey, lockTimeout)
		}
	}()
	return true, nil
}

func releaseLock(ctx context.Context, leaderLockKey string) {
	_, _ = global.Redis.Del(ctx, leaderLockKey).Result()
}
