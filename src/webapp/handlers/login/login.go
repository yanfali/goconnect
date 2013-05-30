package handlers

import (
	"net/http"
	"html/template"
)

type LoginHandler struct {
	Prefix string
	Template *template.Template
}

func NewLoginHandler(prefix string) *LoginHandler {
	return &LoginHandler{Prefix: prefix}
}

func (handler *LoginHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
}
