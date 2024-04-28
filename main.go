package main

import (
	"go-blog/core"
	"go-blog/flag"
	"go-blog/global"
	"go-blog/routers"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.InitGorm()
	option := flag.Parse()
	global.Log.Info(option.DB)
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("Listening on : %s\n", addr)
	router.Run(addr)
}
