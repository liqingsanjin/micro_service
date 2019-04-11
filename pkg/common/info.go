package common

import (
	"userService/pkg/redis"

	"github.com/jinzhu/gorm"
)

var (
	RedisClientService redis.ClientService
	DB                 *gorm.DB
)
