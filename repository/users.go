package repository

import (
	"Api-Aula_1/models"
	"database/sql"
)

type UserRepository struct {
	bd *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{bd: db}
}

func (u UserRepository) Create(user models.Users) (int8, error) {
	query := `INSERT INTO aulago.users (
	name_user, 
	email, 
	password_user,
	cpf
	) VALUES (?, ?, ?, ?)`

	statment, err := u.bd.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Username, user.Email, user.Password, user.CPF)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int8(uint64(lastID)), nil
}
