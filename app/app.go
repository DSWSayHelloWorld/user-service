package app

import (
	"github.com/sunil206b/go-microservices/controller"
	"log"
	"net/http"
)

// StartApp function will start the application
func StartApp() {
	http.HandleFunc("/users", controller.GetUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
