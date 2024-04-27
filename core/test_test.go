package core

import (
	"fmt"
	"go-blog/global"
	"testing"
)

// 测试conf文件
func TestConf(t *testing.T) {
	// 读取配置文件
	InitConf()
	fmt.Println(global.Config)
}

// 测试gorm文件 需要先初始化配置文件的
func TestGorm(t *testing.T) {
	InitConf()
	global.Log = InitLogger()
	global.DB = InitGorm()
	fmt.Println(global.DB)
}

// 测试日志模块
func TestLog(t *testing.T) {
	InitConf()
	global.Log = InitLogger()
	global.Log.Info("wqwqwq")
}
