package controller

import (
	"Api-Aula_1/models"
	"Api-Aula_1/persistency"
	"Api-Aula_1/repository"
	"Api-Aula_1/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Ler os dados do corpo da requisição
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Fazer o unmarshall dos dados para uma struct de usuário
	var newUser models.Users
	if err := json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	log.Println(newUser)

	// Fazer os tratamentos necessários (validação, etc.)
	if err = newUser.Prepare("create"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	log.Println("User prepared successfully")

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	repository := repository.NewUserRepository(db)
	newUser.ID, err = repository.Create(newUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

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
