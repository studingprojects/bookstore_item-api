package app

import (
	"net/http"

	"github.com/studingprojects/bookstore_item-api/controllers"
)

func mapURLs() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
