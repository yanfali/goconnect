package middleware

import (
	"goconnect/core"
	"log"
	"net/http"
	"strconv"
)

type Limit struct {
	LimitInBytes int
}

func NewLimit(limitInKiB int) (*Limit, error) {
	return &Limit{LimitInBytes: limitInKiB * 1024}, nil
}

func (limit *Limit) Name() string {
	return "limit"
}

type resState struct {
	Fail bool
	Msg string	
	Code int
}

func (limit *Limit) ServeHTTP(res http.ResponseWriter, req *http.Request, next core.NextMiddleware) {
	//log.Println(req.Header);
	length, ok := req.Header["Content-Length"]
	if !ok {
		next()
		return
	}
	state := resState{}
	defer func() {
		if state.Fail {
			log.Printf("%s: content-length: %s exceeds limit %d", limit.Name(), length, limit.LimitInBytes)
			http.Error(res, state.Msg, state.Code)
		} else {
			next()
		}
	}()
	contentlen, err := strconv.Atoi(length[0])
	if err != nil {
		log.Println(err)
		state = resState{true, "Internal Server Error", http.StatusInternalServerError}
		return
	}
	if contentlen > limit.LimitInBytes {
		state = resState{true, "Request Entity Too Large", http.StatusRequestEntityTooLarge}
		return
	}
}
