package compressionmanager

import (
	"MediaCompressionManager/compression"
	"MediaCompressionManager/extensionsanitizer"
	"MediaCompressionManager/scanner"
	"MediaCompressionManager/statuschecker"
)

type Builder struct {
	extractor     compression.DecompressorInterface
	sanitizer     extensionsanitizer.SanitizerInterface
	scanner       scanner.ScannerInterface
	statuschecker statuschecker.StatusCheckerInterface
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetDecompressor(iextractor compression.DecompressorInterface) {
	b.extractor = iextractor
}

func (b *Builder) SetExtensionSanitizer(iextension extensionsanitizer.SanitizerInterface) {
	b.sanitizer = iextension
}

func (b *Builder) SetScanner(iscanner scanner.ScannerInterface) {
	b.scanner = iscanner
}

func (b *Builder) SetStatusChecker(istatuschecker statuschecker.StatusCheckerInterface) {
	b.statuschecker = istatuschecker
}

func (b *Builder) Build() *CompressionManager {

	//Initialize interfaces to default implementations if not set.
	if b.extractor == nil {
		b.extractor = &compression.HashigoExtractor{}
	}
	if b.sanitizer == nil {
		b.sanitizer = &extensionsanitizer.RegexSanitizer{}
	}
	if b.scanner == nil {
		b.scanner = &scanner.ClamScanner{}
	}
	if b.statuschecker == nil {
		b.statuschecker = &statuschecker.SyncthingStatusChecker{}
	}

	return &CompressionManager{
		Extractor:     b.extractor,
		Sanitizer:     b.sanitizer,
		Scanner:       b.scanner,
		Statuschecker: b.statuschecker,
	}
}
