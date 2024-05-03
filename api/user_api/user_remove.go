package userapi

import (
	"fmt"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRemoveView 批量删除用户
// @Tags 用户管理
// @Summary 批量删除用户
// @Description 批量删除用户
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "用户id列表"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)

}
