package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	log.Println("Server running at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
