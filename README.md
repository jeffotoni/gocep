# gocep
[![GoDoc](https://godoc.org/gocep?status.svg)](https://godoc.org/gocep) [![Github Release](https://img.shields.io/github/v/release/jeffotoni/gocep?include_prereleases)](https://img.shields.io/github/v/release/jeffotoni/gocep) [![CircleCI](https://dl.circleci.com/status-badge/img/gh/jeffotoni/github.com/jeffotoni/gocep/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/jeffotoni/github.com/jeffotoni/gocep/tree/master) [![Go Report](https://goreportcard.com/badge/gocep)](https://goreportcard.com/badge/gocep) [![License](https://img.shields.io/github/license/jeffotoni/gocep)](https://img.shields.io/github/license/jeffotoni/gocep) ![CircleCI](https://img.shields.io/circleci/build/github/jeffotoni/github.com/jeffotoni/gocep/master) ![Coveralls](https://img.shields.io/coverallsCoverage/github/jeffotoni/gocep)
<!-- 
[![Vulnerabilities](https://snyk.io/test/github/jeffotoni/github.com/jeffotoni/gocep/badge.svg)](https://snyk.io/test/github/jeffotoni/gocep) -->

Um simples pacote para buscar ceps em bases públicas na internet utilizando *concorrência*.
Atualizamos para buscar não somente de bases públicas como também busca do correios que é chamadas SOAPs e busca também de uma base que encontra-se no [ceps github](https://raw.githubusercontent.com/jeffotoni/api.cep/master/v1/cep/) em raw.

Você também pode extender e buscar sua própria base de Ceps se desejar consultar em sua própria base de dados.

Está configurado para buscar em: 
 - viacep 
 - Postmon cep 
 - Republicavirtual 
 - Correio 
 - github Raw Cep (jeffotoni)
 - Cdn api cep
 - Brasil Api

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

O server é extremamente rápido, e usa cache em memória ele está configurado para 2G de Ram, caso queira alterar está tudo bonitinho no /config.

#### Fazendo chamadas do gocep em outras Langs

Da uma conferida em alguns examplos aqui de como fazer chamadas do gocep em diversas linguagens:
 - nodejs
 - python
 - php
 - javascript
 - go lib
 - go server
 - go client
 - rust
 - C
 - C++

[exemplos](https://github.com/jeffotoni/gocep/tree/master/examples)

Você pode fazer seu próprio build usando Go, ou você poderá utilizar docker-compose. 
O server irá funcionar na porta 8080, mas caso queira alterar basta ir na pasta /config.

Para subir o serviço para seu Servidor ou sua máquina local basta compilar, e a porta 8080 será aberta para consumir o endpoint /v1/cep/{cep}

Tudo muito legal não é ?? ❤️😍😍

#### Install gocep

Caso queira utilizar ele como serviço, basta baixa-lo ou usar o docker para utilizado.

#### linux bash
```bash
$ git clone https://gocep
$ cd gocep
$ CGO_ENABLED=0 go build -ldflags="-s -w" 
$ ./gocep
$ 2020/04/21 12:56:46 Port: 0.0.0.0:8080

```

#### docker e docker-compose

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

#### Listando service
```bash
$ docker-compose ps
Creating gocep ... done
Name    Command   State           Ports         
------------------------------------------------
gocep   /gocep    Up      0.0.0.0:8080->8080/tcp
-e Generated Run docker-compose [ok] 

```

#### Executando sua API
```bash

$ curl -i -XGET http://localhost:8080/v1/cep/08226021
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

#### Docker

Também poderá usar o Docker se desejar

```bash
$ docker run --name gocep --rm -p 8080:8080 jeffotoni/gocep:latest
2023/02/19 17:12:03 Server Run Port 0.0.0.0:8080
2023/02/19 17:12:03 /v1/cep/:cep

$ curl -i -XGET http://localhost:8080/v1/cep/08226021
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

#### Usar como Lib

Gocep também poderá ser usado como Lib, ou seja você irá conseguir fazer um import em seu pkg/searchcep e fazer a chamada direto do seu método em seu código.

```go

package main

import (
	"fmt"
	"github.com/jeffotoni/gocep/pkg/cep"
)

func main() {
	result, wecep, err := cep.Search("6233903")
	fmt.Println(err)
	fmt.Println(result) // json
	fmt.Println(wecep) // object WeCep
}

```

Ou se preferir for criar seu próprio serviço em Go e sua api basta fazer como exemplo abaixo:

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

		result, wecep, err := cep.Search(cepstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result))
			return
		}

		if !cep.ValidCep(wecep) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8080"))
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

