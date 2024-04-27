package main

import (
	"fmt"
	"go-blog/core"
	"go-blog/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
