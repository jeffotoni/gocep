package util

import (
	"errors"
	"regexp"
)

func CheckCep(cep string) error {
	re := regexp.MustCompile(`[^0-9]`)
	formatedCEP := re.ReplaceAllString(cep, `$1`)

	if len(formatedCEP) < 8 {
		return errors.New(`{"msg":"error cep tem que ser valido"}`)
	}

	return nil
}
