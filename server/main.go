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

	log.Println("Server running at http://0.0.0.0:5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
