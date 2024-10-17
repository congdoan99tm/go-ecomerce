package global

import (
	"github.com/dinos/go-ecommerce-be-api/pkg/loggers"
	"github.com/dinos/go-ecommerce-be-api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *loggers.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
