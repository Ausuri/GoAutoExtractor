package filewatch

type FileWatcherChannels struct {
	Error         chan error
	EventDetected chan string
}

// FileWatcherInterface defines the methods for monitoring file and directory creation events.
type FileWatcherInterface interface {

	// Detects new files in the specified directory, optionally including subdirectories. The channel will relay the path of the new file.
	MonitorCreatedFiles(folderPath string, watchSubDirectories bool) *FileWatcherChannels

	// Detects new folders in the specified directory, optionally including subdirectories. The channel will relay the path of the new folder.
	MonitorCreatedDirectories(folderPath string, watchSubDirectories bool) *FileWatcherChannels
}
