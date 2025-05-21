package main

import (
	compressionmanager "MediaCompressionManager/compression-manager"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	builder := compressionmanager.NewBuilder()
	compressionmanager := builder.Build()

	//Initialize interfaces.
	extractor := compressionmanager.Extractor
	sanitizer := compressionmanager.Sanitizer
	scanner := compressionmanager.Scanner
	statuschecker := compressionmanager.Statuschecker

	inputFile := os.Args[1]
	folderID := os.Getenv("FOLDER_ID")
	sanitizedFileName := sanitizer.RemoveExtension(inputFile)
	outputDir := sanitizedFileName
	syncTimeoutSeconds := getSyncTimeoutSetting()

	//Wait for the sync to finish before continuing.
	fmt.Println("Waiting for folder to finish syncing.")
	if err := statuschecker.WaitForSync(folderID, syncTimeoutSeconds); err != nil {
		log.Fatal(err)
	}

	//Scan the file for viruses.
	fmt.Println("Scanning compressed file.")
	scanResult := scanner.ScanFile(inputFile)
	if scanResult.VirusFound {
		log.Fatal("Virus found in compressed file:", scanResult.VirusDescription)
	} else if scanResult.Error != nil {
		log.Fatal("Error during scan:", scanResult.Error)
	}

	//Extract the file.
	fmt.Println("Decompressing.")
	if err := extractor.Decompress(inputFile, outputDir); err != nil {
		log.Fatal("Decompression failed:", err)
	}

	//TODO: Move the file to the output directory and possibly delete it.
	logEntry := fmt.Sprintf("File %s has been extracted to %s.", inputFile, outputDir)
	fmt.Println(logEntry)
}

func getSyncTimeoutSetting() int {

	syncTimeoutSecondsStr := os.Getenv("SYNC_TIMEOUT_SECONDS")
	var timeoutSeconds int

	if syncTimeoutSecondsStr == "" {
		timeoutSeconds = 60
	} else {
		var err error
		timeoutSeconds, err = strconv.Atoi(syncTimeoutSecondsStr)

		if err != nil {
			log.Fatal("Invalid SYNC_TIMEOUT_SECONDS:", err)
			timeoutSeconds = 60
		}
	}

	return timeoutSeconds
}
