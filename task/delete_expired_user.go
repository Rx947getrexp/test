package task

import (
	"context"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"os"
	"strings"
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
	ctx := context.Background()
	for {
		isLeader, err := tryAcquireLock(ctx, leaderLockKey, lockTimeout)
		if err != nil {
			global.Logger.Err(err).Msg("tryAcquireLock failed")
		} else if isLeader {
			global.Logger.Info().Msg("I am the leader")
			// 在这里执行主进程的逻辑
			DoDeleteExpiredUser()
			releaseLock(ctx, leaderLockKey)
		} else {
			global.Logger.Info().Msg("I am a follower")
			// 在这里执行从进程的逻辑
		}
		time.Sleep(electionInterval)
	}
}

func DoDeleteExpiredUser() {
	var (
		err  error
		list []*model.TUser
	)

	//whiteList := []string{"zzz@qq.com"}
	// TODO：后续要注意时区
	nowTime := time.Now().Unix() // 过期30分钟后才执行踢人
	err = global.Db.Where("email != 'zzz@qq.com' and status = 0 and expired_time <= ? and expired_time > ?", nowTime, nowTime-24*60*60).OrderBy("expired_time asc").Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("get expired users failed")
		time.Sleep(time.Second * 10)
		return
	}

	if len(list) == 0 {
		global.Logger.Info().Msg("no expired user")
		time.Sleep(time.Minute)
		return
	}

	for _, user := range list {
		if err = DeleteUser(user); err != nil {
			global.Logger.Err(err).Msg("delete expired users from v2ray failed")
			continue
		}
		global.Logger.Err(err).Msgf("delete expired(%d) user(%s) from v2ray success", user.ExpiredTime, user.Email)

		//if err = UpdateUserStatus(user); err != nil {
		//	global.Logger.Err(err).Msg("update user status failed")
		//}
		//global.Logger.Err(err).Msgf("update expired(%d) user(%s) status success", user.ExpiredTime, user.Email)
	}
}

func DeleteUser(user *model.TUser) error {
	req := &request.NodeAddSubRequest{
		Email: user.Email,
		Uuid:  user.V2rayUuid,
		Level: fmt.Sprintf("%d", user.Level),
		Tag:   "2", // TODO：删除用户
	}

	// TODO：白名单逻辑
	nodeList, _ := service.FindNodes(user.Level + 1)
	for _, node := range nodeList {
		url := fmt.Sprintf("https://%s/site-api/node/add_sub", node.Server)
		if strings.Contains(node.Server, "http") {
			url = fmt.Sprintf("%s/node/add_sub", node.Server)
		}

		timestamp := fmt.Sprint(time.Now().Unix())
		headerParam := make(map[string]string)
		res := new(response.Response)
		headerParam["timestamp"] = timestamp
		headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))

		global.Logger.Info().Msgf("level:%d,req.Tag:%s,udid:%s,email:%s,url:%s,level:%s",
			user.Level, req.Tag, req.Uuid, req.Email, url, req.Level)
		err := util.HttpClientPostV2(url, headerParam, req, res)
		if err != nil {
			global.Logger.Err(err).Msg("delete expired user failed.")
			return err
		}
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
