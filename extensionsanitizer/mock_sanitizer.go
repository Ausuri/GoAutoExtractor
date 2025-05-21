package extensionsanitizer

type MockSanitizer struct {
	IsError bool
}

func (m *MockSanitizer) RemoveExtension(fileName string) string {

	if m.IsError {
		return ""
	}

	return fileName
}
