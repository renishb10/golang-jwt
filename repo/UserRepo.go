package repo

import (
	"fmt"

	"github.com/renishb10/golang-jwt/db"
	"github.com/renishb10/golang-jwt/models"
)

var error models.Error

func CreateUser(user *models.User) (e *models.Error) {
	db := db.DB
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"

	fmt.Println(stmt, user.Email, user.Password)
	fmt.Println(db)
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		fmt.Println(err)
		error.Message = err.Error()
		return &error
	}

	user.Password = ""
	return nil
}
