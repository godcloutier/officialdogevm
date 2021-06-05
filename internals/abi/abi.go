package abi

import (
	"encoding/json"
	"fmt"
	"strings"

	"../ast"
)

type ABI struct {
	Methods []Method
}

func New(abiJSON string) (ABI, error) {
	reader := strings.NewReader(abiJSON)
	dec := json.NewDecoder(reader)

	var abi ABI
	if err := dec.Decode(&abi); err != nil {
		return ABI{}, err
	}

	return abi, nil
}

// UnmarshalJSON is implementation of json.Decoder's UnmarshalJSON
func (abi *ABI) UnmarshalJSON(data []byte) error {
	var methods []Method

	if err := json.Unmarshal(data, &methods); err != nil {
		return err
	}

	for _, method := range methods {
		abi.Methods = append(abi.Methods, method)
	}

	return nil
}

// Extract ABI from function of ast.
func ExtractAbiFromFunction(f ast.FunctionLiteral) (Method, error) {
	method := Method{
		Name:      f.Name.String(),
		Arguments: make([]Argument, 0),
		Output:    Argument{},
	}

	args := make([]Argument, 0)

	for _, param := range f.Parameters {
		t, err := convertAstTypeToAbi(param.Type)
		if err != nil {
			return Method{}, err
		}

		arg := Argument{
			Name: param.Identifier.String(),
			Type: t,
		}

		args = append(args, arg)
	}

	method.Arguments = args

	t, err := convertAstTypeToAbi(f.ReturnType)
	if err != nil {
		return Method{}, err
	}

	method.Output = Argument{
		Name: "",
		Type: t,
	}

	return method, nil
}

func convertAstTypeToAbi(p ast.DataStructure) (Type, error) {
	switch p {
	case ast.IntType:
		return NewType("int")
	case ast.StringType:
		return NewType("string")
	case ast.BoolType:
		return NewType("bool")
	case ast.VoidType:
		return NewType("void")
	default:
		return Type{}, fmt.Errorf("Unknown paramter type. got=%v", p)
	}
}