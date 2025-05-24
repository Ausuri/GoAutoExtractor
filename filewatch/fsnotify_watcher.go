package filewatch

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type FSNotifyWatcher struct{}

func (f *FSNotifyWatcher) DetectNewFile(folderPath string, watchSubDirectories bool, fileDetected chan<- string) error {

	watcher := initializeWatcher(folderPath, watchSubDirectories)

	return nil
}

func (f *FSNotifyWatcher) DetectNewFolder(folderPath string, watchSubDirectories bool, directoryDetected chan<- string) error {

	panic("unimplemented")

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

// Watches for new files (or folders) in the specified directory and send the path to the createEvent channel.
func watchFolder(watcher *fsnotify.Watcher, createEvent chan<- string) {

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Create == fsnotify.Create {
				createEvent <- event.Name
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}
