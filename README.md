# gocep

Um simples pacote para buscar ceps em bases públicas na internet utilizando *concorrência*.
Atualizamos para buscar não somente de bases públicas como também busca do correios que é chamadas SOAPs e busca também de uma base que encontra-se no [ceps github](https://raw.githubusercontent.com/jeffotoni/api.cep/master/v1/cep/) em raw.

Você também pode extender e buscar sua própria base de Ceps se desejar consultar em sua própria base de dados.

Está configurado para buscar em: 
 - viacep 
 - Postmon cep 
 - Republicavirtual 
 - Correio 
 - github Raw Cep
 - Cdn api cep

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

O server é extremamente rápido, e usa cache em memória ele está configurado para 2G de Ram, caso queira alterar está tudo bonitinho no /config.

Gocep também poderá ser usado como Lib, ou seja você irá conseguir fazer um import em seu pkg/searchcep e fazer a chamada direto do seu método em seu código.

#### Usar como Lib
```go

package main

import (
	"fmt"
	"github.com/jeffotoni/gocep/pkg/cep"
)

func main() {

	result, err := cep.Search("6233903")
	fmt.Println(err)
	fmt.Println(result)
}

```

Ou se preferir for criar seu próprio serviço e sua api basta fazer como exemplo abaixo:
Da uma conferida em alguns examplos aqui de como integrar a lib gocep em seu projeto: [exemplos](https://github.com/jeffotoni/gocep/tree/master/examples)

#### Criando seu próprio WebServer usando gocep
```bash
package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/jeffotoni/gocep/pkg/cep"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cep/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type", "application/json")
		cepstr := strings.Split(r.URL.Path[1:], "/")[1]
		if len(cepstr) != 8 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := cep.Search(cepstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8080"))
}

```

Você pode fazer seu próprio build usando Go, ou você poderá utilizar docker-compose. 
O server irá funcionar na porta 8080, mas caso queira alterar basta ir na pasta /config.

Para subir o serviço para seu Servidor ou sua máquina local basta compilar, e a porta 8080 será aberta para consumir o endpoint /api/v1/{cep}

Tudo muito legal não é ?? ❤️😍😍

### Install gocep

Caso queira utilizar ele como serviço, basta baixa-lo ou usar o docker para utilizado.

### linux bash
```bash
$ git clone https://github.com/jeffotoni/gocep
$ cd gocep
$ CGO_ENABLED=0 go build -ldflags="-s -w" 
$ ./gocep
$ 2020/04/21 12:56:46 Port: 0.0.0.0:8080

```

### docker e docker-compose

Deixei um script para facilitar a criação de sua imagem, todos os arquivos estão na raiz, docker-compose.yaml, Dockerfile tudo que precisa para personalizar ainda mais se precisar.

```bash
version: '3.5'

services:
  gocep:
    image: jeffotoni/gocep
    container_name: gocep
    hostname: gocep
    domainname: gocep.local.com
    environment:
      - "TZ=America/Sao_Paulo"
      - "API_ENV=prod"
    networks:
        guulawork:
           aliases:
              - gocep.local.com
    ports:
      - 8080:8080
    restart: always

networks:
  guulawork:
      driver: bridge

```

Ao rodar o script ele irá fazer pull da imagem que encontra-se no hub.docker.
```bash

$ make compose

```

### Listando service
```bash
$ docker-compose ps
Creating gocep ... done
Name    Command   State           Ports         
------------------------------------------------
gocep   /gocep    Up      0.0.0.0:8080->8080/tcp
-e Generated Run docker-compose [ok] 

```

### Executando sua API
```bash

$ curl -i -XGET http://localhost:8080/api/v1/08226021
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 19 Feb 2023 13:15:03 GMT
Content-Length: 112
{
	"cidade":"São Paulo",
	"uf":"SP",
	"logradouro":"18 de Abril",
	"bairro":"Cidade Antônio Estevão de Carvalho"
}

```

Temos uma estrutura padrão de retorno do JSON.

#### Struct Go
```go

type WeCep struct {
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
}

```

#### Saida Json
```json

	{
		"cidade":"",
		"uf":"",
		"logradouro":"",
		"bairro":""
	}

```




