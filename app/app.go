package App

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renishb10/golang-jwt/db"
	"github.com/renishb10/golang-jwt/handlers"
	"github.com/renishb10/golang-jwt/middlewares"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	db.Init()

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	a.Router.HandleFunc("/login", handlers.Login).Methods("POST")
	a.Router.HandleFunc("/protected", middlewares.TokenVerifyMiddleware(handlers.Protected)).Methods("GET")
}

func (a *App) Run() {
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
