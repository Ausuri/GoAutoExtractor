package main

import (
	compressionmanager "GoAutoExtractor/compression-manager"
	"GoAutoExtractor/utils"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

type GoAutoExtractorSettings struct {
}

// This is a channel to signal the application to exit to all listeners in go routines.
var appExitChannel = make(chan struct{})

func main() {

	godotenv.Load()

	var signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGQUIT)

	// Setup our channel to close the application when needed.
	go func() {
		<-signalChannel
		close(appExitChannel)
	}()

	runOnce := flag.Bool("once", false, "Run one-time extraction instead of daemon mode")
	inputFile := flag.String("extract", "", "Manually extract a file and exit")

	builder := compressionmanager.NewBuilder()
	compressionmanager := builder.Build()

	//If inputFile is set, then we're just decompressing a single file.
	if (inputFile != nil) && (*inputFile != "") {
		if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
			log.Fatal("Input file does not exist:", *inputFile)
		}

		err := compressionmanager.ScanAndDecompressFile(*inputFile)
		if err != nil {
			log.Fatal("Error during decompression:", err)
		}
		return
	}

	if *runOnce {
		//TODO: Implement run once mode
		panic("unimplemented")
	}

	ctx, cancel := context.WithCancel(context.Background())
	cm := compressionmanager
	runDaemon(ctx, cm)

	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt, syscall.SIGTERM)
	<-osInterruptChannel
	log.Println("Received interrupt signal, shutting down...")
	cancel() // Cancel the context to stop the daemon

	utils.PauseSeconds(1)
}

func runDaemon(ctx context.Context, cm *compressionmanager.CompressionManager) {

	channels, err := cm.RunMonitor()
	if err != nil {
		panic("Could not create channels")
	}

	for range ctx.Done() {

		select {
		case file := <-channels.EventDetected:

		}
	}
}
