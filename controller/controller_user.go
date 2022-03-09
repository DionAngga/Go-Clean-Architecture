package controller

import (
	"crud/entity"
	"crud/usecase"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	//LoginUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetEmailUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) Controller {
	return &controller{usecase}
}

func (u *controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []entity.User
	u.usecase.FindAllUser(&users)
	json.NewEncoder(w).Encode(users)
}

func (u *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	u.usecase.FindUser(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func (u *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	u.usecase.CreateUser(&user)
	json.NewEncoder(w).Encode(user)
}

func (u *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	u.usecase.FindUser(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	u.usecase.UpdateUser(&user)
	json.NewEncoder(w).Encode(user)
}

func (u *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	u.usecase.DeleteUser(&user, params["id"])
	json.NewEncoder(w).Encode("data sukses terhapus")
}

// func (u *controller) LoginUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user entity.User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	u.usecase.Login(&user)
// 	json.NewEncoder(w).Encode(user)
// }

func (u *controller) GetEmailUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *entity.User
	json.NewDecoder(r.Body).Decode(&user)
	aa, ee := u.usecase.FindUserByEmail(user)
	if user.ID == 0 {
		json.NewEncoder(w).Encode(ee)
	} else {
		json.NewEncoder(w).Encode(aa)
	}
}
