package models

type end struct {
	Source string
	Url    string
}

var Endpoints = []end{
	{"viacep", "https://viacep.com.br/ws/%s/json/"},
	{"postmon", "https://api.postmon.com.br/v1/cep/%s"},
	{"republicavirtual", "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json"},
}
