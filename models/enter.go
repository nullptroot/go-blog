package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"updated_at"`           //更新时间
}

// 根据ID删除某些记录
type RemoveRequest struct {
	IDList []uint `json:"id_list"` // 要删除的记录id
}

// 查找图片时分页、排序、模糊查询用到的key
type PageInfo struct {
	Page  int    `form:"page"`  // 从第几页开始，后续会生成offset
	Key   string `form:"key"`   // 模糊查询使用的关键字
	Limit int    `form:"limit"` // 一页限制多少记录
	Sort  string `form:"sort"`  // 排序字段
}
