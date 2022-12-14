// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package config

import (
	"fmt"
	"strings"
)

const (
	// VarTypeString is a VarType of type String.
	// String
	VarTypeString VarType = iota
	// VarTypeBool is a VarType of type Bool.
	// Boolean
	VarTypeBool
	// VarTypeInt is a VarType of type Int.
	// Integer
	VarTypeInt
	// VarTypeArray is a VarType of type Array.
	// Array
	VarTypeArray
)

var ErrInvalidVarType = fmt.Errorf("not a valid VarType, try [%s]", strings.Join(_VarTypeNames, ", "))

const _VarTypeName = "stringboolintarray"

var _VarTypeNames = []string{
	_VarTypeName[0:6],
	_VarTypeName[6:10],
	_VarTypeName[10:13],
	_VarTypeName[13:18],
}

// VarTypeNames returns a list of possible string values of VarType.
func VarTypeNames() []string {
	tmp := make([]string, len(_VarTypeNames))
	copy(tmp, _VarTypeNames)
	return tmp
}

var _VarTypeMap = map[VarType]string{
	VarTypeString: _VarTypeName[0:6],
	VarTypeBool:   _VarTypeName[6:10],
	VarTypeInt:    _VarTypeName[10:13],
	VarTypeArray:  _VarTypeName[13:18],
}

// String implements the Stringer interface.
func (x VarType) String() string {
	if str, ok := _VarTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("VarType(%d)", x)
}

var _VarTypeValue = map[string]VarType{
	_VarTypeName[0:6]:   VarTypeString,
	_VarTypeName[6:10]:  VarTypeBool,
	_VarTypeName[10:13]: VarTypeInt,
	_VarTypeName[13:18]: VarTypeArray,
}

// ParseVarType attempts to convert a string to a VarType.
func ParseVarType(name string) (VarType, error) {
	if x, ok := _VarTypeValue[name]; ok {
		return x, nil
	}
	return VarType(0), fmt.Errorf("%s is %w", name, ErrInvalidVarType)
}

// MarshalText implements the text marshaller method.
func (x VarType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *VarType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseVarType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
