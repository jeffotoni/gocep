package cep

import (
	"encoding/json"

	"github.com/jeffotoni/gocep/models"
)

type CepAPI interface {
	Parse(body []byte) (*models.WeCep, error)
}

type GenericParser struct {
	Source string
}

func (p *GenericParser) Parse(body []byte) (*models.WeCep, error) {
	wecep := &models.WeCep{}
	switch p.Source {
	case "cdnapicep":
		var cdnapi models.CdnApiCep
		if err := json.Unmarshal(body, &cdnapi); err != nil {
			return nil, err
		}
		wecep.Cidade = cdnapi.City
		wecep.Uf = cdnapi.State
		wecep.Logradouro = cdnapi.Address
		wecep.Bairro = cdnapi.District
	case "githubjeffotoni":
		var githubjeff models.GithubJeffotoni
		if err := json.Unmarshal(body, &githubjeff); err != nil {
			return nil, err
		}
		wecep.Cidade = githubjeff.Cidade
		wecep.Uf = githubjeff.Uf
		wecep.Logradouro = githubjeff.Logradouro
		wecep.Bairro = githubjeff.Bairro
	case "viacep":
		var viacep models.ViaCep
		if err := json.Unmarshal(body, &viacep); err != nil {
			return nil, err
		}
		wecep.Cidade = viacep.Localidade
		wecep.Uf = viacep.Uf
		wecep.Logradouro = viacep.Logradouro
		wecep.Bairro = viacep.Bairro
	case "postmon":
		var postmon models.PostMon
		if err := json.Unmarshal(body, &postmon); err != nil {
			return nil, err
		}
		wecep.Cidade = postmon.Cidade
		wecep.Uf = postmon.Estado
		wecep.Logradouro = postmon.Logradouro
		wecep.Bairro = postmon.Bairro
	case "republicavirtual":
		var repub models.RepublicaVirtual
		if err := json.Unmarshal(body, &repub); err != nil {
			return nil, err
		}
		wecep.Cidade = repub.Cidade
		wecep.Uf = repub.Uf
		wecep.Logradouro = repub.Logradouro
		wecep.Bairro = repub.Bairro
	case "brasilapi":
		var brasilapi models.BrasilAPI
		if err := json.Unmarshal(body, &brasilapi); err != nil {
			return nil, err
		}
		wecep.Cidade = brasilapi.City
		wecep.Uf = brasilapi.State
		wecep.Logradouro = brasilapi.Street
		wecep.Bairro = brasilapi.Neighborhood
	}
	return wecep, nil
}

func GetParser(source string) CepAPI {
	return &GenericParser{Source: source}
}
