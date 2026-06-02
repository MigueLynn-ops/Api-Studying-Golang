package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateCPF(cpf string) error {
	// Implementar a lógica de validação do CPF aqui
	// Verificar se o CPF tem 11 dígitos, calcular os dígitos verificadores, etc.
	cpf = onlyDigits(cpf)
	if len(cpf) != 11 {
		return fmt.Errorf("CPF must have 11 digits")
	}

	if !CheckAllEquals(cpf) {
		return fmt.Errorf("CPF cannot have all digits equal")
	}

	if !CalculateDv1(cpf) {
		return fmt.Errorf("Invalid DV1")
	}

	if !CalculateDv2(cpf) {
		return fmt.Errorf("Invalid DV2")
	}
	return nil
}

// Calcular o segundo dígito verificador (DV2) do CPF
func CalculateDv2(cpf string) bool {
	digits := strings.Split(cpf, "")
	if len(digits) < 11 {
		return false
	}
	soma := 0
	// multiplicar os 10 primeiros dígitos por pesos decrescentes de 11 a 2 e somar os resultados
	for i := 0; i < 10; i++ {
		num, err := strconv.Atoi(digits[i])
		if err != nil {
			return false
		}
		soma += num * (11 - i)
	}
	dv := (soma * 10) % 11
	if dv == 10 {
		dv = 0
	}
	return strconv.Itoa(dv) == digits[10]
}

// Calcular o primeiro dígito verificador (DV1) do CPF
func CalculateDv1(cpf string) bool {
	digits := strings.Split(cpf, "")
	if len(digits) < 10 {
		return false
	}
	soma := 0
	// multiplicar os 9 primeiros dígitos por pesos decrescentes de 10 a 2 e somar os resultados
	for i := 0; i < 9; i++ {
		num, err := strconv.Atoi(digits[i])
		if err != nil {
			return false
		}
		soma += num * (10 - i)
	}
	dv := (soma * 10) % 11
	if dv == 10 {
		dv = 0
	}
	return strconv.Itoa(dv) == digits[9]
}

// Verificar se todos os dígitos do CPF são iguais (o que é inválido)
func CheckAllEquals(cpf string) bool {
	if len(cpf) == 0 {
		return false
	}
	first := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != first {
			return true
		}
	}
	return false
}

// Função auxiliar para remover caracteres não numéricos do CPF
func onlyDigits(cpf string) string {
	// Implementar a lógica para remover caracteres não numéricos
	var b strings.Builder
	b.Grow(len(cpf))

	for _, r := range cpf {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
