package routers

import (
	"go-blog/api"
	"go-blog/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	// 获取图片信息路由函数相关的对象
	app := api.ApiGroupApp.UserApi
	// 设置seesion
	// 一个router一个
	router.Use(sessions.Sessions("sessionid", store))
	// 配置图片相关的路由函数
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	router.POST("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.GET("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
	router.POST("users", middleware.JwtAdmin(), app.UserCreateView)
	// router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
