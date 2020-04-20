package main

import (
	"github.com/jeffotoni/gocep/config"
	"github.com/jeffotoni/gocep/handler"
	"log"
	"net/http"
)

type Result struct {
	Body []byte
}

var chResult = make(chan Result, len(models.Endpoints))

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/", Handler.SearchCep)
	mux.HandleFunc("/api/v1", Handler.NotFound)
	mux.HandleFunc("/", Handler.NotFound)

	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	log.Println("Port:", config.Port)
	log.Fatal(server.ListenAndServe())
}
