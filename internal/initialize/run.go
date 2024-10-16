package initialize

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql ...", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
