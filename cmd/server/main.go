package main

import (
	"github.com/dinos/go-ecommerce-be-api/internal/initialize"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/dinos/go-ecommerce-be-api/cmd/swag/docs"
)

import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files

// @title           API Documentation Ecomerrce Backend SHOPDEVGO
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/congdoan99tm/go-ecomerce

// @contact.name   TEAM TIPSGO
// @contact.url	   github.com/congdoan99tm/go-ecomerce
// @contact.email  congdoan99tm@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http

func main() {

	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}
