package main

import (
	"fmt"
	"goconnect/core"
	"goconnect/middleware"
	"log"
	"net/http"
)

func init() {}

type MyApp struct{}

func (app *MyApp) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	log.Printf("%s: Serving App", app.Name())
	fmt.Fprintf(res, "Hello Mom")
	next()
}

func (app *MyApp) Name() string {
	return "app"
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

	auth, _ := middleware.NewBasicAuth("yan", "yan")
	connect.Use(auth)

	limit, _ := middleware.NewLimit(1)
	connect.Use(limit)

	static, _ := middleware.NewStatic("/public", "/tmp")
	connect.Use(static)
	connect.Use(&MyApp{})

	http.Handle("/", connect)
	http.ListenAndServe(":8000", nil)
}
