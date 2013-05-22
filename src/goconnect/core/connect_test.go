package core

import (
	"net/http"
	"testing"
)

func Test_Length(t *testing.T) {
	connect, err := NewConnect()
	if err != nil {
		t.Error(err)
	} else if connect.Length() == 0 {
		t.Log("ok")
	} else {
		t.Error("Connect was not 0 in length")
	}
}

func MyTestHandler(res http.ResponseWriter, req *http.Request) {

}

func Test_Use(t *testing.T) {
	connect, _ := NewConnect()
	connect.Use(MyTestHandler)
	if connect.Length() == 1 {
		t.Log("ok")
	} else {
		t.Error("Connect should be 1 in Length")
	}
}
