package main

import (
	"log"
	"net/http"

	"github.com/myazahq/go-test-exercise/section3/exercise1/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HelloWorld) // updated routing as at Go1.22

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Starting server at port :8080")
	log.Fatal(server.ListenAndServe())
}
