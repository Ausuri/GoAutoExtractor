package main

import (
	"fmt"
	"log"

	"MediaCompressionManager/decompresser"
	"MediaCompressionManager/scanner"
	"MediaCompressionManager/status"
)

func main() {
	inputFile := "example.zip"
	outputDir := inputFile + "_extracted"
	folderID := "your-folder-id"

	fmt.Println("🟡 Waiting for folder to finish syncing...")
	if err := statuscheck.WaitForSync(folderID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("🔍 Scanning compressed file...")
	if err := scanner.ScanFile(inputFile); err != nil {
		log.Fatal("Scan failed:", err)
	}

	fmt.Println("📂 Decompressing...")
	if err := decompresser.Decompress(inputFile, outputDir); err != nil {
		log.Fatal("Decompression failed:", err)
	}

	fmt.Println("🔍 Scanning extracted folder...")
	if err := scanner.ScanFile(outputDir); err != nil {
		log.Fatal("Post-scan failed:", err)
	}

	fmt.Println("✅ All steps completed.")
}
