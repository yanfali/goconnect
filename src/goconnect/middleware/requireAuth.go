package middleware

import (
	"goconnect/core"
	"net/http"
	"strings"
	"log"
)

const ()

var ()

/*
 * If a session starts without the auth cookie, redirect to loginUrl
 */
type RequireAuth struct {
	/* Urls that don't require authentication */
	PublicUrls []string
	LoginUrl   string
}

func NewRequireAuth(publicUrls []string, loginUrl string) (*RequireAuth, error) {
	publicUrls = append(publicUrls, loginUrl)
	return &RequireAuth{PublicUrls: publicUrls, LoginUrl: loginUrl}, nil
}

func (auth *RequireAuth) Name() string {
	return "require-auth"
}

func invalid(cookie *http.Cookie) bool {
	return true
}

func (auth *RequireAuth) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	// if url is in public url list, call next
	// login url automatically gets added to public list
	// url is not in public url list redirect to login url
	for _, url := range auth.PublicUrls {
		log.Printf("url: %s, matching against: %s\n", url, auth.LoginUrl)
		if strings.HasPrefix(req.URL.Path, url) {
			log.Println("matched, calling next handler")
			next();
			return
		}
	}
	authCookie, err := req.Cookie("auth")
	if (err == http.ErrNoCookie || invalid(authCookie)) {
		http.Redirect(res, req, "/login", 302)
		return
	}
	next();
}
