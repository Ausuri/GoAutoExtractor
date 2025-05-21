package statuschecker

type StatusCheckerInterface interface {
	WaitForSync(folderID string, timeoutSeconds int) error
}
