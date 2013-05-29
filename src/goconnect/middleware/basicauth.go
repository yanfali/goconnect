package middleware

import (
	"encoding/base64"
	"goconnect/core"
	"log"
	"net/http"
	"strings"
)

var (
	BASIC     = "Basic "
	BASIC_LEN = len(BASIC)
)

type BasicAuth struct {
	User     string
	Password string
}

func NewBasicAuth(user string, password string) (*BasicAuth, error) {
	return &BasicAuth{User: user, Password: password}, nil
}

func (basic *BasicAuth) Name() string {
	return "basicauth"
}

func (basic *BasicAuth) Unauthorized(res http.ResponseWriter) {
	res.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
	http.Error(res, "Not Authorized", http.StatusUnauthorized)
}

func (basic *BasicAuth) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	auth, ok := req.Header["Authorization"]
	authenticated := false
	defer func() {
		if authenticated {
			next()
		} else {
			basic.Unauthorized(res)
		}
	}()
	if !ok {
		return
	}
	log.Printf("%s: auth = %s", basic.Name(), auth)
	authtoken := auth[0]
	if strings.HasPrefix(authtoken, BASIC) {
		data := []byte(basic.User + ":" + basic.Password)
		str := base64.StdEncoding.EncodeToString(data)
		if str != authtoken[BASIC_LEN:] {
			return
		}
	} else {
		return
	}
	authenticated = true
}
