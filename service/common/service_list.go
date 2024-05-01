package common

import (
	"go-blog/global"
	"go-blog/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}
	// 获取一共有多少条记录，并且把结果存入了list
	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	// -1就是取消limit
	if option.Limit == 0 {
		option.Limit = -1
	}
	// 按照规则取得相应的记录
	err = DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
