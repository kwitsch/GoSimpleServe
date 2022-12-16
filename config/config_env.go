package config

import (
	"os"
	"strconv"
	"strings"
)

func getEnvBool(varname string, defaultval bool) bool {
	if len(varname) > 0 {
		val, exists := os.LookupEnv(varname)
		if exists {
			return (strings.TrimSpace(strings.ToLower(val)) == "true")
		}
	}

	return defaultval
}

func getEnvInt(varname string, defaultval int) int {
	if len(varname) > 0 {
		val, exists := os.LookupEnv(varname)
		if exists {
			if intVal, err := strconv.Atoi(val); err == nil {
				return intVal
			}
		}
	}

	return defaultval
}

func getEnvString(varname string, defaultval string) string {
	if len(varname) > 0 {
		val, exists := os.LookupEnv(varname)
		if exists {
			return strings.TrimSpace(val)
		}
	}

	return defaultval
}
