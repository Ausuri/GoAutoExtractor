package extensionsanitizer

type SanitizerInterface interface {
	RemoveExtension(fileName string) string
}
