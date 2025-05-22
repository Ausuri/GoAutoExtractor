package antivirus

type AntiVirusInterface interface {
	ScanFile(path string) *AntiVirusScanResult
}

type AntiVirusScanResult struct {
	Error            error
	File             string
	VirusDescription string
	VirusFound       bool
}
