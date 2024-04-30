package routers

import (
	"go-blog/global"

	"github.com/gin-gonic/gin"
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
	// 路由分组  也就是url都已api开头的/api/*
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	routerGroupApp.SettingsRouter()
	return router
}
