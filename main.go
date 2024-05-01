package main

import (
	"go-blog/core"
	"go-blog/flag"
	"go-blog/global"
	"go-blog/routers"
)

func main() {
	// 初始化基础信息
	core.InitConf()
	// 初始化日志
	core.InitLogger()
	// 初始化gorm
	core.InitGorm()
	// 获取命令行参数
	option := flag.Parse()
	// global.Log.Info(option.DB)
	// 若是迁移数据库，那么就进行迁移
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 初始化router
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("Listening on : %s\n", addr)
	// 开始运行
	router.Run(addr)
}
