package initialize

import (
	"github.com/dinos/go-ecommerce-be-api/global"
	"github.com/dinos/go-ecommerce-be-api/pkg/loggers"
)

func InitLogger() {
	global.Logger = loggers.NewLogger(global.Config.Logger)
}
