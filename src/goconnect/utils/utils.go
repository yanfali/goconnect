package utils

import (
	"reflect"
	"strings"
)

func GetHandlerName(i interface{}) string {
	typeOf := reflect.TypeOf(i).String()
	return strings.Split(typeOf, ".")[1]
}
