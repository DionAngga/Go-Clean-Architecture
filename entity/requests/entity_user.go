package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Nasabah  string `json:"nasabah"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Nasabah  string `json:"nasabah"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Userx struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Nasabah  string `json:"nasabah"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
