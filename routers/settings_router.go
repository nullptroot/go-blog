package routers

import (
	"go-blog/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", settingsApi.SettingsInfoView)
}
