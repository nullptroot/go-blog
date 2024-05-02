package routers

import (
	"go-blog/global"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	// 组router  这是继承嗷，
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// global.ViewMaps = map[string]any{
	// 	"site":  &global.Config.SiteInfo,
	// 	"email": &global.Config.Email,
	// 	"qq":    &global.Config.QQ,
	// 	"qiniu": &global.Config.QiNiu,
	// 	"jwt":   &global.Config.Jwt,
	// }
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//swag api文档路由
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 路由分组  也就是url都已api开头的/api/*
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	// 设置配置信息相关路由
	routerGroupApp.SettingsRouter()
	// 设置图片相关的路由
	routerGroupApp.ImagesRouter()
	// 设置广告相关的路由
	routerGroupApp.AdvertRouter()
	return router
}
