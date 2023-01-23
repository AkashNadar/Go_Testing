package controller

import (
	"mocking/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	res, err := services.PingService.HandlePing()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
