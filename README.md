# gocep

Um simples pacote para buscar ceps em bases publicas na internet.

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

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

## Start gocep linux bash
```bash
$ go build -ldflags="-s -w" 
$ ./gocep
```

## Start gocep Docker
```bash
$ docker build -f Dockerfile -t jeffotoni/gocep .
$ 
```

