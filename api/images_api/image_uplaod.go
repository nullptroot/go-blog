package imagesapi

import (
	"go-blog/global"
	"go-blog/models/res"
	"go-blog/service"
	imageser "go-blog/service/image_ser"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
)

// ImageUploadView 上传单个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	// 上传多个图片，MultipartForm分析多个上传
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	// images字段是规定的请求的参数
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	// 判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		// 不存在就创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 上传图片的响应信息，回发给客户端的
	var resList []imageser.FileUploadResponse
	for _, file := range fileList {
		// 上传文件
		// 上传的主逻辑
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 成功的，并且不是上传到七牛云的，就需要保存到本地上
		if !global.Config.QiNiu.Enable {
			// 本地还得保存一下
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		// 获取响应数据
		resList = append(resList, serviceRes)
	}

	res.OkWithData(resList, c)

}
