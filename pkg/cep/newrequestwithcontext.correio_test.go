package cep

import (
	"context"
	"testing"
)

func TestNewRequestWithContextCorreio(t *testing.T) {
	type args struct {
		ctx      context.Context
		cancel   context.CancelFunc
		cep      string
		source   string
		method   string
		endpoint string
		payload  string
		chResult chan<- Result
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewRequestWithContextCorreio(tt.args.ctx, tt.args.cancel, tt.args.cep, tt.args.source, tt.args.method, tt.args.endpoint, tt.args.payload, tt.args.chResult)
		})
	}
}
