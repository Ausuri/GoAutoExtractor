package configmanager

import (
	"errors"
	"log"

	"GoAutoExtractor/utils"

	"github.com/spf13/viper"
)

type configManagerViper struct {
	configPaths *configFilePaths
	settings    configObjects
	vpr         *viper.Viper
}

func (v *configManagerViper) Init(configPaths *configFilePaths) *configManagerViper {
	v.configPaths = configPaths
	err := v.createViperInstance()
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (v *configManagerViper) getSetting(settingName string) (any, error) {

	if v.vpr == nil {
		v.createViperInstance()
	}

	settingValue := v.vpr.Get(settingName)
	if settingValue == nil {
		err := errors.New("setting not found: " + settingName)
		return nil, err // or return an error if you prefer
	}

	return settingValue, nil
}

func (v *configManagerViper) createViperInstance() error {

	if (v.vpr) != nil {
		return nil
	}

	viper.SetConfigFile(v.configPaths.UserConfigPath)
	v.mapDefaultFieldsToViper()
	viper.AutomaticEnv()

	v.vpr = viper.GetViper()

	return nil
}

func (v *configManagerViper) mapDefaultFieldsToViper() {

	defaultFields := utils.GetObjectMap(v.settings.defaultConfig)
	for key, value := range defaultFields {
		viper.SetDefault(key, value)
	}

}
