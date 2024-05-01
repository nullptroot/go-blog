package routers

import "go-blog/api"

func (router RouterGroup) ImagesRouter() {
	// 获取图片信息路由函数相关的对象
	app := api.ApiGroupApp.ImagesApi
	// 配置图片相关的路由函数
	router.GET("images", app.ImageListView)
	router.POST("images", app.ImageUploadView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)
	// router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
