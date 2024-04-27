package settings_api

import (
	"go-blog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	// c.JSON(200, gin.H{"msg": "hello world"})
	// res.Ok("sasa", "ok", c)
	res.OkWithData(map[string]any{
		"user":   "admin",
		"passwd": "admin",
	}, c)
}
