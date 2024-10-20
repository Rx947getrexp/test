package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service"
	"net/http/httptest"
)

func main() {
	var (
		ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
		err    error
	)

	defer ctx.Done()

	var nodes []entity.TNode
	err = dao.TNode.Ctx(ctx).Where(do.TNode{Status: 1}).Scan(&nodes)
	if err != nil {
		glog.Fatalf(ctx, "get TNode failed, err: %s", err.Error())
	}

	for _, node := range nodes {
		resp, err := service.GetUserListFromNode(ctx, node.Ip)
		if err != nil {
			glog.Fatalf(ctx, "GetUserListFromNode failed, err: %s", err.Error())
		}

		for _, item := range resp.Items {
			err = addUserNode(ctx, item.Email, item.Password, node.Ip)
			if err != nil {
				glog.Fatalf(ctx, "addUserNode failed, err: %s", err.Error())
			}
		}
	}
}

func addUserNode(ctx *gin.Context, email, uuid, ip string) (err error) {
	var userNode *entity.TUserNode
	err = dao.TUserNode.Ctx(ctx).Where(do.TUserNode{
		Email:     email,
		V2RayUuid: uuid,
		Ip:        ip,
	}).Scan(&userNode)
	if err != nil {
		glog.Fatalf(ctx, "Get TUserNode failed, err: %s", err.Error())
	}

	if userNode == nil {
		var (
			lastInsertId int64
			userInfo     *entity.TUser
		)
		err = dao.TUser.Ctx(ctx).Where(do.TUser{
			Email:     email,
			V2RayUuid: uuid,
		}).Scan(&userInfo)
		if err != nil {
			glog.Fatalf(ctx, "GetUserListFromNode failed, err: %s", err.Error())
		}

		if userInfo == nil {
			glog.Warningf(ctx, "TUser is not existed, email: %s, uuid: %s, ip: %s", email, uuid, ip)
			return nil
		}
		if userInfo.Email != email || userInfo.V2RayUuid != uuid {
			glog.Fatalf(ctx, "userInfo is not match, email: %s/%s, uuid: %s/%s, ip: %s",
				email, userInfo.Email, uuid, userInfo.V2RayUuid, ip)
		}

		lastInsertId, err = dao.TUserNode.Ctx(ctx).Data(do.TUserNode{
			UserId:    userInfo.Id,
			Email:     email,
			Ip:        ip,
			V2RayUuid: uuid,
			Status:    1,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			glog.Fatalf(ctx, "insert TUserNode failed, err: %s", err.Error())
		}
		if lastInsertId < 1 {
			glog.Fatalf(ctx, "insert TUserNode failed, lastInsertId: %d", lastInsertId)
		}
	}
	return nil
}
