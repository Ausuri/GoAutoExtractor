package filewatch

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FSNotifyWatcher struct{}

type runMonitorArgs struct {
	errorChannel        chan error
	eventChannel        chan string
	eventType           EventType
	folderPath          string
	watchSubDirectories bool
}

func (f *FSNotifyWatcher) MonitorCreatedFiles(folderPath string, watchSubDirectories bool) *FileWatcherChannels {

	runArgs := runMonitorArgs{
		errorChannel:        make(chan error),
		eventChannel:        make(chan string),
		eventType:           CreateFile,
		folderPath:          folderPath,
		watchSubDirectories: watchSubDirectories,
	}

	channels := &FileWatcherChannels{
		Error:         runArgs.errorChannel,
		EventDetected: runArgs.eventChannel,
	}

	go runMonitor(runArgs)
	return channels
}

func (f *FSNotifyWatcher) MonitorCreatedDirectories(folderPath string, watchSubDirectories bool) *FileWatcherChannels {
	runArgs := runMonitorArgs{
		errorChannel:        make(chan error),
		eventChannel:        make(chan string),
		eventType:           CreateDirectory,
		folderPath:          folderPath,
		watchSubDirectories: watchSubDirectories,
	}

	channels := &FileWatcherChannels{
		Error:         runArgs.errorChannel,
		EventDetected: runArgs.eventChannel,
	}

	go runMonitor(runArgs)
	return channels
}

func runMonitor(args runMonitorArgs) {

	watcher := initializeWatcher(args.folderPath, args.watchSubDirectories)
	eventChannel := args.eventChannel
	errorChannel := args.errorChannel

	//TODO: Add a way to stop the watcher gracefully.
	go func() {
		for {
			if event := <-watcher.Events; event.Op == fsnotify.Create {

				evType, evErr := getEventType(event.Name)
				if evErr != nil {
					log.Println("Error getting event type:", evErr)
					errorChannel <- evErr
					continue
				}

				if evType != args.eventType {
					continue
				}

				fmt.Println("Created file:", event.Name)
				eventChannel <- event.Name
			}
		}
	}()

	// The channel must be blocked indefinitely
	<-make(chan struct{})
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
