package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renishb10/golang-jwt/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	//r.HandleFunc("/protected", middlewares.TokenVerifyMiddleware(handlers.Protected)).Methods("GET")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
