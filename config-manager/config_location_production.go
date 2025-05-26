package configmanager

type configLocationProduction struct{}

func (clp *configLocationProduction) getPaths() *configFilePaths {

	result := configFilePaths{}
	result.DefaultConfigPath = "/etc/goautoextractor/default_config.json"
	result.EnvironmentConfigPath = "/etc/goautoextractor/environment_config.json"
	result.UserConfigPath = "/etc/goautoextractor/config.json"

	return &result
}
