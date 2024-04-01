package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/lang"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xorm.io/xorm"
)

// 登陆后台
func LoginAdmin(c *gin.Context) {
	param := new(request.LoginAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	reqIP := c.ClientIP()
	dataParam, err := service.PostLoginAdmin(param, reqIP)
	if err != nil {
		global.Logger.Err(err).Send()
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	response.RespOk(c, lang.Translate("cn", "success"), dataParam)
}

// GenerateAuth2Key 生成谷歌验证器私钥
func GenerateAuth2Key(c *gin.Context) {
	key := util.NewGoogleAuth().GetSecret()
	res := &response.GenerateAuth2KeyAdminResponse{
		Auth2Key: key,
	}
	response.RespOk(c, lang.Translate("cn", "success"), res)
}

// SetAuth2Key 设置谷歌验证器
func SetAuth2Key(c *gin.Context) {
	param := new(request.ChangeAuth2KeyAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}

	ok := util.VerifyAuthenticator(util.AesDecrypt(param.Auth2Key), param.Auth2Code)
	if !ok {
		response.ResFail(c, "谷歌验证码错误！")
		return
	}

	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("用户token非法")
		response.ResFail(c, "用户token非法！")
		return
	}

	user.UpdatedAt = time.Now()
	user.Authkey = param.Auth2Key
	user.IsReset = 0
	rows, err := global.Db.Cols("updated_at", "authkey", "is_reset").Where("id = ? and is_reset = 1", user.Id).Update(user)
	if err != nil {
		global.Logger.Err(err)
		response.RespFail(c, lang.Translate("cn", "fail"), nil)
		return
	}
	if rows < 1 {
		response.RespFail(c, "重置状态才能操作！", nil)
		return
	}
	response.ResOk(c, "成功")
}

func EditPasswd(c *gin.Context) {
	param := new(request.EditPasswdRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	if param.NewPass != param.EnterPass {
		response.ResFail(c, "两次密码不一致")
		return
	}
	oldPwdDecode := util.AesDecrypt(param.OldPass)
	oldPwd := util.MD5(oldPwdDecode)
	pwdDecode := util.AesDecrypt(param.NewPass)
	pwd := util.MD5(pwdDecode)
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser := new(model.AdminUser)
	adminUser.Passwd = pwd
	adminUser.UpdatedAt = time.Now()
	rows, err := global.Db.Cols("passwd", "updated_at").Where("id = ? and passwd = ?", claims.UserId, oldPwd).Update(adminUser)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("修改密码失败")
		response.ResFail(c, "修改密码失败！")
		return
	}
	response.ResOk(c, "成功")
}

// AddResource 添加资源
func AddResource(c *gin.Context) {
	param := new(request.AddResourceRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser := new(model.AdminUser)
	has, err := global.Db.Where("id = ?", claims.UserId).Get(adminUser)
	if err != nil || !has {
		global.Logger.Err(err).Msg("添加资源失败")
		response.ResFail(c, "添加资源失败！")
		return
	}
	adminRes := new(model.AdminRes)
	adminRes.Name = param.Name
	adminRes.ResType = param.ResType
	adminRes.CreatedAt = time.Now()
	adminRes.Pid = param.Pid
	adminRes.IsDel = 0
	adminRes.Sort = 0
	adminRes.UpdatedAt = time.Now()
	adminRes.Author = adminUser.Uname
	adminRes.Url = param.Url
	if param.Icon != "" {
		adminRes.Icon = param.Icon
	} else {
		adminRes.Icon = "example"
	}

	rows, err := global.Db.Insert(adminRes)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("添加资源失败")
		response.ResFail(c, "添加资源失败！")
		return
	}
	service.UpdateCachePermission()
	response.ResOk(c, "成功")
}

// DelResource 删除资源
func DelResource(c *gin.Context) {
	param := new(request.DelResourceRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser := new(model.AdminUser)
	has, err := global.Db.Where("id = ?", claims.UserId).Get(adminUser)
	if err != nil || !has {
		global.Logger.Err(err).Msg("添加资源失败")
		response.ResFail(c, "添加资源失败！")
		return
	}
	adminRes := new(model.AdminRes)
	adminRes.IsDel = 1
	adminRes.UpdatedAt = time.Now()
	adminRes.Author = adminUser.Uname
	rows, err := global.Db.Cols("is_del", "updated_at", "author").Where("id = ? and is_del = 0 ", param.Id).Update(adminRes)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("删除资源失败")
		response.ResFail(c, "删除资源失败！")
		return
	}
	service.UpdateCachePermission()
	response.ResOk(c, "成功")
}

// EditResource 修改资源
func EditResource(c *gin.Context) {
	param := new(request.EditResourceRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser := new(model.AdminUser)
	has, err := global.Db.Where("id = ?", claims.UserId).Get(adminUser)
	if err != nil || !has {
		global.Logger.Err(err).Msg("添加资源失败")
		response.ResFail(c, "添加资源失败！")
		return
	}
	adminRes := new(model.AdminRes)
	adminRes.Pid = param.Pid
	adminRes.ResType = param.ResType
	if param.Icon != "" {
		adminRes.Icon = param.Icon
	} else {
		adminRes.Icon = "example"
	}
	adminRes.Url = param.Url
	adminRes.Name = param.Name
	adminRes.UpdatedAt = time.Now()
	adminRes.Author = adminUser.Uname
	_, err = global.Db.Cols("pid", "res_type", "url", "name", "updated_at", "author", "icon").Where("id = ? ", param.Id).Update(adminRes)
	if err != nil {
		global.Logger.Err(err).Msg("修改资源失败")
		response.ResFail(c, "修改资源失败！")
		return
	}
	service.UpdateCachePermission()
	response.ResOk(c, "成功")
}

func AddAdminUser(c *gin.Context) {
	param := new(request.AccountAddAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}

	sess := global.Db.NewSession()
	sess.Begin()
	defer sess.Close()

	user := new(model.AdminUser)
	user.Uname = param.Account
	user.Nickname = param.NickName
	pwd := util.AesDecrypt(param.Password)
	user.Passwd = util.MD5(pwd)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsReset = 1
	rows, err := sess.Insert(user)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("添加用户失败")
		sess.Rollback()
		response.ResFail(c, "添加用户失败！")
		return
	}

	userRole := new(model.AdminUserRole)
	userRole.Uid = user.Id
	userRole.RoleId = param.RoleId
	userRole.CreatedAt = time.Now()
	userRole.UpdatedAt = time.Now()
	userRole.Author = adminUser.Nickname
	rows, err = sess.Insert(userRole)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("添加用户失败")
		sess.Rollback()
		response.ResFail(c, "添加用户失败！")
		return
	}

	sess.Commit()
	response.ResOk(c, "成功")
}

func EditAdminUser(c *gin.Context) {
	param := new(request.AccountEditAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}

	sess := global.Db.NewSession()
	sess.Begin()
	defer sess.Close()

	var columns = []string{"updated_at"}
	user := new(model.AdminUser)
	user.UpdatedAt = time.Now()
	if param.Password != "" {
		newPassword := util.AesDecrypt(param.Password)
		user.Passwd = util.MD5(newPassword)
		columns = append(columns, "passwd")
	}
	if param.NickName != "" {
		user.Nickname = param.NickName
		columns = append(columns, "nickname")
	}
	if param.IsDel != "" {
		if param.AccountId == 100004 {
			global.Logger.Err(err).Msg("不能删除总管理员！")
			response.ResFail(c, "不能删除总管理员！")
			return
		}
		user.IsDel, err = strconv.Atoi(param.IsDel)
		if err != nil {
			global.Logger.Err(err).Msg("参数不合法！")
			response.ResFail(c, "参数不合法！")
			return
		}
		columns = append(columns, "is_del")
	}
	if param.IsReset != "" {
		user.IsReset = 1
		user.Authkey = ""
		columns = append(columns, "is_reset", "authkey")
	}
	if param.Status != "" {
		user.Status, err = strconv.Atoi(param.Status)
		if err != nil {
			global.Logger.Err(err).Msg("参数不合法！")
			response.ResFail(c, "参数不合法！")
			return
		}
		columns = append(columns, "status")
	}
	err = service.UpdateModel(user, sess, param.AccountId, columns...)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错!")
		sess.Rollback()
		response.ResFail(c, "失败！")
		return
	}

	if param.RoleId > 0 {
		userRole := new(model.AdminUserRole)
		userRole.RoleId = param.RoleId
		userRole.UpdatedAt = time.Now()
		userRole.Author = adminUser.Nickname
		rows, err := sess.Cols("role_id", "updated_at", "author").Where("uid = ?", param.AccountId).Update(userRole)
		if err != nil || rows < 1 {
			global.Logger.Err(err).Msg("编辑失败")
			sess.Rollback()
			response.ResFail(c, "失败！")
			return
		}
	}

	sess.Commit()
	response.ResOk(c, "成功")
}

func AddAdminRole(c *gin.Context) {
	param := new(request.RoleAddAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}

	sess := global.Db.NewSession()
	sess.Begin()
	defer sess.Close()

	role := new(model.AdminRole)
	role.Name = param.Name
	role.IsUsed = 1
	role.Remark = param.Remark
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	rows, err := sess.Insert(role)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("添加角色失败")
		sess.Rollback()
		response.ResFail(c, "添加角色失败！")
		return
	}

	roleRes := new(model.AdminRoleRes)
	roleRes.RoleId = role.Id
	roleRes.ResIds = param.ResIds
	roleRes.CreatedAt = time.Now()
	roleRes.UpdatedAt = time.Now()
	roleRes.Author = adminUser.Nickname
	rows, err = sess.Insert(roleRes)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("添加角色权限失败")
		sess.Rollback()
		response.ResFail(c, "添加角色权限失败！")
		return
	}
	sess.Commit()
	service.UpdateCachePermission()
	response.ResOk(c, "成功")
}

func EditAdminRole(c *gin.Context) {
	param := new(request.RoleEditAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	adminUser, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}

	sess := global.Db.NewSession()
	sess.Begin()
	defer sess.Close()

	var columns = []string{"updated_at"}
	role := new(model.AdminRole)
	role.UpdatedAt = time.Now()
	if param.Name != "" {
		role.Name = param.Name
		columns = append(columns, "name")
	}
	if param.Remark != "" {
		role.Remark = param.Remark
		columns = append(columns, "remark")
	}
	if param.IsDel != "" {
		if param.Id == 15 {
			global.Logger.Err(err).Msg("不能删除总管理员！")
			response.ResFail(c, "不能删除总管理员！")
			return
		}
		role.IsDel, err = strconv.Atoi(param.IsDel)
		if err != nil {
			global.Logger.Err(err).Msg("参数不合法！")
			response.ResFail(c, "参数不合法！")
			return
		}
		columns = append(columns, "is_del")
	}

	err = service.UpdateModel(role, sess, param.Id, columns...)
	if err != nil {
		global.Logger.Err(err).Msg("数据库链接出错!")
		sess.Rollback()
		response.ResFail(c, "失败！")
		return
	}

	if param.ResIds != "" {
		roleRes := new(model.AdminRoleRes)
		roleRes.ResIds = param.ResIds
		roleRes.UpdatedAt = time.Now()
		roleRes.Author = adminUser.Nickname
		rows, err := sess.Cols("res_ids", "updated_at", "author").Where("role_id = ?", param.Id).Update(roleRes)
		if err != nil || rows < 1 {
			global.Logger.Err(err).Msg("编辑失败")
			sess.Rollback()
			response.ResFail(c, "失败！")
			return
		}
	}

	sess.Commit()
	service.UpdateCachePermission()
	response.ResOk(c, "成功")
}

func GetFullMenuTree(c *gin.Context) {
	response.RespOk(c, lang.Translate("cn", "success"), service.GetCacheMenuTree())
}

func GetFullTree(c *gin.Context) {
	response.RespOk(c, lang.Translate("cn", "success"), service.GetCacheFullTree())
}

func GetRoleTree(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	userRole := new(model.AdminUserRole)
	has, err := global.Db.Where("uid = ?", user.Id).Get(userRole)
	if err != nil {
		global.Logger.Err(err).Msg("出错！")
		response.ResFail(c, "网络错误！")
		return
	}
	if !has {
		response.ResFail(c, "网络错误！")
		return
	}
	response.RespOk(c, lang.Translate("cn", "success"), service.GetCacheRoleTree(userRole.RoleId))
}

func GetAdminUserList(c *gin.Context) {
	param := new(request.AccountListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	if param.Size >= constant.MaxPageSize {
		param.Size = constant.MaxPageSize
	}
	offset := 0
	if (param.Page - 1) > 0 {
		offset = (param.Page - 1) * param.Size
	}
	session := service.AccountAdminList(param, user)
	count, err := service.AccountAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	var list []map[string]interface{}
	err = session.
		Cols("admin_user.id as id,admin_user.uname as account,admin_user.nickname as nick_name,admin_role.name as role_name,admin_role.id as role_id,admin_user.created_at as created_at,admin_user.status as Status").
		Join("LEFT", model.AdminUserRole{}, "admin_user_role.uid = admin_user.id").
		Join("LEFT", model.AdminRole{}, "admin_user_role.role_id = admin_role.id").
		Limit(param.Size, offset).
		OrderBy("id desc").
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	var dataList response.PageResult
	dataList.Total = count
	dataList.Page = param.Page
	dataList.Size = param.Size
	dataList.List = list
	response.RespOk(c, lang.Translate("cn", "success"), dataList)
}

func GetAdminRoleList(c *gin.Context) {
	param := new(request.RoleListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	if param.Size >= constant.MaxPageSize {
		param.Size = constant.MaxPageSize
	}
	offset := 0
	if (param.Page - 1) > 0 {
		offset = (param.Page - 1) * param.Size
	}
	session := service.RoleAdminList(param, user)
	count, err := service.RoleAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	var list []map[string]interface{}
	err = session.
		Cols("admin_role.*,admin_role_res.res_ids").
		Join("LEFT", model.AdminRoleRes{}, "admin_role_res.role_id = admin_role.id").
		Limit(param.Size, offset).
		OrderBy("id desc").
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	for _, item := range list {
		var tmpIds []int
		resIds := strings.Split(item["res_ids"].(string), ",")
		for _, resId := range resIds {
			id, _ := strconv.Atoi(resId)
			tmpIds = append(tmpIds, id)
		}
		item["resource_ids"] = tmpIds
	}
	var dataList response.PageResult
	dataList.Total = count
	dataList.Page = param.Page
	dataList.Size = param.Size
	dataList.List = list
	response.RespOk(c, lang.Translate("cn", "success"), dataList)
}

func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	var resultMap = make(map[string]interface{})
	resultMap["id"] = user.Id
	resultMap["is_reset"] = user.IsReset
	resultMap["is_first"] = user.IsFirst
	resultMap["nick_name"] = user.Nickname
	resultMap["status"] = user.Status
	response.RespOk(c, "成功", resultMap)
}

func ResetCache(c *gin.Context) {
	service.UpdateCachePermission()
	response.ResOk(c, "成功")
}

func Upload(ctx *gin.Context) {
	param := new(request.UploadFile)
	if err := ctx.ShouldBindWith(param, binding.FormMultipart); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.RespFail(ctx, lang.Translate("cn", "fail"), nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("param: %+v", *param)

	if param.FileType != constant.ImgFileType &&
		param.FileType != constant.OtherFileType {
		response.RespFail(ctx, lang.Translate("cn", "file ext error"), nil)
		return
	}
	fileMap := make(map[string]bool)
	if param.FileType == constant.ImgFileType {
		fileMap[".png"] = true
		fileMap[".jpg"] = true
		fileMap[".jpeg"] = true
	} else if param.FileType == constant.OtherFileType {
		fileMap[".pdf"] = true
		fileMap[".mp4"] = true
		fileMap[".avi"] = true
		fileMap[".dat"] = true
		fileMap[".mkv"] = true
		fileMap[".flv"] = true
		fileMap[".vob"] = true
		fileMap[".mp3"] = true
		fileMap[".wav"] = true
		fileMap[".wma"] = true
		fileMap[".mp2"] = true
		fileMap[".ra"] = true
		fileMap[".ape"] = true
		fileMap[".aac"] = true
		fileMap[".cda"] = true
		fileMap[".mov"] = true
		fileMap[".gif"] = true
		fileMap[".ppt"] = true
		fileMap[".pptx"] = true
		fileMap[".zip"] = true
	}
	global.MyLogger(ctx).Info().Msgf("fileMap: %+v", fileMap)
	fmt.Println("upload请求：", param.FileType)
	resUrl, err := service.Upload(param, fileMap)
	if err != nil {
		global.Logger.Err(err).Msg("upload出错！")
		response.RespFail(ctx, lang.Translate("cn", "fail"), nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("resUrl: %+v", resUrl)
	fmt.Println("upload请求：", param.FileType)
	data := new(response.DataResponse)
	data.Data.Url = resUrl
	response.RespOk(ctx, lang.Translate("cn", "success"), data.Data)
}

// JWTAuth 验证token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization-Token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := service.ParseTokenByUser(token, service.AdminUserType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": "授权已过期",
			})
			c.Abort()
			return
		}

		err = service.AddLog(c, claims.UserId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    100,
				"message": "网络错误",
			})
			c.Abort()
			return
		}
		//isBool := RedisInsert(strconv.FormatInt(claims.UserId, 10))
		//if !isBool {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 301,
		//		"msg":  "授权已过期",
		//	})
		//	c.Abort()
		//	return
		//}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		//uu := c.MustGet("claims").(*service.CustomClaims)
		//fmt.Println("claims...", uu.UserId)
	}
}

func commonPageList(page, size int, total int64, cols string, session *xorm.Session) (response.PageResult, error) {
	if size >= constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	offset := 0
	if (page - 1) > 0 {
		offset = (page - 1) * size
	}
	var list []map[string]interface{}
	err := session.
		Cols(cols).
		Limit(size, offset).
		OrderBy("id desc").
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		return response.PageResult{}, err
	}
	if len(list) == 0 {
		list = []map[string]interface{}{}
	}
	var dataList response.PageResult
	dataList.Total = total
	dataList.Page = page
	dataList.Size = size
	dataList.List = list
	return dataList, nil
}

func commonPageListV2(page, size int, total int64, session *xorm.Session) (response.PageResult, error) {
	if size >= constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	offset := 0
	if (page - 1) > 0 {
		offset = (page - 1) * size
	}
	var list []map[string]interface{}
	err := session.
		Limit(size, offset).
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		return response.PageResult{}, err
	}
	if len(list) == 0 {
		list = []map[string]interface{}{}
	}
	var dataList response.PageResult
	dataList.Total = total
	dataList.Page = page
	dataList.Size = size
	dataList.List = list
	return dataList, nil
}
