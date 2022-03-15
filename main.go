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

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

const port string = ":9001"

func initializeRouter() {
	r := mux.NewRouter()

	DNS := "root:@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to DB")
	}
	db.AutoMigrate(&entity.User{})

	userRepository := repository.NewRepository(db)

	userUsecase := usecase.NewUsecase(userRepository)
	userAuth := auth.NewService()
	// userHandler := handler.NewUserHandler(userUseCase)
	userController := controller.NewController(userUsecase, userAuth)

	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", userController.GetUser).Methods("GET")
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
