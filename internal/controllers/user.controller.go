package controllers

import (
	"github.com/dinos/go-ecommerce-be-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}
func (uc UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"cr7", "m10", "doan"},
	})
}
