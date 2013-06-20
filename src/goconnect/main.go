package main

import (
	"fmt"
	"goconnect/core"
	"goconnect/middleware"
	"log"
	"net/http"
)

func init() {}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome Home")
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Login Page Here")
}

func main() {
	connect, err := core.NewConnect()
	if err != nil {
		log.Panic(err)
	}

	logger, _ := middleware.NewLogger()
	connect.Use(logger)

	userSess, _ := middleware.NewUserSession("yanapp", false)
	connect.Use(userSess)

	publicUrls := []string{}
	auth, _ := middleware.NewRequireAuth(publicUrls, "/login")
	connect.Use(auth)

	limit, _ := middleware.NewLimit(1)
	connect.Use(limit)

	static, _ := middleware.NewStatic("/public", "/tmp")
	connect.Use(static)
	app := core.NewApplication()
	app.Router.HandleFunc("/", HomeHandler)
	app.Router.HandleFunc("/login", LoginHandler)
	connect.Use(app)

	http.Handle("/", connect)
	http.ListenAndServe(":8000", nil)
}
