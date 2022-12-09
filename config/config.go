package config

import (
	"os"
	"strings"
)

var isVerbose = false

func init() {
	isVerbose = (strings.TrimSpace(strings.ToLower(os.Getenv("VERBOSE"))) == "true")
}

func IsVerbose() bool {
	return isVerbose
}
