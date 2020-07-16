package app

import (
	"database/sql"

	"github.com/sunil206b/go-microservices/controller"
)

func mapURL() {
	c := controller.NewUserHandler(&sql.DB{})
	router.GET("/users/:id", c.GetUser)
}
