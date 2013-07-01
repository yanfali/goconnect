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

	staticScripts, _ := middleware.NewStatic("/scripts", config.GetAppDir() + "/scripts")
	connect.Use(staticScripts)
	staticImages, _ := middleware.NewStatic("/images", config.GetAppDir() + "/images")
	connect.Use(staticImages)
	staticStyles, _ := middleware.NewStatic("/styles", config.GetAppDir() + "/styles")
	connect.Use(staticStyles)
	app := core.NewApplication()
	connect.Use(app)


	/* Middleware that wants to custom SubRouting */
 	loginHandler, err := handlers.NewLoginHandler(app.Router.PathPrefix("/login").Subrouter())
	if err != nil {
		panic(err)
	}

	// Set a customer auth required validator
	auth.ValidatorFn = loginHandler.ValidAuthCookie

	http.Handle("/", connect)

	/* Application Goes Here */
	app.Router.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8000", nil)
}
