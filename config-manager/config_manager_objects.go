package configmanager

type JSONConfig struct {
	ClamscanBinary          string `json:"clamscan_binary"`
	Concurrency             int    `json:"concurrency"`
	DeleteAfterExtraction   bool   `json:"delete_after_extraction"`
	EnableClamscan          bool   `json:"enable_clamscan"`
	EnableSyncthing         bool   `json:"enable_syncthing"`
	LogLevel                string `json:"log_level"`
	OutputPath              string `json:"output_path"`
	SyncthingAPIEndpoint    string `json:"syncthing_endpoint"`
	SyncthingTimeoutSeconds int    `json:"syncthing_timeout_seconds"`
	WatchPath               string `json:"watch_path"`
	WatchSubfolders         bool   `json:"watch_subfolders"`
}

type EnvironmentConfig struct {
	SyncthingAPIKey      string
	SyncthingFolderID    string
	SyncthingAPIEndpoint string
	OutputPath           string
}

type ConfigObjects struct {
	defaultConfig    *JSONConfig
	defaultConfigMap map[string]any
	envConfig        *EnvironmentConfig
	envConfigMap     map[string]any
	userConfig       *JSONConfig
	userConfigMap    map[string]any
}
