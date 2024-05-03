package userapi

import (
	"go-blog/global"
	"go-blog/models/res"
	"go-blog/service"
	"go-blog/utils/jwts"

	"github.com/gin-gonic/gin"
)

//	注销就是直接把当前的token放入redis，不再能使用，redis过期 token也就过期了
//
// LogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	res.OkWithMessage("注销成功", c)

}
