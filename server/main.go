package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()
	defer DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/search", SearchCarsHandler)

	log.Println("Server running at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
