package configmanager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"GoAutoExtractor/utils"
)

type appConfigManager struct {
	configManager  configManagerInterface
	configLocation configLocationInterface
}

var appConfig *appConfigManager

// The most important function regarding config files - This must be called at the start of the app (or integration tests)
func InitializeConfig(configType ConfigManagerType) {
	appConfig = intializeAppConfig(configType)
}

// Does a type conversion using the configManager interface provided to grab the setting.
func GetSetting[T Primitive](settingName string) T {

	var zero T

	setting, err := appConfig.configManager.getSetting(settingName)
	if err != nil {
		fmt.Printf("Error getting setting: %s", settingName)
		return zero
	}

	value, ok := setting.(T)
	if !ok {
		fmt.Printf("Error converting setting: %s", settingName)
		return zero
	}

	return value
}

// Maps JSON file objects to maps for key reference. JSON files must be instantiated before calling this.
func (cf *configObjects) createMapFromConfigObjects() error {

	if cf.defaultConfig == nil || cf.envConfig == nil {
		return errors.New("createMapFromConfigObjects(): config files have not been set")
	}

	cf.defaultConfigMap = utils.GetObjectMap(cf.defaultConfig)
	cf.envConfigMap = utils.GetObjectMap(cf.envConfig)

	//User config file is allowed to be null.
	if cf.userConfig != nil {
		cf.userConfigMap = utils.GetObjectMap(cf.userConfig)
	}

	return nil
}

func intializeAppConfig(configType ConfigManagerType) *appConfigManager {

	acm := appConfigManager{}
	var useDevPaths bool
	devEnvPath := os.Getenv("USE_DEV_CONFIG_PATHS")

	if devEnvPath != "" {
		var dErr error
		useDevPaths, dErr = strconv.ParseBool(devEnvPath)
		if dErr != nil {
			log.Print("error parsing dev path environment field")
			useDevPaths = false
		}
	} else {
		useDevPaths = false
	}

	if useDevPaths {
		acm.configLocation = &configDevLocation{}
	} else {
		acm.configLocation = &configLocationProduction{}
	}

	var initErr error
	var cObj *configObjects
	configFilePaths := acm.configLocation.getPaths()
	cObj, initErr = loadAllConfigs(configFilePaths)

	if initErr != nil {
		log.Fatalf("error initializing config file: %v", initErr)
	}

	switch configType {
	case ConfigManagerType(UnknownConfigManagerType):
		acm.configManager = &goexConfigManager{settings: cObj}
	case ConfigManagerType(GoexConfigManagerType):
		acm.configManager = &goexConfigManager{settings: cObj}
	case ConfigManagerType(ViperConfigManagerType):
		acm.configManager = &configManagerViper{}
	default:
		acm.configManager = &goexConfigManager{settings: cObj}
	}

	return &acm
}
