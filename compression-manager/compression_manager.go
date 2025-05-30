package compressionmanager

import (
	"GoAutoExtractor/antivirus"
	"GoAutoExtractor/compression"
	configmanager "GoAutoExtractor/config-manager"
	"GoAutoExtractor/filewatch"
	"GoAutoExtractor/regextools"
	"GoAutoExtractor/statuschecker"
	"fmt"
	"log"
)

type CompressionManager struct {
	antivirus     antivirus.AntiVirusInterface
	extractor     compression.DecompressorInterface
	filewatcher   filewatch.FileWatcherInterface
	regexTool     regextools.RegexToolInterface
	statuschecker statuschecker.StatusCheckerInterface
}

const DEFAULT_TIMEOUT_SECONDS = 60

var StopMonitor chan<- bool

func NewCompressionManager(builder *Builder) *CompressionManager {
	cm := builder.Build()
	return cm
}

func (cm *CompressionManager) RunMonitor() (*filewatch.FileWatcherChannels, error) {

	//Convert settings from any to their expected types.
	pathToWatch := configmanager.GetSetting[string]("WatchPath")
	watchSubDirectories := configmanager.GetSetting[bool]("WatchSubfolders")

	channels := cm.filewatcher.MonitorCreatedFiles(pathToWatch, watchSubDirectories)

	return channels, nil
}

func (cm *CompressionManager) ScanAndDecompressFile(inputFile string) error {

	folderID := configmanager.GetSetting[string]("SyncthingFolderID")
	sanitizedFileName := cm.regexTool.RemoveExtension(inputFile)
	outputDir := sanitizedFileName
	syncTimeoutSeconds := configmanager.GetSetting[int]("SyncthingTimeoutSeconds")

	//Wait for the sync to finish before continuing.
	fmt.Println("Waiting for folder to finish syncing")
	if err := cm.statuschecker.WaitForSync(folderID, syncTimeoutSeconds); err != nil {
		return err
	}

	//Scan the file for viruses.
	fmt.Println("Scanning compressed file.")
	scanResult := cm.antivirus.ScanFile(inputFile)
	if scanResult.VirusFound {
		log.Fatal("Virus found in compressed file:", scanResult.VirusDescription)
	} else if scanResult.Error != nil {
		log.Fatal("Error during scan:", scanResult.Error)
	}

	//Extract the file.
	fmt.Println("Decompressing.")
	if err := cm.extractor.Decompress(inputFile, outputDir); err != nil {
		return err
	}

	//TODO: Move the file to the output directory and possibly delete it.
	logEntry := fmt.Sprintf("File %v has been extracted to %v.", inputFile, outputDir)
	fmt.Println(logEntry)

	return nil
}
