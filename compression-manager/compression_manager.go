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
	"os"
	"strconv"
)

type CompressionManager struct {
	antivirus     antivirus.AntiVirusInterface
	configManager configmanager.ConfigManagerBase
	extractor     compression.DecompressorInterface
	filewatcher   filewatch.FileWatcherInterface
	regexTool     regextools.RegexToolInterface
	statuschecker statuschecker.StatusCheckerInterface

	UserConfig *configmanager.JSONConfig
}

const DEFAULT_TIMEOUT_SECONDS = 60

var StopMonitor chan<- bool

func NewCompressionManager(builder *Builder) *CompressionManager {
	cm := builder.Build()
	return cm
}

func (cm *CompressionManager) RunMonitor() error {

	fileCreatedChannel := make(chan string)

	go func() {
		for {
			select {
			case StopMonitor <- true:
				fmt.Println("Stopping monitor.")
				return
			case newFile := <-fileCreatedChannel:
				fmt.Println("New file detected:", newFile)
			}
		}
	}()

	watchpathSetting, err := cm.configManager.GetSetting("watch_path")
	if err != nil {
		log.Fatal("Error getting watch path from config:", err)
	}

	watchSubDirectoriesSetting, err := cm.configManager.GetSetting("watch_subfolders")
	if err != nil {
		log.Fatal("Error getting watch_subfolders setting from config:", err)
	}

	//Convert settings from any to their expected types.
	pathToWatch, _ := watchpathSetting.(string)
	watchSubDirectories, _ := watchSubDirectoriesSetting.(bool)

	go cm.filewatcher.MonitorCreatedFiles(pathToWatch, watchSubDirectories, fileCreatedChannel)

	return nil
}

func (cm *CompressionManager) ScanAndDecompressFile(inputFile string) error {

	folderID := os.Getenv("FOLDER_ID")
	sanitizedFileName := cm.regexTool.RemoveExtension(inputFile)
	outputDir := sanitizedFileName
	syncTimeoutSeconds := getSyncTimeoutSetting()

	//Wait for the sync to finish before continuing.
	fmt.Println("Waiting for folder to finish syncing.")
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
	logEntry := fmt.Sprintf("File %s has been extracted to %s.", inputFile, outputDir)
	fmt.Println(logEntry)

	return nil
}

func getSyncTimeoutSetting() int {

	syncTimeoutSecondsStr := os.Getenv("SYNC_TIMEOUT_SECONDS")
	var timeoutSeconds int

	if syncTimeoutSecondsStr == "" {
		timeoutSeconds = DEFAULT_TIMEOUT_SECONDS
	} else {
		var err error
		timeoutSeconds, err = strconv.Atoi(syncTimeoutSecondsStr)

		if err != nil {
			log.Fatal("Invalid SYNC_TIMEOUT_SECONDS:", err)
			timeoutSeconds = DEFAULT_TIMEOUT_SECONDS
		}
	}

	return timeoutSeconds
}
