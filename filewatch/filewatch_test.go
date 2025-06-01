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

	fsWatcher := &FSNotifyWatcher{}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "testfile.zip")
	channels := fsWatcher.MonitorCreatedFiles(tmpDir, false)

	fmt.Printf("Created temporary directory: %s", tmpDir)

	go func() {
		utils.PauseSeconds(2)
		_, err := os.Create(filePath)
		if err != nil {
			t.Logf("Failed to create test file: %s", err)
		}
	}()

	if <-channels.EventDetected != "" {
		t.Logf("File created successfully: %s", filePath)
	} else {
		t.Logf("File creation failed: %s", filePath)
		t.Fail()
	}
}
