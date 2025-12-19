package main

import (
	"log"
	"net/http"
)

func main() {
	r := SetupRoutes()

	log.Println("Server Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
