package util

import (
	"errors"
	"regexp"
)

var (
	cepRE = regexp.MustCompile(`^\d{8}$`)
)

func CheckCep(cep string) error {
	if cepRE.MatchString(cep) {
		return nil
	}
	return errors.New(`{"msg":"error cep tem que ser valido"}`)
}
