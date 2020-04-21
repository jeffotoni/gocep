package handler

import (
	"github.com/jeffotoni/gocep/pkg/cep"
	"github.com/jeffotoni/gocep/pkg/util"
	"net/http"
	"strings"
)

func SearchCep(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	validEndpoint := strings.Split(r.URL.Path, "/")
	if len(validEndpoint) > 4 {
		w.WriteHeader(http.StatusFound)
		return
	}

	cep := strings.Split(r.URL.Path[2:], "/")[2]
	if err := util.CheckCep(cep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := cep.Search(cep)
	if err != nil {
		w.WriteHeader(code)
		w.Write(body)
		return
	}
	w.WriteHeader(code)
	w.Write(body)
	return
}
