package config

import (
	"errors"
	"os"
	"strings"
)

const configFilePath = "/config_template.yaml"

var (
	isVerbose         = false
	hasConfigTemplate = false
	configFile        = ""
)

func init() {
	isVerbose = (strings.TrimSpace(strings.ToLower(os.Getenv("VERBOSE"))) == "true")
	configFile, hasConfigTemplate = readConfig(configFilePath)
}

func IsVerbose() bool {
	return isVerbose
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
