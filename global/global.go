package global

import (
	"go-blog/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	// ViewMaps map[string]any // view的maps
)
