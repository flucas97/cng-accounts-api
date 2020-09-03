package app

import (
	"github.com/flucas97/CNG-checknogreen/account/controllers/accounts_controller"
	"github.com/flucas97/CNG-checknogreen/account/controllers/ping"
)

// Routes map all avaliable routes
func Routes() {
	router.GET("/api/ping", ping.Ping)

	// accounts
	router.POST("/api/signup", accounts_controller.New)
	router.POST("/api/validate", accounts_controller.Validate)
	router.POST("/api/account-details", accounts_controller.ShowDetails)
}
