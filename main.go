package main

import (
	compressionmanager "MediaCompressionManager/compression-manager"
	"log"
	"os"
)

func main() {

	inputFile := os.Args[1]
	builder := compressionmanager.NewBuilder()
	compressionmanager := builder.Build()
	err := compressionmanager.ScanAndDecompressFile(inputFile)

	if err != nil {
		log.Fatal("Error during decompression:", err)
	}

}
