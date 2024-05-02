package routers

import "go-blog/api"

func (router RouterGroup) MenuRouter() {
	// 获取图片信息路由函数相关的对象
	app := api.ApiGroupApp.MenuApi
	// 配置图片相关的路由函数
	router.POST("menus", app.MenuCreateView)
	router.GET("menus", app.MenuListView)
	router.GET("menus_names", app.MenuNameList)
	router.PUT("menus/:id", app.MenuUpdateView)
	router.DELETE("menus", app.MenuRemoveView)
	router.GET("menus/:id", app.MenuDetailView)
}
