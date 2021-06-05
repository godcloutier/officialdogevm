package abi

import (

	"../cryptoutil"
	"../encoding"
)

type Pointer []byte
type Size []byte
type Value []byte

// Encode abi parameters
func Encode(params ...interface{}) ([]byte, error) {
	values, err := encodeValues(params...)
	if err != nil {
		return nil, err
	}

	sizes, err := encodeSizes(values)
	if err != nil {
		return nil, err
	}

	pointers, err := encodePointers(sizes, values)
	if err != nil {
		return nil, err
	}

	var Args []byte

	for _, pointer := range pointers {
		Args = append(Args, pointer...)
	}

	for index := 0; index < len(params); index++ {
		Args = append(Args, sizes[index]...)
		Args = append(Args, values[index]...)
	}

	return Args, nil
}

func encodeValues(params ...interface{}) ([]Value, error) {
	values := make([]Value, len(params))

	for index, param := range params {
		bytesValue, err := encoding.EncodeOperand(param)
		if err != nil {
			return nil, err
		}

		values[index] = append(values[index], bytesValue...)
	}

	return values, nil
}

func encodeSizes(values []Value) ([]Size, error) {
	sizes := make([]Size, len(values))

	for index, value := range values {
		size, err := encoding.EncodeOperand(len(value))
		if err != nil {
			return nil, err
		}

		sizes[index] = size
	}

	return sizes, nil
}

func encodePointers(sizes []Size, values []Value) ([]Pointer, error) {
	pointers := make([]Pointer, len(values))

	for index := 0; index < len(values); index++ {
		length := len(values) * 8

		for from := 0; from < index; from++ {
			length += len(sizes[from]) + len(values[from])
		}

		pointer, err := encoding.EncodeOperand(length)
		if err != nil {
			return nil, err
		}

		pointers[index] = pointer
	}


	return pointers, nil
}

// The implementation below is implemented for the abi spec of ethereum.
// https://solidity.readthedocs.io/en/develop/abi-spec.html

// Get function selector(4bytes) from string of function signature
func Selector(functionSignature string) []byte {
	return cryptoutil.Keccak256([]byte(functionSignature))[:4]
}