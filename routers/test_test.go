package routers

import (
	"go-blog/core"
	"go-blog/global"
	"testing"
)

func TestRouter(t *testing.T) {
	core.InitConf()
	core.InitLogger()
	router := InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("Listening on : %s\n", addr)
	router.Run(addr)
}
