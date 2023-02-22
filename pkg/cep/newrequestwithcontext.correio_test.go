package cep

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jeffotoni/gocep/config"
)

// Esse exemplo faz um requisição para a API dos correios
func ExampleNewRequestWithContextCorreio() {
	ctx, cancel := context.WithCancel(context.Background())
	cep := "01001000"
	source := "correio"
	method := http.MethodPost
	endpoint := "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente"
	payload := `<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
	<x:Body>
		<cli:consultaCEP>
			<cep>%s</cep>
		</cli:consultaCEP>
	</x:Body>
</x:Envelope>`
	chResult := make(chan Result)

	go NewRequestWithContextCorreio(ctx, cancel, cep, source, method, endpoint, payload, chResult)

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

// go test -run ^TestNewRequestWithContextCorreio$ -v
func TestNewRequestWithContextCorreio(t *testing.T) {
	type args struct {
		ctx      context.Context
		cancel   context.CancelFunc
		cep      string
		source   string
		method   string
		endpoint string
		payload  string
		chResult chan Result
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test_new_request_with_context_correio_",
			args: args{
				cep:      "01001000",
				source:   "correio",
				method:   "POST",
				endpoint: "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente",
				payload: `<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
				<x:Body>
					<cli:consultaCEP>
						<cep>%s</cep>
					</cli:consultaCEP>
				</x:Body>
			</x:Envelope>`,
				chResult: make(chan Result),
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"Praça da Sé","bairro":"Sé"}`,
		},
		{name: "test_new_request_with_context_correio_",
			args: args{
				cep:      "0",
				source:   "correio",
				method:   "POST",
				endpoint: "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente",
				payload: `<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
				<x:Body>
					<cli:consultaCEP>
						<cep>%s</cep>
					</cli:consultaCEP>
				</x:Body>
			</x:Envelope>`,
				chResult: make(chan Result),
			},
			want: `{"cidade":"","uf":"","logradouro":"","bairro":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())

			go func(cancel context.CancelFunc, cep, source, method, endpoint, payload string, chResult chan<- Result) {
				NewRequestWithContextCorreio(ctx, cancel, cep, source, method, endpoint, payload, chResult)
			}(cancel, tt.args.cep, tt.args.source, tt.args.method, tt.args.endpoint, tt.args.payload, tt.args.chResult)

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
