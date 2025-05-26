package configmanager

type configDevLocation struct{}

func (cdl *configDevLocation) getPaths() *configFilePaths {

	result := configFilePaths{}
	result.DefaultConfigPath = "./config/default_config.json"
	result.EnvironmentConfigPath = "./env"
	result.UserConfigPath = "./config/config.json"

	return &result
}
