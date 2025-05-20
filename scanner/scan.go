package scanner

type ScannerInterface interface {
	ScanFile(path string) *ScanResult
}

type ScanResult struct {
	Error            error
	File             string
	VirusDescription string
	VirusFound       bool
}
