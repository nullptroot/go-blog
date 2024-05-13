package messageapi

import (
	"go-blog/models"
	"go-blog/models/res"
	"go-blog/service/common"

	"github.com/gin-gonic/gin"
)

// 获取全部消息，管理员才能操作
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	// 获取query参数
	err := c.ShouldBindQuery(&cr)
	// fmt.Println(cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 获取根据要求获取图片list ，表中的所有记录都获取了
	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
