package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/jeffotoni/gocep/pkg/cep"

	"github.com/rs/cors"
)

var (
	Port = "0.0.0.0:8080"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/cep/", HandlerCep)
	mux.HandleFunc("/cep", NotFound)
	mux.HandleFunc("/", NotFound)
	muxcors := cors.Default().Handler(mux)
	server := &http.Server{
		Addr:    Port,
		Handler: muxcors,
	}

	log.Println("Run My Server ", Port)
	log.Fatal(server.ListenAndServe())
}

func HandlerCep(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	cepstr := strings.Split(r.URL.Path[1:], "/")[1]
	if len(cepstr) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, wecep, err := cep.Search(cepstr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !cep.ValidCep(wecep) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var b []byte
	if len(result) > 0 {
		b = []byte(string(result))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}
