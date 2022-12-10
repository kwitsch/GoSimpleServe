package config

import (
	"errors"
	"os"
	"strings"
)

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

func readConfig(file string) (string, bool) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return "", false
	}

	res := ""

	return res, true
}

func getEnvBool(varname string, defaultval bool) bool {
	val, exists := os.LookupEnv(varname)
	if exists {
		return (strings.TrimSpace(strings.ToLower(val)) == "true")
	}

	return defaultval
}
