package routers

import (
	"net/http"

	"github.com/dinos/go-ecommerce-be-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping:name", Pong)
     v1.GET("/users",controllers.NewUserController().GetUserById)
	}
	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
	})
}
