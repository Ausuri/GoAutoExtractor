package filewatch

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FSNotifyWatcher struct {
	WatcherChannels []*FileWatcherChannels
}

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

	//watcher := initializeWatcher(args.folderPath, args.watchSubDirectories)
	eventChannel := args.eventChannel
	errorChannel := args.errorChannel

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	//TODO: Add a way to stop the watcher gracefully.
	go func() {
		for {
			if event, ok := <-watcher.Events; event.Has(fsnotify.Create) && ok {

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

	err = watcher.Add(args.folderPath)
	if err != nil {
		log.Fatal(err)
	}

	if args.watchSubDirectories {
		subdirectories, err := getSubDirectories(args.folderPath)

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

	// The channel must be blocked indefinitely
	<-make(chan struct{})
}
