package service

import (
	"fmt"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

func GetAdminUserByClaims(claims *CustomClaims) (*model.AdminUser, error) {
	uid := claims.UserId
	user := new(model.AdminUser)
	has, err := global.Db.Where("id = ? and status = 0 and is_del = 0", uid).Get(user)
	if err != nil {
		global.Logger.Err(err).Msg("查询用户出错！")
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不合法！")
	}
	return user, nil
}

// PostLoginAdmin 登陆后台
func PostLoginAdmin(param *request.LoginAdminRequest, ipAddr string) (*response.LoginAdminParam, error) {
	var dataParam = new(response.LoginAdminParam)
	userName := param.UserName
	password := param.Pass
	password = util.AesDecrypt(password)
	pwd := util.MD5(password)
	fmt.Println(userName)
	fmt.Println(pwd)
	adminUser := new(model.AdminUser)
	has, err := global.Db.Where("uname = ? and passwd = ? and is_del = 0 and status = 0", userName, pwd).Get(adminUser)
	if err != nil {
		global.Logger.Err(err).Msgf("登录出错！%s", userName)
		return nil, errors.New("登录出错！")
	}
	if !has {
		return nil, errors.New("用户名或密码不正确！")
	}
	if adminUser.Status != 0 {
		return nil, errors.New("用户已被冻结！")
	}
	dataParam.UserId = int64(adminUser.Id)
	dataParam.Token = GenerateTokenByUser(dataParam.UserId, AdminUserType)
	return dataParam, nil
}

func AccountAdminList(param *request.AccountListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table(model.AdminUser{})
	session.Where("admin_user.is_del = 0")
	if param.NickName != "" {
		session.Where("admin_user.nickname like ?", "%"+param.NickName+"%")
	}
	if param.Account != "" {
		session.Where("admin_user.uname like ?", "%"+param.Account+"%")
	}
	return session
}

func RoleAdminList(param *request.RoleListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table(model.AdminRole{})
	session.Where("admin_role.is_del = 0")
	if param.RoleName != "" {
		session.Where("admin_role.name like ?", "%"+param.RoleName+"%")
	}
	return session
}

// GenSnowflake 雪花算法生成id
func GenSnowflake() (int64, error) {
	return util.GenSnowflake(1)
}

func AddLog(c *gin.Context, uid int64) error {
	dateStr := time.Now().Format("2006-01-02")
	userLogs := &model.UserLogs{
		Id:        0,
		UserId:    uid,
		Datestr:   dateStr,
		Ip:        c.ClientIP(),
		UserAgent: c.GetHeader("User-Agent"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	sql := "insert ignore into user_logs(user_id,datestr,ip,user_agent,created_at,updated_at) values (?,?,?,?,?,?)"
	_, err := global.Db.Exec(sql, userLogs.UserId, userLogs.Datestr, userLogs.Ip, userLogs.UserAgent, userLogs.CreatedAt, userLogs.UpdatedAt)
	if err != nil {
		global.Logger.Err(err).Msg("db出错")
	}
	return err
}

func UpdateModel(T interface{}, sess *xorm.Session, id int64, columns ...string) error {
	rows, err := sess.Cols(columns...).Where("id = ?", id).Update(T)
	if err != nil {
		return err
	}
	if rows < 1 {
		return errors.New("更新记录不成功")
	}
	return nil
}

func FindNodeDnsByNodeId(nodeId int64, level int) ([]*model.TNodeDns, error) {
	var err error
	var list []*model.TNodeDns
	err = global.Db.Where("node_id = ? and status = 1 and level = ?", nodeId, level).Find(&list)
	return list, err
}

func GetAllRealMachineNodeDns() ([]*model.TNodeDns, error) {
	var err error
	var list []*model.TNodeDns
	err = global.Db.Where("status = 1 and is_machine = 1").Find(&list)
	return list, err
}

func GetAllNodes() ([]*model.TNode, error) {
	var err error
	var list []*model.TNode
	err = global.Db.Where(" status = 1").Find(&list)
	return list, err
}

func FindNodes(level int) ([]*model.TNode, error) {
	var err error
	var list []*model.TNode
	err = global.Db.Where(" status = 1 and is_recommend = ?", level).Find(&list)
	return list, err
}
func FindExpireUsers() ([]*model.TUser, error) {
	var err error
	var list []*model.TUser
	t1 := fmt.Sprint(time.Now().Unix())
	fmt.Printf(t1)
	err = global.Db.Where(" expired_time+3600 <= ? and expired_time+86400 >= ? ", t1, t1).Find(&list)
	fmt.Printf(t1)
	return list, err
}

func FindNodeDnsByLevel(level int) ([]*model.TNodeDns, error) {
	var err error
	var list []*model.TNodeDns
	err = global.Db.Where(" status = 1 and level = ?", level).Find(&list)
	return list, err
}

func FindAppDns(level int) ([]*model.TAppDns, error) {
	var err error
	var list []*model.TAppDns
	err = global.Db.Where("status = 1 and level = ?", level).Find(&list)
	return list, err
}

// RatingMemberLevel 用户评级获取线路
func RatingMemberLevel(user *model.TUser) int {

	return 1
}
