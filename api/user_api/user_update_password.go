package userapi

import (
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"
	"go-blog/utils/jwts"
	"go-blog/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"` // 旧密码
	Pwd    string `json:"pwd"`     // 新密码
}

// UserUpdatePassword 修改登录人的id

// UserUpdatePassword 更新密码
// @Tags 用户管理
// @Summary 更新密码
// @Description 更新密码
// @Param token header string  true  "token"
// @Param data body UpdatePasswordRequest    true  "原密码和新密码"
// @Router /api/user_password [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashPwd, _ := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
}
