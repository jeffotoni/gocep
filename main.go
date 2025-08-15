package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gocep/config"
	handler "github.com/jeffotoni/gocep/handlers"

	"github.com/jeffotoni/gcolor"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/cep/", handler.SearchCep)
	mux.HandleFunc("/v1/cep", handler.NotFound)
	mux.HandleFunc("/", handler.NotFound)
	muxcors := cors.Default().Handler(mux)
	server := &http.Server{
		Addr:    config.Port,
		Handler: muxcors,
	}
	log.Println(gcolor.YellowCor("Server Run Port"), config.Port)
	log.Println(gcolor.YellowCor("/v1/cep/:cep"))
	log.Fatal(server.ListenAndServe())
}
