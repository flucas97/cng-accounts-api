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
	router.POST("/login", accounts_controller.Login)
	router.POST("/details", accounts_controller.ShowDetails)
}
