package models

type end struct {
	Method string
	Source string
	Url    string
	Body   string
}

var Endpoints = []end{
	//{"GET", "apicep", "https://cdn.apicep.com/file/apicep/%s.json", ""},
	{"GET", "githubjeffotoni", "https://raw.githubusercontent.com/jeffotoni/api.cep/master/v1/cep/%s", ""},
	{"GET", "viacep", "https://viacep.com.br/ws/%s/json/", ""},
	{"GET", "postmon", "https://api.postmon.com.br/v1/cep/%s", ""},
	{"GET", "republicavirtual", "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json", ""},
	{"POST", "correio", "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente",
		`<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
    <x:Body>
        <cli:consultaCEP>
            <cep>%s</cep>
        </cli:consultaCEP>
    </x:Body>
</x:Envelope>`}}
