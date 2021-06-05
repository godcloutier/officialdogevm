package abi

import (
	"encoding/json"
	"strings"
)

type Argument struct {
	Name string
	Type Type
}

type Arguments []Argument

type ArgumentMarshaling struct {
	Name string
	Type string
}

// UnmarshalJSON implements json.Unmarshaler interface
func (argument *Argument) UnmarshalJSON(data []byte) error {
	var arg ArgumentMarshaling
	err := json.Unmarshal(data, &arg)
	if err != nil {
		return err
	}
	argument.Type, err = NewType(arg.Type)
	if err != nil {
		return err
	}
	argument.Name = arg.Name

	return nil
}

// Pack returns series of Arguments type
// Example
// function foo(int64 a, bool b)
// Pack() extract "int64,bool"
func (arguments Arguments) Pack() string {
	var packedTypes []string

	for _, argument := range arguments {
		packedTypes = append(packedTypes, string(argument.Type.Type))
	}

	return strings.Join(packedTypes, ",")
}