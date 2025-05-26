package configmanager

var DefaultAppSettings *JsonConfig = &JsonConfig{
	ClamscanBinary:          "clamscan",
	Concurrency:             4,
	DeleteAfterExtraction:   true,
	EnableClamscan:          true,
	EnableSyncthing:         true,
	LogLevel:                "medium",
	OutputPath:              "./tmp/goextests/output/",
	SyncthingAPIEndpoint:    "https://localhost:8384",
	SyncthingTimeoutSeconds: 60,
	WatchPath:               "./tmp/goextests/",
	WatchSubfolders:         true,
}

// Creates a default unit test settings object.
func CreateTestConfig() *JsonConfig {

	unitTestSettings := JsonConfig{
		ClamscanBinary:          "clamscan",
		Concurrency:             4,
		DeleteAfterExtraction:   true,
		EnableClamscan:          true,
		EnableSyncthing:         true,
		LogLevel:                "medium",
		OutputPath:              "./tmp/goextests/output/",
		SyncthingAPIEndpoint:    "https://localhost:8384",
		SyncthingTimeoutSeconds: 60,
		WatchPath:               "./tmp/goextests/",
		WatchSubfolders:         true,
	}

	return &unitTestSettings
}

// Use this initializer for unit tests outside of this package.
func InitializeTestConfigManager(settings map[string]any) {

	mockManager := mockConfigManager{}
	mockManager.MockSettings = settings

	appConfig = &appConfigManager{}
	appConfig.configManager = &mockManager
	appConfig.configLocation = &configDevLocation{}

}
