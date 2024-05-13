package articleapi

import (
	"go-blog/models/res"
	esser "go-blog/service/es_ser"
	redisser "go-blog/service/redis_ser"

	"github.com/gin-gonic/gin"
)

// type ESIDRequest struct {
// 	ID string `json:"id" form:"id" uri:"id"`
// }

// ArticleDetailView 文章细节通过文章id
// @Tags 文章管理
// @Summary 文章细节通过文章id
// @Description 文章细节通过文章id
// @Param id path string true "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	// var cr ESIDRequest
	// err := c.ShouldBindUri(&cr)
	// if err != nil {
	// 	res.FailWithCode(res.ArgumentError, c)
	// 	return
	// }
	ID := c.Param("id")
	// global.Log.Info(ID)
	redisser.Look(ID)
	model, err := esser.CommDetail(ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

// ArticleDetailView 文章细节通过文章题目
// @Tags 文章管理
// @Summary 文章细节通过文章题目
// @Description 文章细节通过文章题目
// @Param data query ArticleDetailRequest    false  "文章标题"
// @Router /api/articles/detail [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// Title := c.Param("title")
	// global.Log.Info(Title)
	model, err := esser.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
