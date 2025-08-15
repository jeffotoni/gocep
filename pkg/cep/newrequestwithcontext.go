package cep

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
		log.Println("Error NewRequestWithContext:", err)
		return
	}

	fmt.Println(endpoint)

	response, err := httpClient.Do(req)
	if err != nil {
		// log.Println("Error httpClient:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error io.ReadAll:", err)
		return
	}

	if len(string(body)) > 0 &&
		response.StatusCode == http.StatusOK {
		parser := GetParser(source)
		if parser != nil {
			wecep, err := parser.Parse(body)
			if err == nil {
				b, err := json.Marshal(wecep)
				if err == nil {
					chResult <- Result{Body: b}
					cancel()
				}
			}
		}
	}
}

func addHyphen(s string) string {
	n := len(s)
	if n <= 5 {
		return s
	}
	return s[:5] + "-" + s[5:]
}
