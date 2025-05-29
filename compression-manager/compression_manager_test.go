package compressionmanager

import (
	"GoAutoExtractor/antivirus"
	"GoAutoExtractor/compression"
	configmanager "GoAutoExtractor/config-manager"
	"GoAutoExtractor/filewatch"
	"GoAutoExtractor/regextools"
	"GoAutoExtractor/statuschecker"
	"fmt"
	"testing"
)

// Builds a CompressionManager using mock interfaces. For mocks that have options, they can be modified through properties after instantation if need be.
func buildMockTester() *CompressionManager {

	builder := Builder{}
	builder.SetAntivirus(&antivirus.MockAntiVirus{})
	builder.SetDecompressor(&compression.MockDecompressor{})
	builder.SetFileWatcher(&filewatch.MockFileWatcher{})
	builder.SetExtensionSanitizer(&regextools.RegexTool{}) //Regex tool doesn't really need a mock - it should work regardless.
	builder.SetStatusChecker(&statuschecker.MockStatusChecker{})

	cm := builder.Build()
	return cm
}

// Loads config files and sets up a CompressionManager for testing. The settings parameter is optional, use if your unit test requires specific config settings.
func initializeCompressionManagerTesting(settingsOverrideMap map[string]any) *CompressionManager {

	configmanager.InitializeTestConfig(settingsOverrideMap)

	//Build a tester with default mock interfaces.
	cm := buildMockTester()
	return cm
}

func TestRunMonitorFile(t *testing.T) {

	stopChannel := make(chan any)
	tester := initializeCompressionManagerTesting(nil)
	tester.filewatcher = &filewatch.MockFileWatcher{
		DirectoryPathFound: "/tmp/nothing/",
		FilePathFound:      "/tmp/nothing/archive.zip",
		LookupMSTime:       250,
		RoutinePauseMSTime: 500,
		StopRoutines:       stopChannel,
		ThrowError:         false,
	}
	channels, err := tester.RunMonitor()

	if err != nil {
		t.Fatalf("error occurred in RunMonitor() %v", err)
	}

	select {
	case file := <-channels.EventDetected:
		fmt.Printf("Received channel item: %s", file)
	case chError := <-channels.Error:
		t.Fatalf("Received channel error %v", chError)
	}

	stopChannel <- true //Stop the mock watcher.

}

func TestScanAndDecompressFile(t *testing.T) {

	tester := initializeCompressionManagerTesting(nil)
	err := tester.ScanAndDecompressFile("/tmp/nothing/archive.zip")

	if err != nil {
		t.Fatalf("error occurred in ScanAndDecompressFile %v", err)
	}

}
