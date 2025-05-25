package filewatch

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FSNotifyWatcher struct{}

func (f *FSNotifyWatcher) MonitorCreatedFiles(folderPath string, watchSubDirectories bool, fileDetected chan<- string) error {

	watcher := initializeWatcher(folderPath, watchSubDirectories)

	//TODO: Add a way to stop the watcher gracefully.
	go func() {
		for {
			if event := <-watcher.Events; event.Op == fsnotify.Create {

				eventType, evErr := GetEventType(event.Name)
				if evErr != nil {
					log.Println("Error getting event type:", evErr)
					continue
				}

				if eventType != CreateFile {
					continue
				}

				fmt.Println("Created file:", event.Name)
				fileDetected <- event.Name
			}
		}
	}()

	return nil
}

func (f *FSNotifyWatcher) MonitorCreatedDirectories(folderPath string, watchSubDirectories bool, directoryDetected chan<- string) error {

	watcher := initializeWatcher(folderPath, watchSubDirectories)

	//TODO: Add a way to stop the watcher gracefully.
	go func() {
		for {
			if event := <-watcher.Events; event.Op == fsnotify.Create {

				eventType, evErr := GetEventType(event.Name)
				if evErr != nil {
					log.Println("Error getting event type:", evErr)
					continue
				}

				if eventType != CreateDirectory {
					continue
				}

				fmt.Println("Created directory:", event.Name)
				directoryDetected <- event.Name
			}
		}
	}()

	return nil

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
		subdirectories, err := GetSubDirectories(folderPath)

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
