package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp function will start the application
func StartApp() {
	mapURL()
	log.Fatal(router.Run(":8080"))
}
