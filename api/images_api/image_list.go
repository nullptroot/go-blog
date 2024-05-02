package imagesapi

import (
	"go-blog/models"
	"go-blog/models/res"
	"go-blog/service/common"

	"github.com/gin-gonic/gin"
)

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	// 获取query参数
	err := c.ShouldBindQuery(&cr)
	// fmt.Println(cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 获取根据要求获取图片list
	list, count, _ := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
