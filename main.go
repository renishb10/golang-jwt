package main

import (
	App "github.com/renishb10/golang-jwt/app"
)

func main() {
	a := App.App{}
	a.Init()

	a.Run()
}
