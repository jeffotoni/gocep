package util

import "testing"

func TestCheckCep(t *testing.T) {
	type args struct {
		cep string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test_chekcep_", args{"08226021"}, false},
		{"test_chekcep_", args{"01010001"}, false},
		{"test_chekcep_", args{"01010900"}, false},
		{"test_chekcep_", args{"xxxxxxxxx"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckCep(tt.args.cep); (err != nil) != tt.wantErr {
				t.Errorf("CheckCep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
