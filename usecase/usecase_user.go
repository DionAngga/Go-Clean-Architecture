package usecase

import (
	"crud/entity"
	"crud/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []entity.User
	repository.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	repository.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	repository.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	repository.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	repository.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	repository.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("data sukses terhapus")
}
