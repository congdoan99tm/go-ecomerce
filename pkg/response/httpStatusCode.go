package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParmaInvalid  = 20003 // Success
	ErrInvalidToken      = 30001 // Success
	ErrInvalidOTP        = 3002
	ErrCodeUserHasExists = 50001 // user has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParmaInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrCodeUserHasExists: "user has already registered",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    ErrCodeSuccess,
		Message: msg[ErrCodeSuccess],
		Data:    data,
	})
}
