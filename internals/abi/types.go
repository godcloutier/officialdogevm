package abi

import (
	"fmt"
)

type ParamType string

const (
	Integer   ParamType = "int"
	Integer64 ParamType = "int64"
	Boolean   ParamType = "bool"
	String    ParamType = "string"
	Void      ParamType = "void"
)

type Type struct {
	Type ParamType
}

func NewType(paramType string) (Type, error) {
	typ := Type{}

	switch paramType {
	case "int":
		typ.Type = Integer
	case "int64":
		typ.Type = Integer64
	case "bool":
		typ.Type = Boolean
	case "string":
		typ.Type = String
	case "void":
		typ.Type = Void
	default:
		return Type{}, fmt.Errorf("unsupported arg type: %s", paramType)
	}

	return typ, nil
}