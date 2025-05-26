package filewatch

// These are the types of events that can be detected by the file watcher.
type EventType int

const (
	UnknownEventType EventType = iota
	CreateFile
	CreateDirectory
)
