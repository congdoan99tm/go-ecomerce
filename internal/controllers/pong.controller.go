package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}
func (p PongController) Pong(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
	})
}
