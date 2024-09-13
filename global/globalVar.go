package global

import (
	"github.com/jordan-wright/email"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"neptune/config"
)

var (
	DB           *gorm.DB
	Redis        *redis.Client
	EmailPool    *email.Pool
	ServerConfig config.ServerConfig
)
