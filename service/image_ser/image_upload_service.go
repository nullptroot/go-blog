package imageser

import (
	"fmt"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/ctype"
	"go-blog/plugins/qiniu"
	"go-blog/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var WhiteImageList = []string{
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
}

// 响应信息的数据结构
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadService 文件上传的方法 上传的主逻辑
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	// 获取文件的信息
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath

	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	// 就是判断文件的后缀是否在白名单中
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		return
	}

	// 读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, _ := io.ReadAll(fileObj)
	// 计算文件的Md5
	imageHash := utils.Md5(byteData)
	// 去数据库中查这个图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 找到了
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	// 这里先假设上传成功了
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	// 看看七牛云是否开启了，开启后就上传七牛云
	if global.Config.QiNiu.Enable {
		// 上传到七牛云的逻辑
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}
	// 图片入库
	global.DB.Debug().Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
