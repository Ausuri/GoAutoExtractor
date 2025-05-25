package configmanager

import (
	"errors"
	"log"

	"GoAutoExtractor/utils"

	"github.com/spf13/viper"
)

type ConfigManagerViper struct {
	vpr      *viper.Viper
	settings ConfigObjects
}

func (v *ConfigManagerViper) Init() *ConfigManagerViper {
	err := v.createViperInstance()
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (v *ConfigManagerViper) GetSetting(settingName string) (any, error) {

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

func (v *ConfigManagerViper) createViperInstance() error {

	if (v.vpr) != nil {
		return nil
	}

	viper.SetConfigFile(USER_CONFIG_PATH)
	v.mapDefaultFieldsToViper()
	viper.AutomaticEnv()

	v.vpr = viper.GetViper()

	return nil
}

func (v *ConfigManagerViper) mapDefaultFieldsToViper() {

	defaultFields := utils.GetObjectMap(v.settings.defaultConfig)
	for key, value := range defaultFields {
		viper.SetDefault(key, value)
	}

}
