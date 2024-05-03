package global

import (
	"go-blog/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	// ViewMaps map[string]any // view的maps
	// mysql的日志信息
	MysqlLog logger.Interface
	Redis    *redis.Client
)
