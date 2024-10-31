package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParmaInvalid  = 20003 // Success
	ErrInvalidToken      = 30001 // Success
	ErrInvalidOTP        = 30002
	ErrSendEmailOtp      = 30003
	ErrCodeUserHasExists = 50001 // user has already registered
	ErrCodeOtpNotExists  = 60009
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParmaInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrSendEmailOtp:      "Failed to send email OTP",
	ErrCodeUserHasExists: "user has already registered",
	ErrCodeOtpNotExists:  "otp exists",
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

func ErrResponse(c *gin.Context, errCode int, errMsg string) {
	c.JSON(http.StatusOK, Response{
		Code:    errCode,
		Message: errMsg,
	})
}
