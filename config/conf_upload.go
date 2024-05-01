package config

// 图片上传的位置和大小限制
type Upload struct {
	Size int    `yaml:"size" json:"size"` //图片上传的大小
	Path string `yaml:"path" json:"path"` //图片上传的目录
}
