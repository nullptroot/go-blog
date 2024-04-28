package flag

import (
	"go-blog/core"
	"go-blog/global"
	"go-blog/routers"
	"testing"
)

// 加参数不太行啊，测试不了
func TestFlag(t *testing.T) {
	core.InitConf()
	core.InitLogger()
	core.InitGorm()
	option := Parse()
	global.Log.Info(option.DB)
	if IsWebStop(option) {
		SwitchOption(option)
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("Listening on : %s\n", addr)
	router.Run(addr)
}
