package models

// MenuBannerModel 自定义菜单和背景图的连接表，方便排序
type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"` // 外表没什么用可以删除
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"` // 外表没什么用可以删除
	Sort        int         `gorm:"size:10" json:"sort"`
}
