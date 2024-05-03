package userapi

import (
	"go-blog/models"
	"go-blog/models/ctype"
	"go-blog/models/res"
	"go-blog/service/common"
	"go-blog/utils/desens"
	"go-blog/utils/jwts"

	"github.com/gin-gonic/gin"
)

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param token header string  true  "token"
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	// 添加了中间件
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		// 脱敏
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)

		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
