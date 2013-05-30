package core

import (
	"github.com/gorilla/mux"
	"net/http"
)

var ()

const ()

type Application struct {
	Router *mux.Router
}

func NewApplication() *Application {
	return &Application{Router: mux.NewRouter()}
}

func (app *Application) Name() string {
	return "application"
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request, next NextMiddleware) {
	app.Router.ServeHTTP(res, req)
	next()
}
