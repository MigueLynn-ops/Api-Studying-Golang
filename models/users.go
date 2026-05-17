package models

type Users struct {
	ID       int8   `json:"id"`
	Username string `json:"username"`
	CPF      string `json:"cpf"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *Users) Prepare() error {
	// Chamar o validate e format
	return nil
}
func (u *Users) Validate() error {
	// Verificar se os campos não estão vazios, etc.
	// Validar o CPF
	return nil
}
func (u *Users) Format(step string) error {
	// Formatar as strings, por exemplo, remover espaços extras, etc.
	// Aplicar hash na senha
	return nil
}
