package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/renishb10/golang-jwt/db"
	"github.com/renishb10/golang-jwt/models"
)

var error models.Error

func CreateUser(user *models.User) (e *models.Error) {
	db := db.DB
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		fmt.Println(err)
		error.Message = err.Error()
		return &error
	}

	user.Password = ""
	return nil
}

func GetUser(user *models.User) (e *models.Error) {
	db := db.DB
	stmt := "select * from users where email = $1;"
	err := db.QueryRow(stmt, user.Email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "The user does not exist"
		} else {
			log.Fatal(err)
			error.Message = err.Error()
		}
		return &error
	}

	return nil
}
