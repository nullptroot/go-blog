package routers

import (
	"go-blog/api"
)

func (router RouterGroup) DiggRouter() {
	app := api.ApiGroupApp.DiggApi
	router.POST("digg/article", app.DiggArticleView)
}
