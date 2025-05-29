package configmanager

import "GoAutoExtractor/utils"

// Creates a default unit test settings object.
func CreateTestConfig() (*JsonConfig, *EnvironmentConfig) {

	unitTestJsonSettings := JsonConfig{
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

	unitTestEnvSettings := EnvironmentConfig{
		OutputPath:           "./tmp/goextests/output/",
		SyncthingAPIKey:      "1234567890",
		SyncthingFolderID:    "1234567890",
		SyncthingAPIEndpoint: "https://localhost:8384",
		UseDevConfigPaths:    true,
		WatchPath:            "./tmp/goextests/",
		WatchSubfolders:      true,
	}

	return &unitTestJsonSettings, &unitTestEnvSettings
}

// Use this initializer for unit tests outside of this package where you want to specify the settings in the constructor, or use nil for default settings.
func InitializeTestConfig(overrideSettingsMap map[string]any) {

	// Create default setting map, and override keys with any values provided by argument(s)
	defaultJsonSettings, defaultEnvSettings := CreateTestConfig()
	defaultJsonSlice := utils.GetObjectMap(defaultJsonSettings)
	defaultEnvSlice := utils.GetObjectMap(defaultEnvSettings)
	settingsMap := utils.MergeMaps(defaultEnvSlice, defaultJsonSlice)

	for key, value := range overrideSettingsMap {
		settingsMap[key] = value
	}

	// Set mock config manager.
	mockManager := mockConfigManager{}
	mockManager.MockSettings = settingsMap

	// Set app config manager.
	appConfig = &appConfigManager{}
	appConfig.configManager = &mockManager
	appConfig.configLocation = &configDevLocation{}

	// Get extensions list from file.
	var extensionFileError error
	configPaths := appConfig.configLocation.getPaths()
	appConfig.configObjects = &configObjects{}
	appConfig.configObjects.allowedExtensions, extensionFileError = getSupportedFileExtensions(configPaths.AcceptedExtensionFilePath)
	if extensionFileError != nil {
		panic(extensionFileError)
	}

}
