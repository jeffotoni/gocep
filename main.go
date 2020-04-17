package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type End struct {
	Source string
	Url    string
}

type Result struct {
	Body []byte
	sync.Mutex
}

var endpoints = []End{
	{"viacep", "https://viacep.com.br/ws/%s/json/"},
	{"postmon", "https://api.postmon.com.br/v1/cep/%s"},
	{"republicavirtual", "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json"},
}

func main() {
	//var mux sync.Mutex
	mux := http.NewServeMux()

	// {cep:[0-9]{8}}
	mux.HandleFunc("/api/v1/", SearchCep)
	mux.HandleFunc("/api/v1", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    ":8084",
		Handler: mux,
	}
	log.Println("Port:", 8084)
	log.Fatal(server.ListenAndServe())
}

func SearchCep(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	cep := strings.Split(r.URL.Path[2:], "/")[2]
	if err := isValidCEP(cep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var result Result
	for _, e := range endpoints {
		endpoint := fmt.Sprintf(e.Url, cep)
		go func(cancel context.CancelFunc, endpoint string, result *Result) {

			req, err := http.NewRequest("GET", endpoint, nil)
			if err != nil {
				return
			}

			req = req.WithContext(ctx)
			response, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Error:", err)
				return
			}

			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Error ioutil.ReadAll:", err)
				return
			}

			if len(string(body)) > 0 &&
				response.StatusCode == http.StatusOK {
				result.Lock()
				result.Body = body
				result.Unlock()
				cancel()
				println("Encontrei o CEP.. propagando cancel de goroutines...:")
				println(endpoint)
				return
			}
		}(cancel, endpoint, &result)
	}

	//go func() {
	select {
	case <-ctx.Done():
		fmt.Println("body:", string(result.Body))
		fmt.Println("Gracefully exit")
		fmt.Println(ctx.Err())

		w.WriteHeader(http.StatusOK)
		w.Write(result.Body)
		return
	case <-time.After(time.Duration(5) * time.Second):
		cancel()
	}
	//}()

	w.WriteHeader(http.StatusNoContent)
	return
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}

func isValidCEP(cep string) error {
	re := regexp.MustCompile(`[^0-9]`)
	formatedCEP := re.ReplaceAllString(cep, `$1`)

	if len(formatedCEP) < 8 {
		return errors.New("Cep deve conter apenas números e no minimo 8 digitos")
	}

	return nil
}
