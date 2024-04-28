package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"` // 显示的标题
	Href   string `json:"href"`                 // 跳转链接
	Images string `json:"images"`               // 图片
	IsShow bool   `json:"is_show"`              // 是否展示
}

// MenuImageModel 自定义菜单和背景图的连接表，方便排序
type MenuImageModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
