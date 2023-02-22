package gocache

import (
	"reflect"
	"testing"
	"time"

	gcache "github.com/patrickmn/go-cache"
)

// go test -run ^TestRun'$ -v
func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		want    *gcache.Cache
		wantErr bool
	}{
		{name: "test_run_", want: Run(), wantErr: false},
		{name: "test_run_", want: gcache.New(24*time.Hour, 24*time.Hour), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -run ^TestSetTTL'$ -v
func TestSetTTL(t *testing.T) {
	TestRun(t)
	type args struct {
		key   string
		value string
		ttl   time.Duration
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_setTTL_",
			args: args{
				key:   `08226021`,
				value: `{"cidade":"São Paulo","uf":"SP","logradouro":"18 de Abril","bairro":"Cidade Antônio Estevão de Carvalho"}`,
				ttl:   time.Duration(5) * time.Second,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetTTL(tt.args.key, tt.args.value, tt.args.ttl); got != tt.want {
				t.Errorf("SetTTL() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -run ^TestGet'$ -v
func TestGet(t *testing.T) {
	TestRun(t)
	TestSetTTL(t)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_get_",
			args: args{
				key: `08226021`,
			},
			want: `{"cidade":"São Paulo","uf":"SP","logradouro":"18 de Abril","bairro":"Cidade Antônio Estevão de Carvalho"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
