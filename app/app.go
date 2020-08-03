package app

import (
	"github.com/flucas97/CNG-checknogreen/account/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts the service
func StartApp() {
	Routes()
	logger.Info("Starting server...")
	router.Run(":8081")
}
