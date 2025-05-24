package filewatch

type EventType int

// These are the types of events that can be detected by the file watcher.
const (
	Unknown EventType = iota
	CreateFile
	CreateDirectory
)
