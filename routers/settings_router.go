package routers

import (
	"go-blog/api"
)

func (router RouterGroup) SettingsRouter() {
	// 获取配置信息路由函数的对象
	settingsApi := api.ApiGroupApp.SettingsApi
	// 设置配置信息相关路由函数
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
