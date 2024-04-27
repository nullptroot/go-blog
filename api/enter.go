package api

import "go-blog/api/settings_api"

// 相当于封装了一组api，目前只有一个，后续可能有多个api都写入里面
type ApiGroup struct {
	// SettingsApi实现了SettingsInfoView的方法，供其gin注册路由，
	SettingsApi settings_api.SettingsApi
}

// 这个用来取得SettingsApi 来给gin设置路由handler的
var ApiGroupApp = new(ApiGroup)
