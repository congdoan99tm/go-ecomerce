package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code   int         `json:"code"`
	Err    string      `json:"error"`
	Detail interface{} `json:"detail"`
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
