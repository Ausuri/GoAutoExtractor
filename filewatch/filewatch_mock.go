package filewatch

import (
	"GoAutoExtractor/utils"
	"errors"
	"fmt"
)

type MockFileWatcher struct {
	DirectoryPathFound string
	FilePathFound      string
	LookupMSTime       int64
	RoutinePauseMSTime int64
	StopRoutines       chan any
	ThrowError         bool
}

func (mfw *MockFileWatcher) MonitorCreatedFiles(folderPath string, watchSubDirectories bool) *FileWatcherChannels {
	result := mfw.runMonitorMock(EventType(CreateFile))
	return result
}

func (mfw *MockFileWatcher) MonitorCreatedDirectories(folderPath string, watchSubDirectories bool) *FileWatcherChannels {
	result := mfw.runMonitorMock(EventType(CreateDirectory))
	return result
}

func (mfw *MockFileWatcher) runMonitorMock(eventType EventType) *FileWatcherChannels {

	var stopFlag bool
	eventChannel := make(chan string)
	errorChannel := make(chan error)
	channels := FileWatcherChannels{
		EventDetected: eventChannel,
		Error:         errorChannel,
	}

	go func() {
		//Listen for the stop event.
		<-mfw.StopRoutines
		stopFlag = true
	}()

	go func() {
		for !stopFlag {

			if mfw.ThrowError {
				errorChannel <- errors.New("mock watcher -> mock error occured")
				return
			}

			if eventType == CreateDirectory {
				fmt.Printf("mock directory watcher -> sleeping for %d milliseconds", mfw.LookupMSTime)
				utils.PauseMilliseconds(mfw.LookupMSTime)
				eventChannel <- mfw.DirectoryPathFound
			} else if eventType == CreateFile {
				fmt.Printf("mock file watcher -> sleeping for %d milliseconds", mfw.LookupMSTime)
				utils.PauseMilliseconds(mfw.LookupMSTime)
				eventChannel <- mfw.FilePathFound
			}

			fmt.Printf("mock filer watcher end of goroutine, sleeping for %d milliseconds", mfw.RoutinePauseMSTime)
			utils.PauseMilliseconds(mfw.RoutinePauseMSTime)

		}
	}()

	return &channels
}
