package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/admin"
)

func AdminRoute(group *gin.RouterGroup) {
	group.POST("login", admin.LoginAdmin)
	group.POST("upload", admin.Upload)

	group.Use(admin.JWTAuth())
	{
		//公共
		group.GET("full_menu_tree", admin.GetFullMenuTree)
		group.GET("role_menu_tree", admin.GetRoleTree)
		group.GET("user_info", admin.UserInfo)
		group.GET("reset_cache", admin.ResetCache) //刷新缓存

		//首页

		//系统管理
		sysGroup := group.Group("sys")
		sysGroup.GET("full_tree", admin.GetFullTree)
		sysGroup.POST("add_resource", admin.AddResource)
		sysGroup.POST("del_resource", admin.DelResource)
		sysGroup.POST("edit_resource", admin.EditResource)
		sysGroup.POST("edit_passwd", admin.EditPasswd)
		sysGroup.POST("generate_auth2Key", admin.GenerateAuth2Key) //生成两步验证器私钥
		sysGroup.POST("set_auth2Key", admin.SetAuth2Key)           //设置两步验证器

		//管理员模块
		adminGroup := group.Group("admin")
		adminGroup.GET("user_list", admin.GetAdminUserList)
		adminGroup.POST("add_user", admin.AddAdminUser)
		adminGroup.POST("edit_user", admin.EditAdminUser)
		adminGroup.GET("role_list", admin.GetAdminRoleList)
		adminGroup.POST("add_role", admin.AddAdminRole)
		adminGroup.POST("edit_role", admin.EditAdminRole)
	}
}
