package routes

import (
	"net/http"

	"canteen-pricing/handlers"
)

func RegisterRoutes() {

	http.HandleFunc("/calculate", handlers.CalculatePrice)

}