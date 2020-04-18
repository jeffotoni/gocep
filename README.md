# gocep

Um simples pacote para buscar ceps em bases publicas na internet.

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

O server é extremamente rápido, e usa cache em memória ele está configurado para 2G de Ram, caso queira alterar está tudo bonitinho no /config.

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

Você pode fazer seu próprio build usando Go, ou você poderá utilizar docker-compose. O server irá funcionar na porta 8084, mas caso queira alterar basta ir na pasta /config.

## Start gocep linux bash
```bash
$ go build -ldflags="-s -w" 
$ ./gocep
```

## Start gocep Docker e docker-compose
```bash
$ sh deploy.gocep.sh
$ docker-compose ps
$ docker-compose logs -f gocep
```

## Executando sua API
```bash
$ curl -i http://localhost:8084/api/v1/08226021
```

## out
```bash

$ {"cidade":"São Paulo","uf":"SP","logradouro":"18 de Abril","bairro":"Cidade Antônio Estevão de Carvalho"}

```






