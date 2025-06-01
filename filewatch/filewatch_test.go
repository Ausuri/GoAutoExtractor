package filewatch

import (
	"GoAutoExtractor/utils"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

//TODO: Technically these are integration tests and should be segregated as such.

func TestRunMonitorFile(t *testing.T) {

	blockerChannel := make(chan bool, 1)
	fsWatcher := &FSNotifyWatcher{}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "testfile.zip")
	channels := fsWatcher.MonitorCreatedFiles(tmpDir, true)

	fmt.Printf("Created temporary directory: %s", tmpDir)

	//TODO: The FSNotify interface needs to be refactored to run but also pass back the channels -> see sandbox.
	go func() {
		for {
			select {
			case event := <-channels.EventDetected:
				t.Logf("Received event: %s", event)
				blockerChannel <- true
			case err := <-channels.Error:
				t.Logf("Received error: %s", err)
				t.Fail()
				blockerChannel <- false
			}
		}
	}()

	go func() {
		utils.PauseSeconds(2)
		_, err := os.Create(filePath)
		if err != nil {
			t.Logf("Failed to create test file: %s", err)
		}
	}()

	if <-blockerChannel {
		t.Logf("File created successfully: %s", filePath)
	} else {
		t.Logf("File creation failed: %s", filePath)
		t.Fail()
	}

}
