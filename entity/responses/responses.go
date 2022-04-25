package responses

import "github.com/jinzhu/gorm"

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Result  map[string]interface{} `json:"result"`
}

type UserLogin struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Nasabah string `json:"nasabah"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

type UserRespon struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Nasabah string `json:"nasabah"`
	Email   string `json:"email"`
}
