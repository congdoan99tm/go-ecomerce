package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}
func (uc UserController) GetUserById(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "doan"},
	})
}
