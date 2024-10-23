package routers

import (
	"fmt"
	. "github.com/dinos/go-ecommerce-be-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}

}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}

}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")

}
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(AA(), BB(), CC)
	//userController, _ := wire.InitUserRouterHandle()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", NewPongController().Pong)
		//v1.GET("/users", NewUserController().GetUserById)
	}
	return r
}
