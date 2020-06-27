package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/renishb10/golang-jwt/models"
	"github.com/renishb10/golang-jwt/repo"
	"github.com/renishb10/golang-jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error parsing signup request body", err)
	}

	defer r.Body.Close()

	validErr := user.Validate()
	if validErr != nil {
		utils.RespondWithError(w, http.StatusBadRequest, *validErr)
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

func generateToken(user models.User) (string, error) {
	var err error
	secret := os.Getenv("MYSECRETKEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		user  models.User
		error models.Error
		jwt   models.JWT
	)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		error.Message = "Error while parsing input"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	defer r.Body.Close()

	validErr := user.Validate()
	if validErr != nil {
		utils.RespondWithError(w, http.StatusBadRequest, *validErr)
		return
	}

	rawPassword := user.Password

	if usrErr := repo.GetUser(&user); usrErr != nil {
		utils.RespondWithError(w, http.StatusNotFound, error)
		return
	}

	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))

	if err != nil {
		error.Message = "Invalid Password"
		utils.RespondWithError(w, http.StatusUnauthorized, error)
		return
	}

	token, err := generateToken(user)
	if err != nil {
		error.Message = "Error generating token"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	utils.ResponseJSON(&w, jwt)
}

func Protected(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Invoked")
}
