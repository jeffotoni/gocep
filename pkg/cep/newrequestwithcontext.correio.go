package cep

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/jeffotoni/gocep/models"
)

// NewRequestWithContextCorreio é responsável por fazer buscas de forma concorrente na API dos correios
func NewRequestWithContextCorreio(ctx context.Context, cancel context.CancelFunc, cep, source, method, endpoint, payload string, chResult chan<- Result) {
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
}
