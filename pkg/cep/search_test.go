package cep

import (
	"testing"
)

// go test -run ^TestSearch$ -v
func TestSearch(t *testing.T) {
	type args struct {
		cep string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want2   string
		wantErr bool
	}{
		{
			name:    "test_search_",
			args:    args{"08226024"},
			want:    `{"cidade":"São Paulo","uf":"SP","logradouro":"Rua Esperança","bairro":"Cidade Antônio Estevão de Carvalho"}`,
			want2:   `{"cidade":"São Paulo","uf":"SP","logradouro":"Esperança","bairro":"Cidade Antônio Estevão de Carvalho"}`,
			wantErr: false, // (err != nill) => false
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.cep)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && got != tt.want2 {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
