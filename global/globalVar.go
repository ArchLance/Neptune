package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"neptune/config"
)

var (
	DB    *gorm.DB
	Redis *redis.Client

	ServerConfig config.ServerConfig
)
