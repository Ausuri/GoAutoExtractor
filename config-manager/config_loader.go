package configmanager

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type configFilePaths struct {
	DefaultConfigPath     string
	EnvironmentConfigPath string
	UserConfigPath        string
}

// Initializes all the configurations. This should be used by most anything that implements the config manager interface.
func loadAllConfigs(configPaths *configFilePaths) (*configObjects, error) {

	var err error
	settings := configObjects{}

	settings.userConfig, err = loadJSONFromFile(configPaths.UserConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error loading user config: %s", err)
	}
	settings.defaultConfig, err = loadJSONFromFile(configPaths.DefaultConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error loading default config: %s", err)
	}
	settings.envConfig, err = mapEnvironmentFile(configPaths.EnvironmentConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error loading environment config: %s", err)
	}

	settings.createMapFromConfigObjects()
	return &settings, nil
}

func loadJSONFromFile(filePath string) (*JsonConfig, error) {

	// Open the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse file to JSON object.
	var configJson JsonConfig
	errJ := json.Unmarshal(file, &configJson)
	if errJ != nil {
		return nil, errJ
	}

	return &configJson, nil
}

// Maps the environment variables to an object.
func mapEnvironmentFile(environmentFilePath string) (*environmentConfig, error) {

	envMap, err := godotenv.Read(environmentFilePath)
	if err != nil {
		return nil, err
	}

	envConfig := environmentConfig{}

	//TODO: May need to map the config files dynamically if this list gets large.
	envConfig.OutputPath = envMap["OUTPUT_DIR"]
	envConfig.SyncthingAPIKey = envMap["SYNCTHING_API_KEY"]
	envConfig.SyncthingFolderID = envMap["SYNCTHING_FOLDER_ID"]
	envConfig.SyncthingAPIEndpoint = envMap["SYNCTHING_API_ENDPOINT"]

	return &envConfig, nil
}
