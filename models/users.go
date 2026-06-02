package models

import (
	"fmt"
	"strings"

	"Api-Aula_1/utils"

	"github.com/badoux/checkmail"
)

type Users struct {
	ID       int8   `json:"id"`
	Username string `json:"username"`
	CPF      string `json:"cpf"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *Users) Prepare(step string) error {
	// Chamar o validate e format
	if err := u.Validate(step); err != nil {
		return err
	}
	if err := u.Format(step); err != nil {
		return err
	}
	return nil
}

func (u *Users) Validate(step string) error {
	// Validar se o username não está vazio
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	// Validar se o email não está vazio
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	// Validar se o email tem um formato válido usando a biblioteca checkmail
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return fmt.Errorf("invalid email format")
	}
	// Validar o CPF usando a função ValidateCPF do pacote utils
	if err := utils.ValidateCPF(u.CPF); err != nil {
		return fmt.Errorf("invalid CPF format")
	}
	// Validar se a senha não está vazia (apenas na criação)
	if step == "create" && u.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (u *Users) Format(step string) error {
	// Formatar as strings, por exemplo, remover espaços extras, etc.
	// Aplicar hash na senha
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)

	return nil
}
