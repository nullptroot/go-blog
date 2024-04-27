package core

import (
	"fmt"
	"go-blog/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	// 若没有读到Host 就不用向下走了
	if global.Config.Mysql.Host == "" {
		global.Log.Errorln("未配置mysql，取消gorm连接")
		return nil
	}
	// 从mysql配置中获取dsn信息
	dsn := global.Config.Mysql.Dsn()
	// 初始化logger
	var mysqlLogger logger.Interface
	// 根据level设置logger level信息
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	// 连接到mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalln(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	// 设置一些信息
	sqlDB, _ := db.DB()
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 最大连接时间
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	global.DB = db
	return db
}
