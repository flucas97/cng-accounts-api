package app

import (
	"github.com/flucas97/CNG-checknogreen/account/controllers/accounts_controller"
	"github.com/flucas97/CNG-checknogreen/account/controllers/ping"
)

// Routes map all avaliable routes
func Routes() {
	router.GET("/ping", ping.Ping)

	// accounts
	router.POST("/new-account", accounts_controller.Create)
	router.POST("/validate", accounts_controller.Validate)
	router.POST("/account-details", accounts_controller.ShowDetails)
}
