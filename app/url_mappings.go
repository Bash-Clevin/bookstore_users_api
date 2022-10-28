package app

import (
	"github.com/Bash-Clevin/bookstore_users_api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
