package controller

import (
	"Api-Aula_1/models"
	"Api-Aula_1/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Ler os dados do corpo da requisição
	// Fazer o unmarshall dos dados para uma struct de usuário
	// Fazer os tratamentos necessários (validação, etc.)

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	var newUser models.Users
	if err := json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	log.Println(newUser)
	if err = newUser.Prepare("create"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, newUser)
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
