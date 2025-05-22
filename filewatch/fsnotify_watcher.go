package filewatch

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type FSNotifyWatcher struct{}

func (f *FSNotifyWatcher) DetectNewFile(folderPath string, watchSubDirectories bool) (filePath string, err error) {
	watcher, err := fsnotify.NewWatcher()
}

func (f *FSNotifyWatcher) DetectNewFolder(folderPath string, watchSubDirectories bool) (newolderPath string, err error) {
	watcher, err := fsnotify.NewWatcher()

}

func (f *FSNotifyWatcher) initializeWatcher(folderPath string, watchSubDirectories bool) *fsnotify.Watcher {
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
		// Watch all subdirectories
		err = filepath.Walk(dirToWatch, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				err = watcher.Add(path)
				if err != nil {
					log.Fatal(err)
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	return watcher
}
