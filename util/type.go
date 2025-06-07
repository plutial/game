package util

import (
	"reflect"
)

// Get the type of a generic type interface as reflect.Type variable
func GetType[T any]() string {
	var temp T
	return reflect.TypeOf(temp).Name()
}
