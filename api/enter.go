package api

import (
	advertapi "go-blog/api/advert_api"
	articleapi "go-blog/api/article_api"
	diggapi "go-blog/api/digg_api"
	imagesapi "go-blog/api/images_api"
	menuapi "go-blog/api/menu_api"
	messageapi "go-blog/api/message_api"
	"go-blog/api/settings_api"
	tagapi "go-blog/api/tag_api"
	userapi "go-blog/api/user_api"
)

// 相当于封装了一组api，目前只有一个，后续可能有多个api都写入里面
type ApiGroup struct {
	SettingsApi settings_api.SettingsApi // SettingsApi实现了SettingsInfoView的方法，供其gin注册路由，
	ImagesApi   imagesapi.ImagesApi      // 图片相关的接口
	AdvertApi   advertapi.AdvertApi      // 广告相关的接口
	MenuApi     menuapi.MenuApi          // 菜单管理相关的接口
	UserApi     userapi.UserApi          // 登录相关的接口
	TagApi      tagapi.TagApi            // 标签相关接口
	MessageApi  messageapi.MessageApi    // 消息相关的接口
	ArticleApi  articleapi.ArticleApi    // 文章相关的接口
	DiggApi     diggapi.DiggApi          //点赞相关的接口
}

// 这个用来取得SettingsApi 来给gin设置路由handler的
var ApiGroupApp = new(ApiGroup)
