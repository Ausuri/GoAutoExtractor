package extensionsanitizer

import (
	"regexp"
)

type RegexSanitizer struct{}

func (r RegexSanitizer) RemoveExtension(fileName string) string {
	tarballRegex := regexp.MustCompile(`\.tar\.\S{2,3}$`)
	fileExtensionRegex := regexp.MustCompile(`\.\S{2,4}$`)

	var result string = tarballRegex.ReplaceAllString(fileName, "")
	result = fileExtensionRegex.ReplaceAllString(result, "")

	return result
}
