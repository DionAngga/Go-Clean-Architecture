package controller

import (
	"crud/entity"
	"crud/repository"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users entity.User
	json.NewDecoder(r.Body).Decode(&users)
	repository.DB.Create(users)
	json.NewEncoder(w).Encode(users)
}
