package app

import "github.com/flucas97/CNG-checknogreen/account/controllers/ping"

// Routes map all avaliable routes
func Routes() {
	router.GET("/ping", ping.Ping)
}
