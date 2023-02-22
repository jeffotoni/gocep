package cep

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jeffotoni/gocep/config"
)

// Esse exemplo faz um requisição para a API do viacep
func ExampleNewRequestWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cep := "01001000"
	source := "viacep"
	method := http.MethodGet
	endpoint := "https://viacep.com.br/ws/%s/json/"

	chResult := make(chan Result)

	go NewRequestWithContext(ctx, cancel, cep, source, method, endpoint, chResult)

	var result string
	select {
	case got := <-chResult:
		result = string(got.Body)
	case <-time.After(time.Duration(15) * time.Second):
		// tratar o erro, apenas a resposta está presente no chResult
	}
	fmt.Println(result)
	// Output: {"cidade":"São Paulo","uf":"SP","logradouro":"Praça da Sé","bairro":"Sé"}
}

// go test -run ^TestNewRequestWithContext$ -v
func TestNewRequestWithContext(t *testing.T) {
	type args struct {
		cep      string
		source   string
		method   string
		endpoint string
		chResult chan Result
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test_new_request_with_context_cdnapicep_",
			args: args{
				cep:      "01001000",
				source:   "cdnapicep",
				method:   "GET",
				endpoint: "https://cdn.apicep.com/file/apicep/%s.json",
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"Praça da Sé - lado ímpar","bairro":"Sé"}`,
		},
		{name: "test_new_request_with_context_githubjeffotoni_",
			args: args{
				cep:      "01001000",
				source:   "githubjeffotoni",
				method:   "GET",
				endpoint: "https://raw.githubusercontent.com/jeffotoni/api.cep/master/v1/cep/%s",
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"da Sé","bairro":"Sé"}`,
		},
		{name: "test_new_request_with_context_viacep_",
			args: args{
				cep:      "01001000",
				source:   "viacep",
				method:   "GET",
				endpoint: "https://viacep.com.br/ws/%s/json/",
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"Praça da Sé","bairro":"Sé"}`,
		},
		{name: "test_new_request_with_context_postmon_",
			args: args{
				cep:      "01001000",
				source:   "postmon",
				method:   "GET",
				endpoint: "https://api.postmon.com.br/v1/cep/%s",
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"Praça da Sé","bairro":"Sé"}`,
		},
		{name: "test_new_request_with_context_republicavirtual_",
			args: args{
				cep:      "01001000",
				source:   "republicavirtual",
				method:   "GET",
				endpoint: "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json",
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"da Sé","bairro":"Sé"}`,
		},
		{name: "test_new_request_with_context_brasilapi_",
			args: args{
				cep:      "01001000",
				source:   "brasilapi",
				method:   "GET",
				endpoint: "https://brasilapi.com.br/api/cep/v1/%s",
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"Praça da Sé","bairro":"Sé"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())

			go func(cancel context.CancelFunc, cep, source, method, endpoint string, chResult chan<- Result) {
				NewRequestWithContext(ctx, cancel, cep, source, method, endpoint, chResult)
			}(cancel, tt.args.cep, tt.args.source, tt.args.method, tt.args.endpoint, tt.args.chResult)

			select {
			case got := <-tt.args.chResult:
				if string(got.Body) != tt.want {
					t.Errorf("NewRequestWithContext() = %v, want %v", string(got.Body), tt.want)
				}
			case <-time.After(time.Duration(config.TimeOutSearchCep) * time.Second):
				t.Errorf("NewRequestWithContext() = %v, want %v", "timeout", tt.want)
			}
		})
	}
}
