package main

import (
	"log"
	"net/http"

	"crud/auth"
	"crud/controller"
	entity "crud/entity/requests"
	"crud/repository"
	"crud/usecase"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

const port string = ":9001"

func initializeRouter() {
	r := mux.NewRouter()

	DNS := auth.DNS()
	db, err := gorm.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to DB")
	}
	db.AutoMigrate(&entity.User{})

	userRepository := repository.NewRepository(db)

	userUsecase := usecase.NewUsecase(userRepository)
	userAuth := auth.NewService()
	userController := controller.NewController(userUsecase, userAuth)

	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", userController.GetUser).Methods("GET")
	r.HandleFunc("/userx/{id}", userController.GetUserx).Methods("GET")
	r.HandleFunc("/user", userController.CreateUser).Methods("POST")
	r.HandleFunc("/userEmail", userController.GetEmailUser).Methods("POST")
	r.HandleFunc("/userlogin", userController.LoginUser).Methods("POST")
	r.HandleFunc("/user/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", userController.DeleteUser).Methods("DELETE")

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	//repository.InitialMigration()
	initializeRouter()

}
