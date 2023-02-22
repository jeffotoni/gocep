package cep

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jeffotoni/gocep/models"
)

// NewRequestWithContext responsavel em fazer buscas de forma concorrente em seus respectivos
// servidores
func NewRequestWithContext(ctx context.Context, cancel context.CancelFunc, cep, source, method,
	endpoint string, chResult chan<- Result) {
	if source == "cdnapicep" && len(cep) > 7 {
		cep = addHyphen(cep)
	}
	endpoint = fmt.Sprintf(endpoint, cep)
	req, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
	if err != nil {
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error io.ReadAll:", err)
		return
	}

	defer response.Body.Close()

	if len(string(body)) > 0 &&
		response.StatusCode == http.StatusOK {
		var wecep = &models.WeCep{}
		switch source {
		case "githubjeffotoni":
			githubjeffotoni(wecep, body, chResult, cancel)
		case "viacep":
			viacep(wecep, body, chResult, cancel)
		case "postmon":
			postmon(wecep, body, chResult, cancel)
		case "republicavirtual":
			republicavirtual(wecep, body, chResult, cancel)
		case "cdnapicep":
			println("cdnapicep")
			cdnapicep(wecep, body, chResult, cancel)
		case "brasilapi":
			brasilapi(wecep, body, chResult, cancel)
		}
	}
	return
}

func addHyphen(s string) string {
	n := len(s)
	if n <= 5 {
		return s
	}
	return s[:5] + "-" + s[5:]
}

func githubjeffotoni(wecep *models.WeCep, body []byte, chResult chan<- Result, cancel context.CancelFunc) {
	var githubjeff = models.GithubJeffotoni{}
	err := json.Unmarshal(body, &githubjeff)
	if err == nil {
		wecep.Cidade = githubjeff.Cidade
		wecep.Uf = githubjeff.Uf
		wecep.Logradouro = githubjeff.Logradouro
		wecep.Bairro = githubjeff.Bairro
		b, err := json.Marshal(wecep)
		if err == nil {
			chResult <- Result{Body: b}
			cancel()
		}
	}
}

func viacep(wecep *models.WeCep, body []byte, chResult chan<- Result, cancel context.CancelFunc) {
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
}

func postmon(wecep *models.WeCep, body []byte, chResult chan<- Result, cancel context.CancelFunc) {
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
}

func republicavirtual(wecep *models.WeCep, body []byte, chResult chan<- Result, cancel context.CancelFunc) {
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

func cdnapicep(wecep *models.WeCep, body []byte, chResult chan<- Result, cancel context.CancelFunc) {
	var cdnapi = models.CdnApiCep{}
	err := json.Unmarshal(body, &cdnapi)
	if err == nil {
		wecep.Cidade = cdnapi.City
		wecep.Uf = cdnapi.State
		wecep.Logradouro = cdnapi.Address
		wecep.Bairro = cdnapi.District
		b, err := json.Marshal(wecep)
		if err == nil {
			chResult <- Result{Body: b}
			cancel()
		}
	}
}

func brasilapi(wecep *models.WeCep, body []byte, chResult chan<- Result, cancel context.CancelFunc) {
	var brasilapi = models.BrasilAPI{}
	err := json.Unmarshal(body, &brasilapi)
	if err == nil {
		wecep.Cidade = brasilapi.City
		wecep.Uf = brasilapi.State
		wecep.Logradouro = brasilapi.Street
		wecep.Bairro = brasilapi.Neighborhood
		b, err := json.Marshal(wecep)
		if err == nil {
			chResult <- Result{Body: b}
			cancel()
		}
	}
}
