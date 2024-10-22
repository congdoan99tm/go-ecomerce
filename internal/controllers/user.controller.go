package controllers

import (
	"github.com/dinos/go-ecommerce-be-api/internal/service"
	"github.com/dinos/go-ecommerce-be-api/pkg/response"
	"github.com/gin-gonic/gin"
)

//type UserController struct {
//	userService *service.UserService
//}
//
//func NewUserController() *UserController {
//	return &UserController{userService: service.NewUserService()}
//}
//func (uc UserController) GetUserById(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": uc.userService.GetInfoUser(),
//		"users":   []string{"cr7", "m10", "doan"},
//	})
//}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result)
}
