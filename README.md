# gocep

Um simples pacote para buscar ceps em bases publicas na internet utilizando concorência.

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

O server é extremamente rápido, e usa cache em memória ele está configurado para 2G de Ram, caso queira alterar está tudo bonitinho no /config.

Gocep também poderá ser usado como Lib, ou seja você irá conseguir fazer um import em seu pkg/searchcep  e fazer a chamada direto do seu método em seu código.

## Usar como Lib
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

Você pode fazer seu próprio build usando Go, ou você poderá utilizar docker-compose. O server irá funcionar na porta 8084, mas caso queira alterar basta ir na pasta /config.

Para subir o serviço para seu Servidor ou sua máquina local basta compilar, e a porta 8084 será aberta para consumir o endpoint /api/v1/{cep}

## Start gocep linux bash
```bash

$ go build -ldflags="-s -w" 
$ ./gocep
$ 2020/04/21 12:56:46 Port: :8084

```

## Start gocep Docker e docker-compose
```bash

$ sh deploy.gocep.sh

$ docker-compose ps

Creating gocep ... done
Name    Command   State           Ports         
------------------------------------------------
gocep   /gocep    Up      0.0.0.0:8084->8084/tcp
-e Generated Run docker-compose [ok] 

$ docker-compose logs -f gocep
Attaching to gocep
gocep    | 2020/04/21 12:58:13 Port: :8084

```

## Executando sua API
```bash

$ curl -i http://localhost:8084/api/v1/08226021

```

## out
```bash

$ {"cidade":"São Paulo","uf":"SP","logradouro":"18 de Abril","bairro":"Cidade Antônio Estevão de Carvalho"}

```

Temos uma estrutura padrão de retorno do JSON.

## Struct Go
```go

type WeCep struct {
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
}

```

## Saida Json
```json

	{
		"cidade":"",
		"uf":"",
		"logradouro":"",
		"bairro":""
	}

```




