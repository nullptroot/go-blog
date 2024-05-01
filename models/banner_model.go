package models

import (
	"go-blog/global"
	"go-blog/models/ctype"
	"os"

	"gorm.io/gorm"
)

// BannerModel banner表
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型， 本地还是七牛
}

// 删除回调，删除某记录的时候会调用这个东西
// 目前删除图片记录会删除本地文件，应该有引用计数部分
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
