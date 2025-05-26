package compressionmanager

import (
	"GoAutoExtractor/antivirus"
	"GoAutoExtractor/compression"
	"GoAutoExtractor/filewatch"
	"GoAutoExtractor/regextools"
	"GoAutoExtractor/statuschecker"
)

type Builder struct {
	extractor     compression.DecompressorInterface
	filewatcher   filewatch.FileWatcherInterface
	regexTool     regextools.RegexToolInterface
	antivirus     antivirus.AntiVirusInterface
	statuschecker statuschecker.StatusCheckerInterface
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetDecompressor(iextractor compression.DecompressorInterface) {
	b.extractor = iextractor
}

func (b *Builder) SetExtensionSanitizer(iextension regextools.RegexToolInterface) {
	b.regexTool = iextension
}

func (b *Builder) SetFileWatcher(ifilewatcher filewatch.FileWatcherInterface) {
	b.filewatcher = ifilewatcher
}

func (b *Builder) SetAntivirus(iscanner antivirus.AntiVirusInterface) {
	b.antivirus = iscanner
}

func (b *Builder) SetStatusChecker(istatuschecker statuschecker.StatusCheckerInterface) {
	b.statuschecker = istatuschecker
}

func (b *Builder) Build() *CompressionManager {

	//Initialize interfaces to default implementations if not set.
	if b.extractor == nil {
		b.extractor = &compression.HashigoExtractor{}
	}
	if b.filewatcher == nil {
		b.filewatcher = &filewatch.FSNotifyWatcher{}
	}
	if b.regexTool == nil {
		b.regexTool = &regextools.RegexTool{}
	}
	if b.antivirus == nil {
		b.antivirus = &antivirus.ClamAntiVirus{}
	}
	if b.statuschecker == nil {
		b.statuschecker = &statuschecker.SyncthingStatusChecker{}
	}

	return &CompressionManager{
		antivirus:     b.antivirus,
		extractor:     b.extractor,
		filewatcher:   b.filewatcher,
		regexTool:     b.regexTool,
		statuschecker: b.statuschecker,
	}
}
