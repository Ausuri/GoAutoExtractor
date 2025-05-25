package configmanager

import (
	"GoAutoExtractor/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigManagerBase struct {
	settings   ConfigObjects
	GetSetting func(settingName string) (any, error)
}

const DEFAULT_CONFIG_PATH = "/etc/goautoextractor/default_config.json"
const ENVIRONMENT_CONFIG_PATH = "/etc/goautoextractor/environment_config.json"
const USER_CONFIG_PATH = "/etc/goautoextractor/config.json"

// New creates a new instance of ConfigManagerBase, loading the user, default, and environment configurations. onInit is an optional constructor action that can be passed in.
func (cmb *ConfigManagerBase) New(onInit func(c *ConfigManagerBase)) *ConfigManagerBase {

	var userConfigError, defaultConfigError, environmentConfigError error
	cmb.settings.userConfig, userConfigError = cmb.loadJSONFromFile(USER_CONFIG_PATH)
	if userConfigError != nil {
		log.Fatal(fmt.Printf("Error loading user config: %v", userConfigError))
	}

	cmb.settings.defaultConfig, defaultConfigError = cmb.loadJSONFromFile(DEFAULT_CONFIG_PATH)
	if defaultConfigError != nil {
		log.Fatal(fmt.Printf("Error loading default config: %v", defaultConfigError))
	}

	cmb.settings.envConfig, environmentConfigError = cmb.mapEnvironmentFile(ENVIRONMENT_CONFIG_PATH)
	if environmentConfigError != nil {
		log.Fatal(fmt.Printf("Error loading environment config: %v", environmentConfigError))
	}

	cmb.createMapFromConfigObjects()

	//Run any custom initialization function passed in.
	if onInit != nil {
		onInit(cmb)
	}

	return cmb
}

// Maps the environment variables to an object.
func (cmb *ConfigManagerBase) mapEnvironmentFile(environmentFilePath string) (*EnvironmentConfig, error) {

	envMap, err := godotenv.Read(environmentFilePath)
	if err != nil {
		return nil, err
	}

	//TODO: May need to map the config files dynamically if this list gets large.
	cmb.settings.envConfig.OutputPath = envMap["OUTPUT_DIR"]
	cmb.settings.envConfig.SyncthingAPIKey = envMap["SYNCTHING_API_KEY"]
	cmb.settings.envConfig.SyncthingFolderID = envMap["SYNCTHING_FOLDER_ID"]
	cmb.settings.envConfig.SyncthingAPIEndpoint = envMap["SYNCTHING_API_ENDPOINT"]

	return cmb.settings.envConfig, nil
}

func (cmb *ConfigManagerBase) loadJSONFromFile(filePath string) (*JSONConfig, error) {

	// Open the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse file to JSON object.
	var configJson JSONConfig
	errJ := json.Unmarshal(file, &configJson)
	if errJ != nil {
		return nil, errJ
	}

	return &configJson, nil
}

// Maps JSON file objects to maps for key reference. JSON files must be instantiated before calling this.
func (cmb *ConfigManagerBase) createMapFromConfigObjects() error {

	if cmb.settings.defaultConfig == nil || cmb.settings.envConfig == nil {
		return errors.New("createMapFromConfigObjects(): config files have not been set")
	}

	cmb.settings.defaultConfigMap = utils.GetObjectMap(cmb.settings.defaultConfig)
	cmb.settings.envConfigMap = utils.GetObjectMap(cmb.settings.envConfig)

	//User config file is allowed to be null.
	if cmb.settings.userConfig != nil {
		cmb.settings.userConfigMap = utils.GetObjectMap(cmb.settings.userConfig)
	}

	return nil
}
