package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gocep/config"
	handler "github.com/jeffotoni/gocep/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/", handler.SearchCep)
	mux.HandleFunc("/api/v1", handler.NotFound)
	mux.HandleFunc("/", handler.NotFound)
	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	log.Println("Port:", config.Port)
	log.Fatal(server.ListenAndServe())
}
