package util

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/go-playground/validator/v10"
// )

// func TestValidCPF(t *testing.T) {
// 	v := validator.New()

// 	tests := []struct {
// 		Cpf string
// 	}{
// 		{Cpf: "14041364019"},
// 		{Cpf: "52523685035"},
// 		{Cpf: "09830834018"},
// 		{Cpf: "25852090085"},
// 		{Cpf: "49965314012"},
// 	}

// 	for _, test := range tests {
// 		v.Struct(test)
// 		got := v.RegisterValidation("cpf", ValidateCPF)
// 		fmt.Println(got)

// 		if got != nil {

// 			t.Errorf("Cpf: %s is not valid!", test.Cpf)
// 		}
// 	}
// }

// func TestUnvalidCPF(t *testing.T) {
// 	v := validator.New()

// 	tests := []struct {
// 		Cpf string
// 	}{
// 		{Cpf: "12345678900"}, // Invalid CPF (last digit changed)
// 		{Cpf: "98765432101"}, // Invalid CPF (last digit changed)
// 		{Cpf: "00000000000"}, // Invalid CPF (all digits the same)
// 		{Cpf: "99999999999"}, // Invalid CPF (all digits the same)
// 		{Cpf: "87654321098"}, // Invalid CPF (reversed digits)
// 	}

// 	for _, test := range tests {
// 		v.Struct(test)
// 		got := v.RegisterValidation("cpf", ValidateCPF)
// 		fmt.Println(got)
// 		if got == nil {
// 			t.Errorf("Cpf: %s should be invalid", test.Cpf)
// 		}
// 	}
// }

// func TestWrongLengthCPF(t *testing.T) {

// }

// func TestEmptyCPF(t *testing.T) {

// }
