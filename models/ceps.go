package models

// Default
type WeCep struct {
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
}

//viacep
type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

//postmon
type PostMon struct {
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Logradouro string `json:"logradouro"`
	EstadoInfo struct {
		AreaKm2    string `json:"area_km2"`
		CodigoIbge string `json:"codigo_ibge"`
		Nome       string `json:"nome"`
	} `json:"estado_info"`
	Cep        string `json:"cep"`
	CidadeInfo struct {
		AreaKm2    string `json:"area_km2"`
		CodigoIbge string `json:"codigo_ibge"`
	} `json:"cidade_info"`
	Estado string `json:"estado"`
}

//republicavirtual
type RepublicaVirtual struct {
	Resultado      string `json:"resultado"`
	ResultadoTxt   string `json:"resultado_txt"`
	Uf             string `json:"uf"`
	Cidade         string `json:"cidade"`
	Bairro         string `json:"bairro"`
	TipoLogradouro string `json:"tipo_logradouro"`
	Logradouro     string `json:"logradouro"`
}

type Correio struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                string `xml:",chardata"`
		ConsultaCEPResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text         string `xml:",chardata"`
				Bairro       string `xml:"bairro"`
				Cep          string `xml:"cep"`
				Cidade       string `xml:"cidade"`
				Complemento2 string `xml:"complemento2"`
				End          string `xml:"end"`
				Uf           string `xml:"uf"`
			} `xml:"return"`
		} `xml:"consultaCEPResponse"`
	} `xml:"Body"`
}
