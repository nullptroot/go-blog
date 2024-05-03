package tagapi

import (
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Param token header string  true  "token"
// @Description 更新标签
// @Param data body TagRequest    true  "标签的一些参数"
// @Param id path int true "id"
// @Router /api/tags/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagUpdateView(c *gin.Context) {

	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 结构体转map的第三方包  直接使用cr 遇见false会直接忽略，那么就更新不了数据库了
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", c)
		return
	}

	res.OkWithMessage("修改标签成功", c)
}
