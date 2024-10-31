package initialize

import (
	"github.com/dinos/go-ecommerce-be-api/global"
	"github.com/dinos/go-ecommerce-be-api/internal/database"
	"github.com/dinos/go-ecommerce-be-api/internal/service"
	"github.com/dinos/go-ecommerce-be-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
