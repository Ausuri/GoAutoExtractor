package configmanager

import (
	"errors"
	"fmt"
)

type GoexConfigManager struct {
	settings *ConfigObjects
}

func (gcm *GoexConfigManager) GetSetting(settingName string) (any, error) {

	if gcm.settings.defaultConfigMap == nil || gcm.settings.envConfigMap == nil {
		return nil, errors.New("config not initialized")
	}

	if gcm.settings.envConfigMap[settingName] != nil {
		return gcm.settings.envConfigMap[settingName], nil
	}

	if gcm.settings.userConfig != nil && gcm.settings.userConfigMap[settingName] != nil {
		return gcm.settings.userConfigMap[settingName], nil
	}

	if gcm.settings.defaultConfigMap != nil && gcm.settings.defaultConfigMap[settingName] != nil {
		return gcm.settings.defaultConfigMap[settingName], nil
	}

	notFoundError := fmt.Errorf("setting %s not found", settingName)
	return nil, notFoundError
}
