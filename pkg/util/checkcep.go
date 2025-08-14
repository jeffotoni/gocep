package util

import (
	"errors"
	"regexp"
)

var (
	cepRE = regexp.MustCompile(`^\d{8}$`)
)

// CheckCep verifica se o cep informado é um CEP válido, ou seja, uma string composta por 8 números "00000000",
// e retorna um erro caso contrário.
func CheckCep(cep string) error {
	if cepRE.MatchString(cep) {
		return nil
	}
	return errors.New(`{"msg":"error cep tem que ser valido"}`)
}
