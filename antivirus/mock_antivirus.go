package antivirus

import "errors"

type MockAntiVirus struct {
	IsErrorExpected bool
	IsVirusFound    bool
}

func (m *MockAntiVirus) ScanFile(path string) *AntiVirusScanResult {

	var err error

	if m.IsErrorExpected {
		err = errors.New("mock scan error")
	} else {
		err = nil
	}

	virusFound := m.IsVirusFound

	var virusDescription string
	if virusFound {
		virusDescription = "mock virus found"
	} else {
		virusDescription = ""
	}

	return &AntiVirusScanResult{
		Error:            err,
		File:             path,
		VirusDescription: virusDescription,
		VirusFound:       m.IsVirusFound,
	}

}
