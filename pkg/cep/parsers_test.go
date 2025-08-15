package cep

import (
	"reflect"
	"testing"

	"github.com/jeffotoni/gocep/models"
)

func TestGetParser(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want CepAPI
	}{
		{
			name: "GetParser cdnapicep",
			args: args{source: "cdnapicep"},
			want: &GenericParser{Source: "cdnapicep"},
		},
		{
			name: "GetParser viacep",
			args: args{source: "viacep"},
			want: &GenericParser{Source: "viacep"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetParser(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericParser_Parse(t *testing.T) {
	type fields struct {
		Source string
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.WeCep
		wantErr bool
	}{
		{
			name:   "Parse cdnapicep",
			fields: fields{Source: "cdnapicep"},
			args: args{
				body: []byte(`{"status":200,"code":"01001-000","state":"SP","city":"São Paulo","district":"Sé","address":"Praça da Sé - lado ímpar"}`),
			},
			want: &models.WeCep{
				Cidade:     "São Paulo",
				Uf:         "SP",
				Logradouro: "Praça da Sé - lado ímpar",
				Bairro:     "Sé",
			},
			wantErr: false,
		},
		{
			name:   "Parse viacep",
			fields: fields{Source: "viacep"},
			args: args{
				body: []byte(`{"cep": "01001-000", "logradouro": "Praça da Sé", "localidade": "São Paulo", "uf": "SP", "bairro": "Sé"}`),
			},
			want: &models.WeCep{
				Cidade:     "São Paulo",
				Uf:         "SP",
				Logradouro: "Praça da Sé",
				Bairro:     "Sé",
			},
			wantErr: false,
		},
		{
			name:   "Parse invalid json",
			fields: fields{Source: "viacep"},
			args: args{
				body: []byte(`invalid json`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GenericParser{
				Source: tt.fields.Source,
			}
			got, err := p.Parse(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenericParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenericParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
