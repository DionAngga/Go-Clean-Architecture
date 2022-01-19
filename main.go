package main

import (
	"log"
	"net/http"

	"crud/controller"
	"crud/repository"

	"github.com/gorilla/mux"
)

const port string = ":9001"

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controller.UpdateUsers).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	repository.InitialMigration()
	initializeRouter()

}
