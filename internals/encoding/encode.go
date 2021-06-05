package encoding

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"


	"github.com/ethereum/go-ethereum/common/math"
)

// In koa, we use hexadecimal encoding
type EncodeError struct {
	Operand interface{}
}

func (e EncodeError) Error() string {
	return fmt.Sprintf("EncodeOperand() error - operand %v could not encoded", e.Operand)
}

// EncodeOperand() encodes operand to bytes.
func EncodeOperand(operand interface{}) ([]byte, error) {
	switch op := operand.(type) {
	case int:
		return encodeInt(int64(op))

	case int64:
		return encodeInt(op)

	case string:
		return encodeString(op)

	case bool:
		return encodeBool(op)

	case []byte:
		return encodeBytes(op)

	default:
		return nil, EncodeError{op}
	}
}

// Encode integer to hexadecimal bytes
// ex) int 123 => 0x000000000000007b
func encodeInt(operand int64) ([]byte, error) {
	byteSlice := make([]byte, 8)
	binary.BigEndian.PutUint64(byteSlice, uint64(operand))

	return byteSlice, nil
}

// Encode string to hexadecimal bytes
// ex) string "abc" => 0x6162630000000000
func encodeString(operand string) ([]byte, error) {
	if len(operand) > 8 {
		return nil, errors.New("Length of string must shorter than 8")
	}

	src := hex.EncodeToString([]byte(operand))
	if len(src)&1 == 1 {
		src = "0" + src
	}

	dst, err := hex.DecodeString(src)
	if err != nil {
		return nil, err
	}

	for len(dst) < 8 {
		dst = append(dst, 0)
	}

	return dst, nil
}

// Encode boolean to hexadecimal bytes
// ex) bool true  => 0x0000000000000001
// ex) bool false => 0x0000000000000000
func encodeBool(operand bool) ([]byte, error) {
	byteSlice := make([]byte, 8)

	if operand {
		binary.BigEndian.PutUint64(byteSlice, 1)
	} else {
		binary.BigEndian.PutUint64(byteSlice, 0)
	}

	return byteSlice, nil
}

// Encode boolean to hexadecimal bytes
// ex) []byte{0x01, 0x02, 0x03, 0x04}  => 0x0000000001020304
func encodeBytes(operand []byte) ([]byte, error) {
	copiedBytes := make([]byte, 8)

	tmp := new(big.Int)
	tmp.SetBytes(operand)

	math.ReadBits(tmp, copiedBytes)

	return copiedBytes, nil
}