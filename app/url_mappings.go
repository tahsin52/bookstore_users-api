package app

import (
	"github.com/tahsin52/bookstore_users-api/controllers/ping"
	"github.com/tahsin52/bookstore_users-api/controllers/users"
)

func mapUrls()  {
	rooter.GET("/ping", ping.Ping)

	rooter.GET("/users/:user_id", users.GetUser)
	//rooter.GET("/users/search", controllers.SearchUser)
	rooter.POST("/users", users.CreateUser)
	
}