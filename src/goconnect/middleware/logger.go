package middleware

import (
	"fmt"
	"goconnect/core"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	Format string
	Stream *os.File
}

func (logger *Logger) Name() string {
	return "logger"
}

func NewLogger() (*Logger, error) {
	return &Logger{Stream: os.Stdout}, nil
}

func getReferrer(req *http.Request) string {
	referrer, ok := req.Header["Referer"]
	if ok {
		return referrer[0]
	}
	return ""
}

func getUserAgent(req *http.Request) string {
	useragent, ok := req.Header["User-Agent"]
	if ok {
		return useragent[0]
	}
	return ""
}

func getStatus(res http.ResponseWriter) string {
	return res.Header().Get("Status")
}

func getContentLength(res http.ResponseWriter) string {
	return res.Header().Get("Content-Length")
}

func (logger *Logger) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	fmt.Fprintf(logger.Stream, "%s - - [%s] \"%s %s HTTP/:%d.%d\" %s %s \"%s\" \"%s\"\n", req.RemoteAddr, time.Now(), req.Method, req.URL, req.ProtoMajor, req.ProtoMinor, getStatus(res), getContentLength(res), getReferrer(req), getUserAgent(req))
	next()
}
