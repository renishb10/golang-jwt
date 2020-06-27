package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/renishb10/golang-jwt/models"
	"github.com/renishb10/golang-jwt/repo"
	"github.com/renishb10/golang-jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var error models.Error

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error parsing signup request body", err)
	}

	defer r.Body.Close()

	if user.Email == "" {
		error.Message = "Email is missing"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)

	dbErr := repo.CreateUser(&user)
	fmt.Printf("%v", dbErr)
	if dbErr != nil {
		utils.RespondWithError(w, http.StatusBadRequest, *dbErr)
		return
	}

	utils.ResponseJSON(&w, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func Protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected"))
}
