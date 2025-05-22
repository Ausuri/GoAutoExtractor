package antivirus

import (
	"os"

	"github.com/sheenobu/go-clamscan"
)

type ClamAntiVirus struct{}

func (c ClamAntiVirus) ScanFile(path string) (result *AntiVirusScanResult) {

	options := &clamscan.Options{
		BinaryLocation: os.Getenv("CLAMSCAN_BINARY"),
	}

	sr, cErr := clamscan.Scan(options, path)
	scanResult := <-sr

	result = &AntiVirusScanResult{
		Error:            cErr,
		File:             scanResult.File,
		VirusDescription: scanResult.Virus,
		VirusFound:       scanResult.Found,
	}

	return result
}
