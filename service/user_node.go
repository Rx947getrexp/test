package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"io/ioutil"
	"sync"
	"time"
)

const (
	UserNodeStatusInit  = 0
	UserNodeStatusAdded = 1

	V2rayConfigAdd    = "1"
	V2rayConfigDelete = "2"
)

var nodeUsers map[string]map[string]string
var nodeUsersRWMutex sync.Mutex

func init() {
	nodeUsers = make(map[string]map[string]string)
}

func getFileName(ip string) string {
	return fmt.Sprintf("/wwwroot/go/go-api/config/v2ray/users-%s.json", ip)
}

func isUserInConfig(ctx *gin.Context, userEmail, userUUID, ip string) (flag bool) {
	nodeUsersRWMutex.Lock()
	defer nodeUsersRWMutex.Unlock()

	users, ok := nodeUsers[ip]
	if !ok {
		nodeUsers[ip] = make(map[string]string)

		users = loadConfig(ctx, ip)
		if users != nil && len(users) > 0 {
			nodeUsers[ip] = users
		}
	}
	if users == nil || len(users) < 1 {
		return false
	}

	uuid, ok := users[userEmail]
	if ok && uuid == userUUID {
		return true
	}
	return false
}

func loadConfig(ctx *gin.Context, ip string) map[string]string {
	fileName := getFileName(ip)

	// 从文件中读取JSON数据
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("read config file failed, %s", fileName)
		return nil
	}

	// 定义一个map变量来存储数据
	var data map[string]string

	// 将JSON数据解码到map变量中
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("JSON unmarshaling failed, %s", fileName)
		return nil
	}
	return data
}

func addUserConfig(ctx *gin.Context, user *entity.TUser, node *entity.TNode) (err error) {
	url := fmt.Sprintf("http://%s:15003/node/add_sub", node.Ip)
	global.MyLogger(ctx).Info().Msgf(">>>>>>>>> url: %s", url)

	nodeAddSubRequest := &request.NodeAddSubRequest{
		Email: user.Email,
		Uuid:  user.V2RayUuid,
		Level: fmt.Sprintf("%d", user.Level),
		Tag:   V2rayConfigAdd,
	}
	timestamp := fmt.Sprint(time.Now().Unix())
	headerParam := make(map[string]string)
	res := new(response.Response)
	headerParam["timestamp"] = timestamp
	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
	err = util.HttpClientPostV2(url, headerParam, nodeAddSubRequest, res)
	if res != nil {
		global.MyLogger(ctx).Info().Msgf("nodeAddSubRequest: %+v, res: %+v", *nodeAddSubRequest, *res)
	} else {
		global.MyLogger(ctx).Warn().Msgf("nodeAddSubRequest: %+v, res: is nil", *nodeAddSubRequest)
	}
	if err != nil {
		return gerror.Wrap(err, "add_sub failed")
	}
	return nil
}

func saveToFile(ctx *gin.Context, ip string) {
	m, ok := nodeUsers[ip]
	if !ok {
		global.MyLogger(ctx).Warn().Msgf("map is not found key: %s", ip)
		return
	}
	if len(m) == 0 {
		return
	}

	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Marshal resp failed")
		return
	}

	// 将JSON数据写入到文件中
	err = ioutil.WriteFile(getFileName(ip), jsonData, 0644)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("WriteFile failed")
		return
	}
}

func AddUserConfigToNode(ctx *gin.Context, user *entity.TUser, node *entity.TNode) (err error) {
	if isUserInConfig(ctx, user.Email, user.V2RayUuid, node.Ip) {
		global.MyLogger(ctx).Debug().Msgf("already in config, skip call api to addsub")
		return nil
	}

	err = addUserConfig(ctx, user, node)
	if err != nil {
		return err
	}

	nodeUsersRWMutex.Lock()
	defer nodeUsersRWMutex.Unlock()
	nodeUsers[node.Ip][user.Email] = user.V2RayUuid
	saveToFile(ctx, node.Ip)
	return nil
}

//func doAddConfigBak() {
//
//var items []entity.TUserNode
//err = dao.TUserNode.Ctx(ctx).Where(do.TUserNode{
//	Email:     user.Email,
//	Ip:        node.Ip,
//	V2RayUuid: user.V2RayUuid,
//}).Scan(&items)
//if err != nil {
//	return gerror.Wrap(err, "get TUserNode failed")
//}
//
//var userNode *entity.TUserNode
//for i, item := range items {
//	if item.Email == user.Email && item.V2RayUuid == user.V2RayUuid {
//		userNode = &items[i]
//		break
//	}
//}
//
//if userNode == nil {
//	return doAddUserConfigToNode(ctx, user, node)
//} else if userNode.Status == UserNodeStatusAdded {
//	return doUpdateTime(ctx, userNode.Id)
//} else {
//	return doAddUserConfigToV2ray(ctx, user, node, userNode.Id)
//}
//}

func doUpdateTime(ctx *gin.Context, userNodeId int64) (err error) {
	_, e := dao.TUserNode.Ctx(ctx).Data(
		do.TUserNode{UpdatedAt: gtime.Now()}).Where(do.TUserNode{
		Id: userNodeId,
	}).Update()
	if e != nil {
		global.MyLogger(ctx).Err(e).Msgf("update TUserNode.updated_at failed")
	}
	return nil
}

func doAddUserConfigToNode(ctx *gin.Context, user *entity.TUser, node *entity.TNode) (err error) {
	var lastInsertId int64
	lastInsertId, err = dao.TUserNode.Ctx(ctx).Data(do.TUserNode{
		UserId:    user.Id,
		Email:     user.Email,
		Ip:        node.Ip,
		V2RayUuid: user.V2RayUuid,
		Status:    UserNodeStatusInit,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return gerror.Wrap(err, "insert TUserNode failed")
	}
	global.MyLogger(ctx).Info().Msgf("lastInsertId: %d", lastInsertId)
	return doAddUserConfigToV2ray(ctx, user, node, lastInsertId)
}

func doAddUserConfigToV2ray(ctx *gin.Context, user *entity.TUser, node *entity.TNode, userNodeId int64) (err error) {
	err = addUserConfig(ctx, user, node)
	if err != nil {
		return err
	}

	_, e := dao.TUserNode.Ctx(ctx).Data(
		do.TUserNode{Status: UserNodeStatusAdded, UpdatedAt: gtime.Now()}).Where(do.TUserNode{
		Id: userNodeId,
	}).Update()
	if e != nil {
		global.MyLogger(ctx).Err(e).Msgf("update TUserNode.staus = 1 failed")
	}
	return nil
}

func GetUserListFromNode(ctx *gin.Context, ip string) (resp response.GetUserListResponse, err error) {
	type userListReq struct{}
	var req userListReq

	url := fmt.Sprintf("http://%s:15003/node/get_user_list", ip)
	res := new(response.Response)
	timestamp := fmt.Sprint(time.Now().Unix())
	headerParam := make(map[string]string)
	headerParam["timestamp"] = timestamp
	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
	err = util.HttpClientPostV2(url, headerParam, req, res)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s http failed: %s", url, err.Error())
		return
	}

	if res.Code != response.Success {
		err = fmt.Errorf("Code: %d, Msg: %s ", res.Code, res.Msg)
		global.MyLogger(ctx).Err(err).Msgf("%s return code is not success: Code: %d, Msg: %s", url, res.Code, res.Msg)
		return
	}

	err = json.Unmarshal(res.Data, &resp)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s Unmarshal failed, Data: %s, err: %s", url, string(res.Data), err.Error())
		return
	}
	return resp, nil
}
