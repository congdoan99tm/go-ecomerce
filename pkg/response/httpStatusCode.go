package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParmaInvalid = 20003 // Success
	ErrorInvalidToken     = 30001 // Success
	ErrCodeUserHasExists  = 50001 // user has already registered
)

var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParmaInvalid: "Email is invalid",
	ErrorInvalidToken:     "Token is invalid",
	ErrCodeUserHasExists:  "user has already registered",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    ErrorCodeSuccess,
		Message: msg[ErrorCodeSuccess],
		Data:    data,
	})
}
