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
