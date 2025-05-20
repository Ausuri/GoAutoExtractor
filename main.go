package main

import (
	regexextension "MediaCompressionManager/extensionsanitizer"
	"MediaCompressionManager/scanner"
	statuscheck "MediaCompressionManager/status"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputFile string = os.Args[1]
	var folderID string = os.Getenv("FOLDER_ID")
	var sanitizedFileName string = regexextension.RemoveExtension(inputFile)
	var outputDir string = sanitizedFileName

	//Wait for the sync to finish before continuing.
	fmt.Println("Waiting for folder to finish syncing.")
	if err := statuscheck.WaitForSync(folderID); err != nil {
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
	extractor := decompresser.HashigoExtractor{}
	if err := extractor.Decompress(inputFile, outputDir); err != nil {
		log.Fatal("Decompression failed:", err)
	}

	//TODO: Move the file to the output directory and possibly delete it.
	logEntry := fmt.Sprintf("File %s has been extracted to %s.", inputFile, outputDir)
	fmt.Println(logEntry)
}
