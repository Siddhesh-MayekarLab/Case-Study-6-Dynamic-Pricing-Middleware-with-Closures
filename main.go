package main

import (
	"fmt"
	"net/http"

	"canteen-pricing/routes"
)

func main() {

	// Register API routes
	routes.RegisterRoutes()

	// Serve frontend files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("===================================")
	fmt.Println(" Smart Canteen Server Started")
	fmt.Println(" http://localhost:8080")
	fmt.Println("===================================")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}