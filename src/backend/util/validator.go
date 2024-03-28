package util

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// custom validators

type CPFValidator struct {
	Validate *validator.Validate
}

var ValidateNumbers validator.Func = func(field validator.FieldLevel) bool {
	if field.Field().Kind() != reflect.Slice {
		return false
	}

	slice := field.Field()
	if slice.Len() != MAX_USER_GAMBLE_NUMBERS {
		return false
	}
	for i := 0; i < slice.Len(); i++ {
		if slice.Index(i).Kind() != reflect.Int {
			return false
		}
		if slice.Index(i).Int() < MIN_GAMBLE_RANGE {
			return false
		}
		if slice.Index(i).Int() > MAX_GAMBLE_RANGE {
			return false
		}
	}
	return true
}

var ValidateCPF validator.Func = func(field validator.FieldLevel) bool {
	cpf := strings.TrimSpace(field.Field().String())

	if len(cpf) != 11 {
		return false
	}

	// Calculate CPF verifier digits
	var sum int
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	if sum%11 < 2 {
		sum = 0
	} else {
		sum = 11 - sum%11
	}
	if sum != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	if sum%11 < 2 {
		sum = 0
	} else {
		sum = 11 - sum%11
	}
	if sum != int(cpf[10]-'0') {
		return false
	}

	return true
}
