package cep

import (
	"context"
	"runtime"
	"time"

	"github.com/jeffotoni/gocep/config"
	"github.com/jeffotoni/gocep/models"
	"github.com/jeffotoni/gocep/service/gocache"
)

// Result representa a resposta da requisição em uma das APIs
type Result struct {
	Body []byte
	//WeCep *models.WeCep
}

// Search busca o cep informado de forma concorrente nas APIs definidas em [pkg/models/endpoints.go],
// retornando a primeira resposta(string em formato JSON) e um erro.
func Search(cep string) (string, error) {
	jsonCep := gocache.Get(cep)
	if len(jsonCep) > 0 {
		return jsonCep, nil
	}

	var chResult = make(chan Result, len(models.Endpoints))
	runtime.GOMAXPROCS(config.NumCPU)

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
		gocache.SetTTL(cep, string(result.Body), time.Duration(config.TTlCache)*time.Second)
		return string(result.Body), nil

	case <-time.After(time.Duration(config.TimeOutSearchCep) * time.Second):
		cancel()
	}
	return config.JsonDefault, nil
}
