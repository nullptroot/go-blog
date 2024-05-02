package menuapi

import (
	"fmt"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 删除菜单

// MenuRemoveView 删除菜单
// @Tags 菜单管理
// @Summary 删除菜单
// @Description 删除菜单
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "图片id列表"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 查出要删除的menu记录
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 事务  接受一个闭包，返回error就回滚
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除MenuModel和BannersModel相关联的记录，也就是MenuBannerModel中的记录
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)

}
