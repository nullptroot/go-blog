package tagapi

import (
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
)

// 前端传递false会直接忽略 ,下面的IsShow 不能binding了
type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

// TagCreateView 添加标签
// @Tags 标签管理
// @Summary 添加标签
// @Description 添加标签
// @Param data body TagRequest    true  "标签标题"
// @Param token header string  true  "token"
// @Router /api/tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复的判断
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}
	// 直接写入数据库即可
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}

	res.OkWithMessage("添加标签成功", c)
}
