package initialize

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Println("Loading configuration mysql ...", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "Success"))
	InitMysql()
	InitServiceInterface()
	InitRedis()
	InitKafka()
	r := InitRouter()
	return r
}
