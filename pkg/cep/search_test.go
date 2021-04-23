package cep

import "testing"

// go test -run ^TestSearch$ -v
func TestSearch(t *testing.T) {
	type args struct {
		cep string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test_search_", args{"08226021"},
			`{"cidade":"São Paulo","uf":"SP","logradouro":"18 de Abril","bairro":"Cidade Antônio Estevão de Carvalho"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.cep)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
