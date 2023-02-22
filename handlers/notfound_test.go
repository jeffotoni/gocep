package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -run ^TestNotFound'$ -v
func TestNotFound(t *testing.T) {
	type args struct {
		method string
		target string
	}
	tests := []struct {
		name string
		want int //statuscode
	}{
		{"test_not_found_", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

			rr := httptest.NewRecorder()

			NotFound(rr, req)

			if rr.Code != tt.want {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.want)
			}

		})
	}
}
