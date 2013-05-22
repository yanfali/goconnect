package middleware

import (
	"goconnect/core"
	"log"
	"net/http"
	"strings"
)

type Static struct {
	Path    string
	MaxAge  int
	Root    string
	Handler http.Handler
}

func NewStatic(path string, root string) (*Static, error) {
	return &Static{Path: path, Root: root, Handler: http.FileServer(http.Dir(root))}, nil
}

func (static *Static) Name() string {
	return "static"
}

func (static *Static) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	log.Printf("%s", req.Method)
	if req.Method != "GET" && req.Method != "HEAD" {
		next()
		return
	}
	path := req.URL.Path
	if strings.HasPrefix(path, static.Path) {
		log.Printf("%s: path %s contains %s, serving files", static.Name(), path, static.Path)
		req.URL.Path = strings.TrimPrefix(path, static.Path)
		res.Header().Set("Cache-Control", "max-age=0")
		static.Handler.ServeHTTP(res, req)
		log.Printf("%s: should be done", static.Name())
	} else {
		next()
	}
}
