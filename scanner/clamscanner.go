package scanner

import (
	"os"

	"github.com/sheenobu/go-clamscan"
)

type ClamScanner struct{}

func (c ClamScanner) ScanFile(path string) (result *ScanResult) {

	options := &clamscan.Options{
		BinaryLocation: os.Getenv("CLAMSCAN_BINARY"),
	}

	sr, cErr := clamscan.Scan(options, path)
	scanResult := <-sr

	result = &ScanResult{
		Error:            cErr,
		File:             scanResult.File,
		VirusDescription: scanResult.Virus,
		VirusFound:       scanResult.Found,
	}

	return result
}
