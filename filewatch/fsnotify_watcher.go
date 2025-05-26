package filewatch

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FSNotifyWatcher struct{}

func (f *FSNotifyWatcher) MonitorCreatedFiles(folderPath string, watchSubDirectories bool) *FileWatcherChannels {

	result := runMonitor(folderPath, watchSubDirectories, EventType(CreateFile))
	return result

}

func (f *FSNotifyWatcher) MonitorCreatedDirectories(folderPath string, watchSubDirectories bool) *FileWatcherChannels {

	result := runMonitor(folderPath, watchSubDirectories, EventType(CreateDirectory))
	return result

}

func runMonitor(folderPath string, watchSubDirectories bool, eventType EventType) *FileWatcherChannels {

	watcher := initializeWatcher(folderPath, watchSubDirectories)
	eventChannel := make(chan string)
	errorChannel := make(chan error)
	channelObj := FileWatcherChannels{
		Error:         errorChannel,
		EventDetected: eventChannel,
	}

	//TODO: Add a way to stop the watcher gracefully.
	go func() {
		for {
			if event := <-watcher.Events; event.Op == fsnotify.Create {

				evType, evErr := getEventType(event.Name)
				if evErr != nil {
					log.Println("Error getting event type:", evErr)
					continue
				}

				if evType != eventType {
					continue
				}

				fmt.Println("Created file:", event.Name)
				eventChannel <- event.Name
			}
		}
	}()

	return &channelObj
}

func initializeWatcher(folderPath string, watchSubDirectories bool) *fsnotify.Watcher {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Watch this directory
	dirToWatch := folderPath
	err = watcher.Add(dirToWatch)
	if err != nil {
		log.Fatal(err)
	}

	if watchSubDirectories {
		subdirectories, err := getSubDirectories(folderPath)

		if err != nil {
			log.Fatal(err)
		}

		for _, subdirectory := range subdirectories {
			err = watcher.Add(subdirectory)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return watcher
}
