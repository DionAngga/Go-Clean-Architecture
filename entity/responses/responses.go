package responses

import "gorm.io/gorm"

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Result  map[string]interface{} `json:"result"`
}

type UserRespon struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Nasabah string `json:"nasabah"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}
