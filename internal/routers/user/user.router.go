package user

import (
	"github.com/dinos/go-ecommerce-be-api/internal/controllers/account"
	"github.com/dinos/go-ecommerce-be-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandle()
	// public routers
	userRouterPublic := Router.Group("/user", userController.Register)
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/login", account.Login.Login)

	}
	// private routers
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authenticate())
	//userRouterPrivate.Use(Permission())

	{
		userRouterPrivate.GET("/get_info")

	}
}
