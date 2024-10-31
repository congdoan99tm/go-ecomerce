package account

import (
	"github.com/dinos/go-ecommerce-be-api/global"
	"github.com/dinos/go-ecommerce-be-api/internal/model"
	"github.com/dinos/go-ecommerce-be-api/internal/service"
	"github.com/dinos/go-ecommerce-be-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// management controller Login User
var Login = new(cUserLogin)

type cUserLogin struct {
}

func (c *cUserLogin) Login(ctx *gin.Context) {

	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrResponse(ctx, response.ErrCodeParmaInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, nil)
}
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrResponse(ctx, response.ErrCodeParmaInvalid, err.Error())
		return
	}
	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus)
}
