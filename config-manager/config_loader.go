package configmanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type configFilePaths struct {
	AcceptedExtensionFilePath string
	DefaultConfigPath         string
	EnvironmentConfigPath     string
	UserConfigPath            string
}

// Parses file line by line to get list of items (1 item = 1 line)
func getSupportedFileExtensions(acceptedListFilePath string) ([]string, error) {

	fileSlice, err := parseExtensionFile(acceptedListFilePath)
	if err != nil {
		return nil, err
	}

	return fileSlice, nil
}

// Initializes all the configurations. This should be used by most anything that implements the config manager interface.
func loadAllConfigs(configPaths *configFilePaths) (*configObjects, error) {

	var err error
	settings := configObjects{}

	settings.allowedExtensions, err = getSupportedFileExtensions(configPaths.AcceptedExtensionFilePath)
	if err != nil {
		return nil, fmt.Errorf("error loading extension file list: %v", err)
	}
	settings.userConfig, err = loadJSONFromFile(configPaths.UserConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error loading user config: %v", err)
	}
	settings.defaultConfig, err = loadJSONFromFile(configPaths.DefaultConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error loading default config: %v", err)
	}
	settings.envConfig, err = mapEnvironmentFile(configPaths.EnvironmentConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error loading environment config: %v", err)
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
func mapEnvironmentFile(environmentFilePath string) (*EnvironmentConfig, error) {

	envMap, err := godotenv.Read(environmentFilePath)
	if err != nil {
		return nil, err
	}

	envConfig := EnvironmentConfig{}

	//TODO: May need to map the config files dynamically if this list gets large.
	envConfig.OutputPath = envMap["OUTPUT_DIR"]
	envConfig.SyncthingAPIKey = envMap["SYNCTHING_API_KEY"]
	envConfig.SyncthingFolderID = envMap["SYNCTHING_FOLDER_ID"]
	envConfig.SyncthingAPIEndpoint = envMap["SYNCTHING_API_ENDPOINT"]

	return &envConfig, nil
}

func parseExtensionFile(filePath string) ([]string, error) {

	fileSlice := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileSlice = append(fileSlice, line)
	}

	return fileSlice, nil
}
