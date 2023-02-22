package handler

import (
	"net/http"
	"strings"

	"github.com/jeffotoni/gocep/pkg/cep"
	"github.com/jeffotoni/gocep/pkg/util"
)

func SearchCep(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	validEndpoint := strings.Split(r.URL.Path, "/")
	if len(validEndpoint) > 4 {
		w.WriteHeader(http.StatusFound)
		return
	}

	cepstr := strings.Split(r.URL.Path[2:], "/")[2]
	if err := util.CheckCep(cepstr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, wecep, err := cep.Search(cepstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}

	if !cep.ValidCep(wecep) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}
