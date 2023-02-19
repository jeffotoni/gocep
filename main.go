package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gocep/config"
	handler "github.com/jeffotoni/gocep/handlers"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/", handler.SearchCep)
	mux.HandleFunc("/api/v1", handler.NotFound)
	mux.HandleFunc("/", handler.NotFound)
	muxcors := cors.Default().Handler(mux)
	server := &http.Server{
		Addr:    config.Port,
		Handler: muxcors,
	}
	log.Println(gcolor.YellowCor("Server Run Port"), config.Port)
	log.Println(gcolor.YellowCor("/api/v1/:cep"))
	log.Fatal(server.ListenAndServe())
}
