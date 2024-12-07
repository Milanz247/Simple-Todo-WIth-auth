package main

import (
	"Golang/config"
	"Golang/models"
	"Golang/routes"
	"log"
	"net/http"
)

func main() {

	config.InitDB() // Initialize the database
	
	r := http.NewServeMux() // Initialize the router
	routes.RegisterRoutes(r) // Load all routes


	config.InitDB().AutoMigrate(&models.User{},&models.Todo{})
	
	
	// Start the server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
