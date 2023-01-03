package config

import (
	"fmt"
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

func getEnvString(varname, defaultval string) string {
	if len(varname) > 0 {
		val, exists := os.LookupEnv(varname)
		if exists {
			return strings.TrimSpace(val)
		}
	}

	return defaultval
}

func getEnvArrayString(varname, defaultval, separator string) string {
	res := "["
	for _, s := range strings.Split(getEnvString(varname, defaultval), separator) {
		if len(res) > 2 {
			res += ","
		}
		res += fmt.Sprintf("\"%s\"", s)
	}
	res += "]"

	return res
}
