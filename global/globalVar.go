package global

import (
	"gorm.io/gorm"
	"neptune/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)
