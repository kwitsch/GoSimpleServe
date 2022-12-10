//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names
package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// Variable type ENUM(
// string // String
// bool   // Boolean
// int   // Integer
// )
type VarType uint8

type ConfigField struct {
	EnvVariable  string  `yaml:"envVariable"`
	DefaultValue string  `yaml:"defaultValue"`
	VariableType VarType `yaml:"variableType"`
}

type ConfigFile struct {
	Fields map[string]ConfigField `yaml:",inline"`
}

func readConfig(file string) (string, bool) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return "", false
	}

	var data []byte
	data, err := os.ReadFile(file)
	if err != nil {
		return "", false
	}

	var cfgf ConfigFile
	err = yaml.UnmarshalStrict(data, &cfgf)
	if err != nil {
		return "", false
	}

	res := ""

	return res, true
}
