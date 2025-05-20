package compressionmanager

import (
	"MediaCompressionManager/compression"
	"MediaCompressionManager/extensionsanitizer"
	"MediaCompressionManager/scanner"
)

type CompressionManager struct {
	extractor compression.DecompressorInterface
	sanitizer extensionsanitizer.SanitizerInterface
	scanner   scanner.ScannerInterface
}
