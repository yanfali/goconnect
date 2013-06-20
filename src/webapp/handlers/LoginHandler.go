package handlers

import (
	"html/template"
	"log"
	"net/http"
	"webapp/config"
)

type LoginHandler struct {
	Prefix    string
	Templates *template.Template
}

func NewLoginHandler(prefix string) (*LoginHandler, error) {
	h := &LoginHandler{Prefix: prefix}
	appdir := config.GetAppDir()
	h.Templates = template.Must(template.ParseGlob(appdir + "/templates/common/*.tmpl"))
	template.Must(h.Templates.ParseGlob(appdir + "/templates/" + prefix + "/*.tmpl"))
	return h, nil
}

func (h *LoginHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("Running login handler")
	page := struct{ Title, Description, Username, Password, Prefix string }{
		"Login", "Enter credentials to login", "Username", "Password", "login",
	}
	err := h.Templates.ExecuteTemplate(res, "html5.tmpl", page)
	if err != nil {
		log.Print(err)
	}
}
