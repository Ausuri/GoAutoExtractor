package regextools

type MockRegexTool struct {
	IsError bool
}

func (m *MockRegexTool) RemoveExtension(fileName string) string {

	if m.IsError {
		return ""
	}

	return fileName
}
