package regextools

import (
	"GoAutoExtractor/utils"
	"fmt"
	"regexp"
	"slices"
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

	fileExtensionSlice, err := utils.GetSupportedFileExtensions()
	if err != nil {
		panic(fmt.Sprintf("valid archive retrievel of file extension list error: %v", err))
	}

	//TODO: Implement a regex string that also uses the AllowedExtensions.txt to verify extensions.
	tarballRegex := regexp.MustCompile(`\.tar\.\S{2,3}$`)
	fileExtensionRegex := regexp.MustCompile(`\.\S{2,4}$`)

	tarballExtension := tarballRegex.FindString(fileName)
	archiveExtension := fileExtensionRegex.FindString(fileName)

	if tarballExtension != "" {
		isValid = slices.Contains(fileExtensionSlice, tarballExtension)
	} else if archiveExtension != "" {
		isValid = slices.Contains(fileExtensionSlice, archiveExtension)
	}

	return isValid
}
