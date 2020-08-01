package app

import "github.com/flucas97/CNG-checknogreen/account/controllers/ping"

func Routes() {
	router.GET("/ping", ping.Ping)
}
