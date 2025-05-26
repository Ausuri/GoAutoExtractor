package filewatch

import (
	"GoAutoExtractor/utils"
	"errors"
	"fmt"
	"time"
)

type MockFileWatcher struct {
	DirectoryPathFound string
	FilePathFound      string
	FindDirectory      bool
	FindFile           bool
	LookupMSTime       int
	RoutinePauseMSTime int64
	StopRoutines       chan any
	ThrowError         bool
}

func (mfw *MockFileWatcher) MonitorCreatedFiles(folderPath string, watchSubDirectories bool) *FileWatcherChannels {
	result := mfw.runMonitorMock()
	return result
}

func (mfw *MockFileWatcher) MonitorCreatedDirectories(folderPath string, watchSubDirectories bool) *FileWatcherChannels {
	result := mfw.runMonitorMock()
	return result
}

func (mfw *MockFileWatcher) runMonitorMock() *FileWatcherChannels {

	eventChannel := make(chan string)
	errorChannel := make(chan error)
	channels := FileWatcherChannels{
		EventDetected: eventChannel,
		Error:         errorChannel,
	}

	go func() {
		select {

		case <-mfw.StopRoutines:
			{
				fmt.Println("mock directory watcher -> stop routine command received")
				return
			}
		default:
			if mfw.ThrowError {
				errorChannel <- errors.New("mock directory watcher -> error occured")
			}

			if mfw.FindDirectory {
				fmt.Printf("mock directory watcher -> sleeping for %d milliseconds", mfw.LookupMSTime)
				time.Sleep(time.Duration(mfw.LookupMSTime))
				eventChannel <- mfw.DirectoryPathFound
			}

			utils.PauseMilliseconds(mfw.RoutinePauseMSTime)
		}
	}()

	return &channels
}
