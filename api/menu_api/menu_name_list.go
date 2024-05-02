package menuapi

import (
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// 返回menu一些简单的信息
// MenuNameList 菜单名称列表
// @Tags 菜单管理
// @Summary 菜单名称列表
// @Description 菜单名称列表
// @Router /api/menus_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=MenuNameResponse}
func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	res.OkWithData(menuNameList, c)
}
