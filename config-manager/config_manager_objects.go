package configmanager

type JsonConfig struct {
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

type environmentConfig struct {
	OutputPath           string
	SyncthingAPIKey      string
	SyncthingFolderID    string
	SyncthingAPIEndpoint string
	UseDevConfigPaths    bool
	WatchPath            string
	WatchSubfolders      bool
}

type configObjects struct {
	defaultConfig    *JsonConfig
	defaultConfigMap map[string]any
	envConfig        *environmentConfig
	envConfigMap     map[string]any
	userConfig       *JsonConfig
	userConfigMap    map[string]any
}

type ConfigManagerType int

const (
	UnknownConfigManagerType ConfigManagerType = iota
	MockConfigManagerType
	GoexConfigManagerType
	ViperConfigManagerType
)
