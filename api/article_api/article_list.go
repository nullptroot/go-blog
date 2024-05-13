package articleapi

import (
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"
	esser "go-blog/service/es_ser"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag  string `json:"tag" form:"tag"`
	Sort string `json:"sort" form:"sort"`
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param data query ArticleSearchRequest    false  "查询参数"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := esser.CommList(esser.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}
	// json-filter 空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
