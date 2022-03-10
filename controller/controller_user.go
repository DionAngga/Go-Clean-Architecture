package controller

import (
	entity "crud/entity/requests"
	"crud/entity/responses"
	"crud/usecase"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetEmailUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	usecase usecase.Usecase
}

type ErrResponse struct {
	Message string
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

func (u *controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.Login
	json.NewDecoder(r.Body).Decode(&user)
	respon := u.usecase.Login(user)
	if respon.ID == 0 {
		resp := responses.Response{Status: http.StatusNotFound, Message: "Email atau password salah", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := responses.Response{Status: http.StatusOK, Message: "Data User Ditemukan", Result: map[string]interface{}{"data": respon}}
		json.NewEncoder(w).Encode(resp)
	}
}

func (u *controller) GetEmailUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *entity.Login
	json.NewDecoder(r.Body).Decode(&user)
	respon := u.usecase.FindUserByEmail(user)
	if respon.ID == 0 {
		resp := responses.Response{Status: http.StatusNotFound, Message: "Data User Tidak Ditemukan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := responses.Response{Status: http.StatusOK, Message: "Data User Ditemukan", Result: map[string]interface{}{"data": respon}}
		json.NewEncoder(w).Encode(resp)
	}
}
