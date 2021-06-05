/*
 *	Copyright 2018-2019 De-labtory
 *	Copyright 2021 Dogecoin Smart Contract Developers
* Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.

*/
package opcode_test

import (
	"testing"

	"../../internals/opcode"
)

func TestType_String(t *testing.T) {
	tests := []struct {
		input    opcode.Type
		expected string
	}{
		{
			opcode.Add,
			"Add",
		},
		{
			opcode.Mul,
			"Mul",
		},
		{
			opcode.Sub,
			"Sub",
		},
		{
			opcode.Div,
			"Div",
		},
		{
			opcode.Mod,
			"Mod",
		},
		{
			opcode.And,
			"And",
		},
		{
			opcode.Or,
			"Or",
		},
		{
			opcode.LT,
			"LT",
		},
		{
			opcode.LTE,
			"LTE",
		},
		{
			opcode.GT,
			"GT",
		},
		{
			opcode.GTE,
			"GTE",
		},
		{
			opcode.EQ,
			"EQ",
		},
		{
			opcode.NOT,
			"NOT",
		},
		{
			opcode.Minus,
			"Minus",
		},
		{
			opcode.Pop,
			"Pop",
		},
		{
			opcode.Push,
			"Push",
		},
		{
			opcode.Mload,
			"Mload",
		},
		{
			opcode.Mstore,
			"Mstore",
		},
		{
			opcode.Msize,
			"Msize",
		},
		{
			opcode.LoadFunc,
			"LoadFunc",
		},
		{
			opcode.LoadArgs,
			"LoadArgs",
		},
		{
			opcode.Returning,
			"Returning",
		},
		{
			opcode.Jump,
			"Jump",
		},
		{
			opcode.JumpDst,
			"JumpDst",
		},
		{
			opcode.Jumpi,
			"Jumpi",
		},
		{
			opcode.DUP,
			"DUP",
		},
		{
			opcode.SWAP,
			"SWAP",
		},
		{
			opcode.Exit,
			"Exit",
		},
		{
			0x97,
			"String() error - Not defined opcode",
		},
	}

	for i, test := range tests {
		str, err := test.input.String()
		if str != "" && str != test.expected {
			t.Fatalf("test[%d] - TestType_String() wrong result.\n"+
				"expected: %s\n"+
				"got: %s", i, test.expected, str)
		}
		if err != nil && err.Error() != test.expected {
			t.Fatalf("test[%d] - TestType_String() wrong error.\n"+
				"expected: %s\n"+
				"got: %s", i, test.expected, err.Error())
		}
	}
}
