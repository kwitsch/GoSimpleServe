//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

// Variable type ENUM(
// string // String
// bool   // Boolean
// int    // Integer
// array  // Array
// )
type VarType uint8

type ConfigField struct {
	EnvVariable  string  `yaml:"envVariable"`
	DefaultValue string  `yaml:"defaultValue"`
	VariableType VarType `yaml:"variableType" default:"string"`
	Separator    string  `yaml:"separator" default:" "`
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
	if err := defaults.Set(&cfgf); err != nil {
		return "", false
	}

	err = yaml.UnmarshalStrict(data, &cfgf)
	if err != nil {
		return "", false
	}

	return buildJson(&cfgf)
}

func buildJson(cfgf *ConfigFile) (string, bool) {
	iterator := 0
	var writeBuf bytes.Buffer
	fmt.Fprint(&writeBuf, "{ ")
	for name, val := range cfgf.Fields {
		if iterator > 0 {
			fmt.Fprint(&writeBuf, ", ")
		}

		if sval, err := val.String(); err == nil {
			fmt.Fprintf(&writeBuf, "\"%s\": %s", name, sval)
		}

		iterator++
	}

	fmt.Fprint(&writeBuf, "}")

	var out bytes.Buffer
	if err := json.Indent(&out, writeBuf.Bytes(), "", "  "); err == nil {
		return out.String(), true
	}

	return "", false
}

func (f *ConfigField) String() (string, error) {
	switch f.VariableType {
	case VarTypeString:
		return fmt.Sprintf("\"%s\"", getEnvString(f.EnvVariable, f.DefaultValue)), nil
	case VarTypeBool:
		if def, err := strconv.ParseBool(f.DefaultValue); err == nil {
			return fmt.Sprintf("%t", getEnvBool(f.EnvVariable, def)), nil
		}
	case VarTypeInt:
		if def, err := strconv.Atoi(f.DefaultValue); err == nil {
			return fmt.Sprintf("%d", getEnvInt(f.EnvVariable, def)), nil
		}
	case VarTypeArray:
		return getEnvArrayString(f.EnvVariable, f.DefaultValue, f.Separator), nil
	}

	return "", errors.New("Field Error")
}
