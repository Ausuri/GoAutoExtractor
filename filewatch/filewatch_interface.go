package filewatch

type FileWatcherInterface interface {
	DetectNewFile(folderPath string, watchSubDirectories bool) (filePath string, err error)        // Detects file changes in the specified directory.
	DetectNewFolder(folderPath string, watchSubDirectories bool) (newFolderPath string, err error) // Detects new folders in the specified directory.
}
