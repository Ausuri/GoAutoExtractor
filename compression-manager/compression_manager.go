package compressionmanager

import (
	"MediaCompressionManager/compression"
	"MediaCompressionManager/extensionsanitizer"
	"MediaCompressionManager/scanner"
	"MediaCompressionManager/statuschecker"
)

type CompressionManager struct {
	Extractor     compression.DecompressorInterface
	Sanitizer     extensionsanitizer.SanitizerInterface
	Scanner       scanner.ScannerInterface
	Statuschecker statuschecker.StatusCheckerInterface
}
