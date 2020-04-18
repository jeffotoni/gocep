package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jeffotoni/gocep/models"
	//"golang.org/x/sync/singleflight"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type End struct {
	Source string
	Url    string
}

type Result struct {
	Body []byte
}

var endpoints = []End{
	{"viacep", "https://viacep.com.br/ws/%s/json/"},
	{"postmon", "https://api.postmon.com.br/v1/cep/%s"},
	{"republicavirtual", "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json"},
}

var chResult = make(chan Result, len(endpoints))

var (
	Port = ":8084"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/", SearchCep)
	mux.HandleFunc("/api/v1", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    Port,
		Handler: mux,
	}

	log.Println("Port:", Port)
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

	//var requestGroup singleflight.Group

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, e := range endpoints {
		endpoint := fmt.Sprintf(e.Url, cep)
		go func(cancel context.CancelFunc, endpoint string, chResult chan<- Result) {
			//res2, err2, shared2 := requestGroup.Do("singleflight", func() (interface{}, error) {
			req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
			if err != nil {
				//return []byte(``), err
				return
			}

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				//return []byte(``), err
				return
			}

			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Error ioutil.ReadAll:", err)
				//return []byte(``), err
				return
			}

			if len(string(body)) > 0 &&
				response.StatusCode == http.StatusOK {
				var wecep = models.WeCep{}
				//println(e.Source)
				switch e.Source {
				case "viacep":
					var viacep = models.ViaCep{}
					err := json.Unmarshal(body, &viacep)
					if err == nil {
						wecep.Cidade = viacep.Localidade
						wecep.Uf = viacep.Uf
						wecep.Logradouro = viacep.Logradouro
						wecep.Bairro = viacep.Bairro
						b, err := json.Marshal(&wecep)
						if err == nil {
							//return b, err
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
						b, err := json.Marshal(&wecep)
						if err == nil {
							//return b, err
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
						b, err := json.Marshal(&wecep)
						if err == nil {
							//								return b, err
							chResult <- Result{Body: b}
							cancel()
						}
					}
				}
			}
			return
			//return []byte(``), errors.New(`nao encontramos o resultado`)
			//return string(responseData), err
			//})

			// if err2 != nil {
			// 	http.Error(w, err2.Error(), http.StatusInternalServerError)
			// 	return
			// }

			// result2 := res2.(string)
			// fmt.Println("shared = ", shared2)
			// fmt.Fprintf(w, "%q", result2)

		}(cancel, endpoint, chResult)
	}

	select {
	//case <-ctx.Done():
	case result := <-chResult:
		w.WriteHeader(http.StatusOK)
		w.Write(result.Body)
		return
	case <-time.After(time.Duration(5) * time.Second):
		cancel()
	}

	w.WriteHeader(http.StatusNoContent)
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
