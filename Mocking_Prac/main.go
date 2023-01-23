package main

import (
	"mocking/controller"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controller.Ping)
	router.Run(":8111")
}
