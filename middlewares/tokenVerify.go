package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/renishb10/golang-jwt/models"
	"github.com/renishb10/golang-jwt/utils"
)

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	var cErr models.Error

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error parsing token")
				}

				return []byte(os.Getenv("MYSECRETKEY")), nil
			})

			if err != nil {
				cErr.Message = err.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, cErr)
				return
			}

			spew.Dump(token)

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				cErr.Message = "Invalid Token"
				utils.RespondWithError(w, http.StatusUnauthorized, cErr)
				return
			}
		} else {
			cErr.Message = "Token Format is invalid"
			utils.RespondWithError(w, http.StatusBadRequest, cErr)
			return
		}
	})
}
