package middleware

import (
	"encoding/hex"
	"github.com/gorilla/securecookie"
	"goconnect/core"
	"net/http"
)

const ()

type UserSession struct {
	CookieName string
	Secure bool
}

func NewUserSession(cookieName string, secure bool) (*UserSession, error) {
	return &UserSession{CookieName: cookieName + "_session", Secure: secure}, nil
}

func (user *UserSession) Name() string {
	return "user-session"
}

func (user *UserSession) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	if _, err := req.Cookie(user.CookieName); err == http.ErrNoCookie {
		rawid := securecookie.GenerateRandomKey(16)
		cookie := &http.Cookie{
			Name:     user.CookieName,
			Value:    hex.EncodeToString(rawid),
			Path:     "/",
			HttpOnly: true,
			Secure: user.Secure,
		}
		http.SetCookie(res, cookie)
	}
	next()
}
