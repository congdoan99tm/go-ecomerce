package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/register")
	}

	userRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())

	{
		userRouterPrivate.GET("/active_user")
	}
}
