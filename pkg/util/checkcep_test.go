package util

import (
	"fmt"
	"testing"
)

// Esse exemplo passa um cep inválido para função CheckCep
func ExampleCheckCep() {
	err := CheckCep("não é um cep")
	fmt.Println(err)
	// Output: {"msg":"error cep tem que ser valido"}
}

// go test -run ^TestCheckCep'$ -v
func TestCheckCep(t *testing.T) {
	type args struct {
		cep string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test_chekcep_", args{"08226021"}, false},
		{"test_chekcep_", args{"01010001"}, false},
		{"test_chekcep_", args{"01010900"}, false},
		{"test_chekcep_", args{"xxxxxxxxx"}, true},
		{"test_chekcep_", args{"1234567"}, true},
		{"test_chekcep_", args{"123456789"}, true},
		{"test_chekcep_", args{"abc12345"}, true},
		{"test_chekcep_", args{"#$%&*^@"}, true},
		{"test_chekcep_", args{""}, true},
		{"test_chekcep_", args{"      "}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckCep(tt.args.cep); (err != nil) != tt.wantErr {
				t.Errorf("CheckCep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
