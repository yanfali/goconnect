package main

import (
	"fmt"
	"goconnect/core"
	"goconnect/middleware"
	"log"
	"net/http"
	"os"
	"webapp/config"
	"webapp/handlers"
)

var (
	baseDir = ""
	appDir  = ""
)

func init() {
	err := config.SetBaseDir(os.Args[0])
	if err != nil {
		panic(err)
	}
	log.Printf("basedir: %s, appdir: %s\n", config.GetBaseDir(), config.GetAppDir())
}

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
 	loginHandler, err := handlers.NewLoginHandler("login")
	if err != nil {
		panic(err)
	}
	app.Router.Handle("/login", loginHandler)
	connect.Use(app)

	http.Handle("/", connect)
	http.ListenAndServe(":8000", nil)
}
