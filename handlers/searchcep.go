package handler

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
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
		endpoint := e.Url
		source := e.Source
		method := e.Method
		payload := e.Body
		go func(cancel context.CancelFunc, cep, method, source, endpoint, payload string, chResult chan<- Result) {

			if source == "correio" {
				NewRequestWithContextCorreio(ctx, cancel, cep, source, method, endpoint, payload, chResult)
			} else {
				NewRequestWithContext(ctx, cancel, cep, source, method, endpoint, chResult)
			}

		}(cancel, cep, method, source, endpoint, payload, chResult)
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

func NewRequestWithContextCorreio(ctx context.Context, cancel context.CancelFunc, cep, source, method, endpoint, payload string, chResult chan<- Result) {

	var err error
	payload = fmt.Sprintf(payload, cep)
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewReader([]byte(payload)))
	if err != nil {
		return
	}

	req.Header.Set("Content-type", "text/xml; charset=utf-8")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var wecep = &models.WeCep{}
	correio := new(models.Correio)
	err = xml.NewDecoder(response.Body).Decode(correio)
	if err == nil {
		c := correio.Body.ConsultaCEPResponse.Return
		wecep.Cidade = c.Cidade
		wecep.Uf = c.Uf
		wecep.Logradouro = c.End
		wecep.Bairro = c.Bairro
		b, err := json.Marshal(wecep)
		if err == nil {
			chResult <- Result{Body: b}
			cancel()
		}
	}

	return
}

func NewRequestWithContext(ctx context.Context, cancel context.CancelFunc, cep, source, method, endpoint string, chResult chan<- Result) {

	endpoint = fmt.Sprintf(endpoint, cep)
	req, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
	if err != nil {
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error ioutil.ReadAll:", err)
		return
	}

	defer response.Body.Close()

	if len(string(body)) > 0 &&
		response.StatusCode == http.StatusOK {
		var wecep = &models.WeCep{}
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
}

func checkCep(cep string) error {
	re := regexp.MustCompile(`[^0-9]`)
	formatedCEP := re.ReplaceAllString(cep, `$1`)

	if len(formatedCEP) < 8 {
		return errors.New(`{"msg":"error cep tem que ser valido"}`)
	}

	return nil
}
