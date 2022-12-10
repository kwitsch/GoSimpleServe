package config

const configFilePath = "/config_template.yaml"

var (
	isVerbose         = false
	filesEPEnabled    = false
	hasConfigTemplate = false
	configFile        = ""
)

func init() {
	isVerbose = getEnvBool("VERBOSE", false)
	filesEPEnabled = getEnvBool("ENDPOINT_FILES", false)
	configFile, hasConfigTemplate = readConfig(configFilePath)
}

func IsVerbose() bool {
	return isVerbose
}

func FilesEndpointEnabled() bool {
	return filesEPEnabled
}

func HasConfigTemplate() bool {
	return hasConfigTemplate
}

func GetConfig() string {
	return configFile
}
