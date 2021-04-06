package app

import (
	"github.com/JenniO/bookstore_users-api/controllers/ping"
	"github.com/JenniO/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("ping", ping.Ping)

	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)   // full update
	router.PATCH("/users/:user_id", users.Update) // partial update
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
