package menuapi

import (
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
)

// MenuDetailView 菜单细节
// @Tags 菜单管理
// @Summary 菜单细节
// @Description 菜单细节
// @Param id path int true "id"
// @Router /api/menus/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=MenuResponse}
func (MenuApi) MenuDetailView(c *gin.Context) {
	// 先查菜单 获取menuMode id等于id的记录
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Debug().Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 查连接表MenuBannerModel  查找menu_id等于id的记录
	var menuBanners []models.MenuBannerModel
	// 预加载，把BannerModel的数据挂载到了MenuBannerModel上，按照bannerID连接的
	global.DB.Debug().Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var banners = make([]Banner, 0)
	// 组合banners的数据
	for _, banner := range menuBanners {
		// 下面的去重没有意义的
		// if menuModel.ID != banner.MenuID {
		// 	continue
		// }
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	// 拼接到一起
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
}
