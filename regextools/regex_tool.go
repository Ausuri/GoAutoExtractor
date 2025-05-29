package regextools

import (
	configmanager "GoAutoExtractor/config-manager"
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

	fileExtensionSlice := configmanager.GetAllowedExtensions()

	//TODO: Implement a regex string that also uses the AllowedExtensions.txt to verify extensions.
	tarballRegex := regexp.MustCompile(`\.tar\.\S{2,3}$`)
	fileExtensionRegex := regexp.MustCompile(`\.\S{2,4}$`)

	tarballExtension := tarballRegex.FindString(fileName)
	archiveExtension := fileExtensionRegex.FindString(fileName)

	if tarballExtension != "" {
		extensionSplitRegex := regexp.MustCompile(`\.\w{2,4}`)
		extSlice := extensionSplitRegex.FindAllString(tarballExtension, -1)
		isValid = (extSlice[0] == ".tar") && slices.Contains(fileExtensionSlice, extSlice[1])
	} else if archiveExtension != "" {
		isValid = slices.Contains(fileExtensionSlice, archiveExtension)
	}

	return isValid
}
