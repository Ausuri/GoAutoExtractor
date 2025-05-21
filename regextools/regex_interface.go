package regextools

type RegexToolInterface interface {
	RemoveExtension(fileName string) string  // Sanitizes the filename by removing the extension.
	VerifyValidArchive(fileName string) bool // Check if the file is a supported archive type and not a multipart archive.
}
