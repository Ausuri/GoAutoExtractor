package utils

import (
	"fmt"
	"time"

	"github.com/goccy/go-reflect"
)

// Uses reflection to return a map of key/value pairs for all the properties of a struct or pointer.
func GetObjectMap(obj interface{}) map[string]any {

	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	// If it's a pointer, dereference
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return nil
	}

	result := make(map[string]any)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		result[fieldType.Name] = field.Interface()
	}

	return result
}

func PauseMilliseconds(msPauseTime int64) {
	time.Sleep(time.Duration(msPauseTime) * time.Millisecond)
}

func PauseSeconds(secondsPauseTime int64) {
	time.Sleep(time.Duration(secondsPauseTime) * time.Second)
}
