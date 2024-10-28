package global

import (
	"database/sql"
	"github.com/dinos/go-ecommerce-be-api/pkg/loggers"
	"github.com/dinos/go-ecommerce-be-api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        settings.Config
	Logger        *loggers.LoggerZap
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
)
