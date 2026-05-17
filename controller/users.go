package controller

import (
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Ler os dados do corpo da requisição
	// Fazer o unmarshall dos dados para uma struct de usuário
	// Fazer os tratamentos necessários (validação, etc.)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("User created successfully"))
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching user...")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User fetched successfully"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User updated successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User deleted successfully"))
}
