package app

import (
	"net/http"

	"github.com/hladera100/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodPost)
	router.HandleFunc("/items", controllers.ItemsControllers.Create).Methods(http.MethodPost)
}
