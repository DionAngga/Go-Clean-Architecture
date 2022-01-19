package main

import (
	"log"
	"net/http"

	"crud/repository"
	"crud/usecase"

	"github.com/gorilla/mux"
)

const port string = ":9001"

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", usecase.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", usecase.GetUser).Methods("GET")
	r.HandleFunc("/user", usecase.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", usecase.UpdateUsers).Methods("PUT")
	r.HandleFunc("/user/{id}", usecase.DeleteUser).Methods("DELETE")

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	repository.InitialMigration()
	initializeRouter()

}
