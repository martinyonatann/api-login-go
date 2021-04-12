package main

import (
	"api-login-go/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/registrasi", controller.RegistrasiHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
