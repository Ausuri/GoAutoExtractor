package scanner

import "errors"

type MockScanner struct {
	IsErrorExpected bool
	IsVirusFound    bool
}

func (m *MockScanner) ScanFile(path string) *ScanResult {

	var err error

	if m.IsErrorExpected {
		err = errors.New("Mock scan error.")
	} else {
		err = nil
	}

	virusFound := m.IsVirusFound

	var virusDescription string
	if virusFound {
		virusDescription = "Mock virus found."
	} else {
		virusDescription = ""
	}

	return &ScanResult{
		Error:            err,
		File:             path,
		VirusDescription: virusDescription,
		VirusFound:       m.IsVirusFound,
	}

}
