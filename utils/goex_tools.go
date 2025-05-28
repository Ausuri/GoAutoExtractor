package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-reflect"
)

// Parses file line by line to get list of items (1 item = 1 line)
func GetSupportedFileExtensions() ([]string, error) {

	var extensionFilePath string
	useDevConfigPath, parseError := strconv.ParseBool(os.Getenv("USE_DEV_CONFIG_PATHS"))
	if parseError != nil {
		return nil, parseError
	}

	// TODO: At some point maybe this needs to be refactored? There's probably a better way to check here.
	if useDevConfigPath {
		extensionFilePath = "./config/AllowedExtensions.txt"
	} else {
		//TODO put the production path here and remove exception.
		return nil, fmt.Errorf("path not implemented")
	}

	fileSlice, err := parseExtensionFile(extensionFilePath)
	if err != nil {
		return nil, err
	}

	return fileSlice, nil
}

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

func parseExtensionFile(filePath string) ([]string, error) {

	fileSlice := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileSlice = append(fileSlice, line)
	}

	return fileSlice, nil
}
