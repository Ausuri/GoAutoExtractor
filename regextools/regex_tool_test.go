package regextools

import (
	configmanager "GoAutoExtractor/config-manager"
	"fmt"
	"testing"
)

const TEST_FILE = "./testarchive"

func TestRemoveExtension(t *testing.T) {

	configmanager.InitializeTestConfig(nil)
	acceptedExtensions := configmanager.GetAllowedExtensions()
	regexTool := RegexTool{}

	for _, slice := range acceptedExtensions {

		// Create file name with extension from list.
		extension := slice
		file := TEST_FILE + extension

		// Test removing extension
		baseFileName := regexTool.RemoveExtension(file)
		if baseFileName != TEST_FILE {
			fmt.Printf("failure removing extension, expected %v but received %v", TEST_FILE, baseFileName)
			t.FailNow()
		}

		fmt.Printf("Success removing extension %v", extension)
	}

}

func TestVerifyValidArchive(t *testing.T) {

	configmanager.InitializeTestConfig(nil)
	acceptedExtensions := configmanager.GetAllowedExtensions()
	regexTool := RegexTool{}

	// Test for file extensions in accepted list.
	for _, slice := range acceptedExtensions {

		// Create file name with extension from list.
		extension := slice
		file := TEST_FILE + extension

		// Test removing extension
		isValidArchive := regexTool.VerifyValidArchive(file)
		if !isValidArchive {
			fmt.Printf("failure verifying archive, filename %v should be recognized as valid\n", file)
			t.Fail()
		}

		fmt.Printf("Success verifying archive %v\n", file)

		// Test for all variations of tar.{extension} where extension is an item from the accepted list.
		if extension == ".tar" {
			continue
		}

		tarFile := TEST_FILE + ".tar" + extension
		if !regexTool.VerifyValidArchive(tarFile) {
			fmt.Printf("failure verifying archive, filename %v should be recognized as valid\n", file)
			t.Fail()
		}

		fmt.Printf("Success verifying tar archive %v\n", tarFile)

	}
}
