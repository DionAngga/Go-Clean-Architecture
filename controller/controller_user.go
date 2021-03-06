package controller

import (
	"crud/auth"
	entity "crud/entity/requests"
	"crud/entity/responses"
	"crud/usecase"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUserx(w http.ResponseWriter, r *http.Request)
	GetEmailUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	usecase usecase.Usecase
	auth    auth.Service
}

type ErrResponse struct {
	Message string
}

func NewController(usecase usecase.Usecase, auth auth.Service) Controller {
	return &controller{usecase, auth}
}

func (u *controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []entity.User
	Users, err := u.usecase.FindAllUser(&users)
	if err != nil {
		respon := responses.Response{Status: http.StatusNotFound, Message: "Data user tidak ditemukan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	}
	var newuser []responses.UserRespon
	for _, user := range *Users {
		resps := responses.UserRespon{
			Model:   user.Model,
			Nasabah: user.Nasabah,
			Age:     user.Age,
			Name:    user.Name,
			Email:   user.Email,
		}
		newuser = append(newuser, resps)
	}
	respon := responses.Response{Status: http.StatusOK, Message: "Data user ditemukan", Result: map[string]interface{}{"data": newuser}}
	json.NewEncoder(w).Encode(respon)
}

func (u *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	User, err := u.usecase.FindUser(params["id"])
	if err != nil {
		fmt.Println("+++++++++++++ error ++++++++++++++")
		respon := responses.Response{Status: http.StatusNotFound, Message: "Data user tidak ditemukan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	} else {
		var newuser = &responses.UserRespon{}
		newuser.Model = User.Model
		newuser.Nasabah = User.Nasabah
		newuser.Age = User.Age
		newuser.Name = User.Name
		newuser.Email = User.Email
		respon := responses.Response{Status: http.StatusFound, Message: "Data user ditemukan", Result: map[string]interface{}{"data": newuser}}
		fmt.Println("+++++++++++++ success ++++++++++++++")
		json.NewEncoder(w).Encode(respon)
	}
}

func (u *controller) GetUserx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	User, err := u.usecase.FindUserx(params["id"])
	if err != nil {
		fmt.Println("+++++++++++++ error ++++++++++++++")
		respon := responses.Response{Status: http.StatusNotFound, Message: "Data user tidak ditemukan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	} else {
		var newuser = &responses.UserRespon{}
		newuser.Nasabah = User.Nasabah
		newuser.Age = User.Age
		newuser.Name = User.Name
		newuser.Email = User.Email
		respon := responses.Response{Status: http.StatusFound, Message: "Data user ditemukan", Result: map[string]interface{}{"data": newuser}}
		fmt.Println("+++++++++++++ success ++++++++++++++")
		json.NewEncoder(w).Encode(respon)
	}
}

func (u *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	User, err := u.usecase.CreateUser(&user)
	if err != nil {
		respon := responses.Response{Status: http.StatusBadRequest, Message: "Terjadi kesalahan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	} else {
		var newuser responses.UserRespon
		newuser.Model = User.Model
		newuser.Nasabah = User.Nasabah
		newuser.Age = User.Age
		newuser.Name = User.Name
		newuser.Email = User.Email
		respon := responses.Response{Status: http.StatusFound, Message: "akun berhasil dibuat", Result: map[string]interface{}{"data": newuser}}
		json.NewEncoder(w).Encode(respon)
	}
}

func (u *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user, _ := u.usecase.FindUser(params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	User, err := u.usecase.UpdateUser(user)
	if err != nil {
		respon := responses.Response{Status: http.StatusBadRequest, Message: "Terjadi kesalahan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	} else {
		var newuser responses.UserRespon
		newuser.Model = User.Model
		newuser.Nasabah = User.Nasabah
		newuser.Age = User.Age
		newuser.Name = User.Name
		newuser.Email = User.Email
		respon := responses.Response{Status: http.StatusFound, Message: "akun berhasil diperbaharui", Result: map[string]interface{}{"data": newuser}}
		json.NewEncoder(w).Encode(respon)
	}
}

func (u *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	User, error := u.usecase.FindUser(params["id"])
	var newuser responses.UserRespon
	newuser.Model = User.Model
	newuser.Nasabah = User.Nasabah
	newuser.Age = User.Age
	newuser.Name = User.Name
	newuser.Email = User.Email
	if error != nil {
		respon := responses.Response{Status: http.StatusNotFound, Message: "Data user tidak ditemukan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	}
	_, err := u.usecase.DeleteUser(&user, params["id"])
	if err != nil {
		respon := responses.Response{Status: http.StatusBadRequest, Message: "Terjadi kesalahan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	} else {
		respon := responses.Response{Status: http.StatusOK, Message: "Data user telah di Hapus", Result: map[string]interface{}{"data": newuser}}
		json.NewEncoder(w).Encode(respon)
	}
}

func (u *controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.Login
	json.NewDecoder(r.Body).Decode(&user)
	respon, err := u.usecase.Login(user)
	if err != nil {
		respon := responses.Response{Status: http.StatusBadRequest, Message: "Terjadi kesalahan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	}
	if respon.ID == 0 {
		resp := responses.Response{Status: http.StatusNotFound, Message: "Email atau password salah", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(resp)
	} else {
		tokenUser, err := u.auth.GenerateToken(respon.ID, respon.Email)
		if err != nil {
			resp := responses.Response{Status: http.StatusBadRequest, Message: "Token tidak muncul", Result: map[string]interface{}{"data": nil}}
			json.NewEncoder(w).Encode(resp)
		}
		respon.Token = tokenUser
		resp := responses.Response{Status: http.StatusOK, Message: "Data User Ditemukan", Result: map[string]interface{}{"data": respon}}
		json.NewEncoder(w).Encode(resp)
	}
}

func (u *controller) GetEmailUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *entity.Login
	json.NewDecoder(r.Body).Decode(&user)
	respon, err := u.usecase.FindUserByEmail(user)
	if err != nil {
		respon := responses.Response{Status: http.StatusBadRequest, Message: "Terjadi kesalahan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(respon)
	}
	if respon.ID == 0 {
		resp := responses.Response{Status: http.StatusNotFound, Message: "Data User Tidak Ditemukan", Result: map[string]interface{}{"data": nil}}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := responses.Response{Status: http.StatusOK, Message: "Data User Ditemukan", Result: map[string]interface{}{"data": respon}}
		json.NewEncoder(w).Encode(resp)
	}
}
