package regextools

import (
	"regexp"
)

type RegexTool struct{}

func (r RegexTool) RemoveExtension(fileName string) string {

	//TODO: Implement a regex string that also uses the AllowedExtensions.txt to verify extensions.
	tarballRegex := regexp.MustCompile(`\.tar\.\S{2,3}$`)
	fileExtensionRegex := regexp.MustCompile(`\.\S{2,4}$`)

	var result string = tarballRegex.ReplaceAllString(fileName, "")
	result = fileExtensionRegex.ReplaceAllString(result, "")

	return result
}

func (r RegexTool) VerifyValidArchive(fileName string) (isValid bool) {

	//TODO: Implement a regex string that also uses the AllowedExtensions.txt to verify extensions.
	zipRegex := regexp.MustCompile(`\.zip$`)
	tarballRegex := regexp.MustCompile(`\.tar\.\S{2,3}$`)
	fileExtensionRegex := regexp.MustCompile(`\.\S{2,4}$`)

	if zipRegex.MatchString(fileName) || tarballRegex.MatchString(fileName) || fileExtensionRegex.MatchString(fileName) {
		return true
	}

	return false
}
