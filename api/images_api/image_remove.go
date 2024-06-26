package imagesapi

import (
	"fmt"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 删除图片
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "图片id列表"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)

}
