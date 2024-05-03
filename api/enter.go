package api

import (
	advertapi "go-blog/api/advert_api"
	imagesapi "go-blog/api/images_api"
	menuapi "go-blog/api/menu_api"
	"go-blog/api/settings_api"
	tagapi "go-blog/api/tag_api"
	userapi "go-blog/api/user_api"
)

// 相当于封装了一组api，目前只有一个，后续可能有多个api都写入里面
type ApiGroup struct {
	// SettingsApi实现了SettingsInfoView的方法，供其gin注册路由，
	SettingsApi settings_api.SettingsApi
	// 图片相关的接口
	ImagesApi imagesapi.ImagesApi
	// 广告相关的接口
	AdvertApi advertapi.AdvertApi
	// 菜单管理相关的接口
	MenuApi menuapi.MenuApi
	// 登录相关的接口
	UserApi userapi.UserApi
	// 标签相关接口
	TagApi tagapi.TagApi
}

// 这个用来取得SettingsApi 来给gin设置路由handler的
var ApiGroupApp = new(ApiGroup)
