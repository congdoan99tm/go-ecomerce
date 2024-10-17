package initialize

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql ...", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "Success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
