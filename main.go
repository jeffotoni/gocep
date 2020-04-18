package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jeffotoni/gocep/config"
	"github.com/jeffotoni/gocep/models"
	"github.com/jeffotoni/gocep/service/ristretto"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Result struct {
	Body []byte
}

var chResult = make(chan Result, len(models.Endpoints))

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/", SearchCep)
	mux.HandleFunc("/api/v1", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	log.Println("Port:", config.Port)
	log.Fatal(server.ListenAndServe())
}

func SearchCep(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}
	validEndpoint := strings.Split(r.URL.Path, "/")
	if len(validEndpoint) > 4 {
		w.WriteHeader(http.StatusFound)
		return
	}

	cep := strings.Split(r.URL.Path[2:], "/")[2]
	if err := checkCep(cep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonCep := ristretto.Get(cep)
	if len(jsonCep) > 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonCep))
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, e := range models.Endpoints {
		endpoint := fmt.Sprintf(e.Url, cep)
		source := e.Source
		go func(cancel context.CancelFunc, source, endpoint string, chResult chan<- Result) {
			req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
			if err != nil {
				return
			}

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				return
			}

			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Error ioutil.ReadAll:", err)
				return
			}

			var wecep = &models.WeCep{}
			if len(string(body)) > 0 &&
				response.StatusCode == http.StatusOK {
				switch source {
				case "viacep":
					var viacep = models.ViaCep{}
					err := json.Unmarshal(body, &viacep)
					if err == nil {
						wecep.Cidade = viacep.Localidade
						wecep.Uf = viacep.Uf
						wecep.Logradouro = viacep.Logradouro
						wecep.Bairro = viacep.Bairro
						b, err := json.Marshal(wecep)
						if err == nil {
							chResult <- Result{Body: b}
							cancel()
						}
					}
				case "postmon":
					var postmon = models.PostMon{}
					err := json.Unmarshal(body, &postmon)
					if err == nil {
						wecep.Cidade = postmon.Cidade
						wecep.Uf = postmon.Estado
						wecep.Logradouro = postmon.Logradouro
						wecep.Bairro = postmon.Bairro
						b, err := json.Marshal(wecep)
						if err == nil {
							chResult <- Result{Body: b}
							cancel()
						}
					}

				case "republicavirtual":
					var repub = models.RepublicaVirtual{}
					err := json.Unmarshal(body, &repub)
					if err == nil {
						wecep.Cidade = repub.Cidade
						wecep.Uf = repub.Uf
						wecep.Logradouro = repub.Logradouro
						wecep.Bairro = repub.Bairro
						b, err := json.Marshal(wecep)
						if err == nil {
							chResult <- Result{Body: b}
							cancel()
						}
					}
				}
			}
			return

		}(cancel, source, endpoint, chResult)
	}

	select {
	case result := <-chResult:
		ristretto.Set(cep, string(result.Body))
		w.WriteHeader(http.StatusOK)
		w.Write(result.Body)
		return
	case <-time.After(time.Duration(4) * time.Second):
		cancel()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(config.JsonDefault))
	return
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}

func checkCep(cep string) error {
	re := regexp.MustCompile(`[^0-9]`)
	formatedCEP := re.ReplaceAllString(cep, `$1`)

	if len(formatedCEP) < 8 {
		return errors.New(`{"msg":"error cep tem que ser valido"}`)
	}

	return nil
}
