package searchcep

import (
	"context"
	"github.com/jeffotoni/gocep/config"
	"github.com/jeffotoni/gocep/models"
	"github.com/jeffotoni/gocep/service/ristretto"
	"net/http"
	"time"
)

type Result struct {
	Body  []byte
	WeCep models.WeCep
}

var chResult = make(chan Result, len(models.Endpoints))

func Search(cep string) (*models.WeCep, error) {

	jsonCep := ristretto.Get(cep)
	if len(jsonCep) > 0 {
		return http.StatusOK, []byte(jsonCep)
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
		return result.WeCep, nil

	case <-time.After(time.Duration(4) * time.Second):
		cancel()
	}

	var wecep = &models.WeCep{}
	return wecep, nil
}
