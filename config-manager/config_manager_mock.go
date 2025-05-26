package configmanager

import (
	"errors"
	"fmt"
)

type mockConfigManager struct {
	MockSettings map[string]any
}

func (mcm *mockConfigManager) getSetting(settingName string) (any, error) {

	if mcm.MockSettings != nil && mcm.MockSettings[settingName] != nil {
		return mcm.MockSettings[settingName], nil
	}

	if mcm.MockSettings == nil {
		return nil, errors.New("mock settings map is null")
	}

	return nil, fmt.Errorf("setting not found: %s", settingName)
}
