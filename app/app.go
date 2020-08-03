package app

import (
	"github.com/flucas97/CNG-checknogreen/account/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts server
func StartApp() {
	Routes()
	logger.Info("Starticng server...")
	router.Run(":8080")
}
