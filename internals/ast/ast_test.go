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

package ast

import (
	"testing"
)

func TestAssignStatement_String(t *testing.T) {
	tests := []struct {
		input    AssignStatement
		expected string
	}{
		{
			input: AssignStatement{
				Type:     IntType,
				Variable: Identifier{Name: "a"},
				Value:    &IntegerLiteral{Value: 1},
			},
			expected: "int a = 1",
		},
		{
			input: AssignStatement{
				Type:     BoolType,
				Variable: Identifier{Name: "asdf"},
				Value:    &StringLiteral{Value: "hello, world"},
			},
			// type mismatch is not considered here
			expected: "bool asdf = hello, world",
		},
		{
			input: AssignStatement{
				Type:     StringType,
				Variable: Identifier{Name: "ff"},
				Value:    &BooleanLiteral{Value: true},
			},
			// type mismatch is not considered here
			expected: "string ff = true",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestIdentifier_String(t *testing.T) {
	tests := []struct {
		input    Identifier
		expected string
	}{
		{
			input:    Identifier{Name: "a"},
			expected: "a",
		},
		{
			input:    Identifier{Name: "zcf"},
			expected: "zcf",
		},
		{
			input:    Identifier{Name: "zcf asdf"},
			expected: "zcf asdf",
		},
		{
			input:    Identifier{Name: "123"},
			expected: "123",
		},
		{
			input:    Identifier{Name: ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestStringLiteral_String(t *testing.T) {
	tests := []struct {
		input    StringLiteral
		expected string
	}{
		{
			StringLiteral{"hello"},
			"hello",
		},
		{
			StringLiteral{"hello, world"},
			"hello, world",
		},
		{
			StringLiteral{"123"},
			"123",
		},
		{
			StringLiteral{"123, hello"},
			"123, hello",
		},
		{
			StringLiteral{""},
			"",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestIntegerLiteral_String(t *testing.T) {
	tests := []struct {
		input    IntegerLiteral
		expected string
	}{
		{
			IntegerLiteral{Value: 123},
			"123",
		},
		{
			IntegerLiteral{Value: -1},
			"-1",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestBooleanLiteral_String(t *testing.T) {
	tests := []struct {
		input    BooleanLiteral
		expected string
	}{
		{
			BooleanLiteral{true},
			"true",
		},
		{
			BooleanLiteral{false},
			"false",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestFunctionLiteral_String(t *testing.T) {
	tests := []struct {
		input    FunctionLiteral
		expected string
	}{
		{
			FunctionLiteral{
				Name:       &Identifier{Name: "foo"},
				Parameters: []*ParameterLiteral{},
				Body:       &BlockStatement{},
				ReturnType: StringType,
			},
			`func foo() string {
}`,
		},
		{
			FunctionLiteral{
				Name:       &Identifier{Name: "foo"},
				Parameters: []*ParameterLiteral{},
				Body: &BlockStatement{
					Statements: []Statement{
						&AssignStatement{
							Type:     IntType,
							Variable: Identifier{Name: "a"},
							Value:    &IntegerLiteral{Value: 1},
						},
					},
				},
				ReturnType: IntType,
			},
			`func foo() int {
int a = 1
}`,
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestReassignStatement_String(t *testing.T) {
	tests := []struct {
		input    ReassignStatement
		expected string
	}{
		{
			input: ReassignStatement{
				Variable: &Identifier{
					Name: "foo",
				},
				Value: &BooleanLiteral{
					Value: true,
				},
			},
			expected: "foo = true",
		},
		{
			input: ReassignStatement{
				Variable: &Identifier{
					Name: "foo",
				},
				Value: &PrefixExpression{
					Operator: Minus,
					Right: &IntegerLiteral{
						Value: 1,
					},
				},
			},
			expected: "foo = (-1)",
		},
		{
			input: ReassignStatement{
				Variable: &Identifier{
					Name: "foo",
				},
				Value: &InfixExpression{
					Left: &StringLiteral{
						Value: "hello",
					},
					Operator: Plus,
					Right: &IntegerLiteral{
						Value: 2,
					},
				},
			},
			expected: "foo = (hello + 2)",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestReturnStatement_String(t *testing.T) {
	tests := []struct {
		input    ReturnStatement
		expected string
	}{
		{
			input: ReturnStatement{
				ReturnValue: &IntegerLiteral{Value: 1},
			},
			expected: "return 1",
		},
		{
			input: ReturnStatement{
				ReturnValue: &StringLiteral{Value: "\"hello, world\""},
			},
			expected: "return \"hello, world\"",
		},
		{
			input: ReturnStatement{
				ReturnValue: &StringLiteral{Value: "\"hello, world\""},
			},
			expected: "return \"hello, world\"",
		},
		{
			input: ReturnStatement{
				ReturnValue: &Identifier{Name: "a"},
			},
			expected: "return a",
		},
		{
			input:    ReturnStatement{},
			expected: "return",
		},
	}

	for _, tt := range tests {
		result := tt.input.String()
		testString(t, result, tt.expected)
	}
}

func TestFunctionLiteral_Signature(t *testing.T) {
	tests := []struct {
		input    FunctionLiteral
		expected string
	}{
		{
			input: FunctionLiteral{
				Name: &Identifier{Name: "foo"},
			},
			expected: "foo()",
		},
		{
			input: FunctionLiteral{
				Name: &Identifier{Name: "foo"},
				Parameters: []*ParameterLiteral{
					{
						Identifier: &Identifier{Name: "a"},
						Type:       IntType,
					},
				},
			},
			expected: "foo(int)",
		},
		{
			input: FunctionLiteral{
				Name: &Identifier{Name: "foo"},
				Parameters: []*ParameterLiteral{
					{
						Identifier: &Identifier{Name: "a"},
						Type:       IntType,
					},
					{
						Identifier: &Identifier{Name: "b"},
						Type:       StringType,
					},
				},
			},
			expected: "foo(int,string)",
		},
	}

	for _, tt := range tests {
		result := tt.input.Signature()
		testString(t, result, tt.expected)
	}
}

func testString(t *testing.T, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf(`String() wrong result. expected="%s", got="%s"`, expected, got)
	}
}
