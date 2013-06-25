package handlers

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"goconnect/utils"
	"html/template"
	"log"
	"net/http"
	"webapp/config"
	"time"
)

type LoginHandler struct {
	Router       *mux.Router
	Prefix       string
	Templates    *template.Template
	SecureCookie *securecookie.SecureCookie
}

func NewLoginHandler(router *mux.Router) (*LoginHandler, error) {
	h := &LoginHandler{Router: router}
	h.Prefix = utils.GetHandlerName(h)
	appdir := config.GetAppDir()
	h.Templates = template.Must(template.ParseGlob(appdir + "/templates/common/*.tmpl"))
	template.Must(h.Templates.ParseGlob(appdir + "/templates/" + h.Prefix + "/*.tmpl"))

	h.Router.Methods("GET").HandlerFunc(h.GetHandler)
	h.Router.Methods("POST").HandlerFunc(h.PostHandler)

	h.SecureCookie = securecookie.New(config.GetHashKey(), config.GetBlockKey())
	return h, nil
}

func (h *LoginHandler) GetHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Running login handler")

	_, err := req.Cookie("auth")
	if err == nil {
		http.Redirect(res, req, "/", 302)
		return
	}


	page := struct{ Title, Description, Username, Password, Prefix string }{
		"Login", "Enter credentials to login", "Username", "Password", "login",
	}
	err = h.Templates.ExecuteTemplate(res, "html5.tmpl", page)
	if err != nil {
		log.Print(err)
	}
}

type AuthCookie struct {
	Username string
	Timestamp int64
}

func (h *LoginHandler) ValidAuthCookie(authCookie *http.Cookie) (error) {
	var err error
	var value AuthCookie
	if err = h.SecureCookie.Decode("auth", authCookie.Value, &value); err == nil {
		log.Printf("The auth cookie was valid and contained: u=%s t=%d", value.Username, value.Timestamp)
	}
	return err;
}

func (h *LoginHandler) PostHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	username, password := req.PostForm["username"][0], req.PostForm["password"][0]
	log.Printf("Username: %s, Password: %s", username, password)
	if username == "abc" && password == "123" {
		val := AuthCookie{
			username,
			time.Now().Unix(),
		}
		if encoded, err := h.SecureCookie.Encode("auth", val); err == nil {
			cookie := &http.Cookie{
				Name: "auth",
				Value: encoded,
				Path: "/",
			}
			http.SetCookie(res, cookie)
			log.Printf("Auth Successful. Redirecting to /");
			http.Redirect(res, req, "/", 302)
		}
	}
	// TODO - return errors to page

}
