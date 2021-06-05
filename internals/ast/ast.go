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
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Node represent ast node
type Node interface {
	String() string
}

// Represent Statement
type Statement interface {
	Node
	do()
}

// Represent Expression
type Expression interface {
	Node
	produce()
}

// Represent Contract.
// Contract consists of multiple functions.
type Contract struct {
	Functions []*FunctionLiteral
}

func (c *Contract) do() {}
func (c *Contract) String() string {
	var buf bytes.Buffer

	// start by change line for readability
	buf.WriteString("\ncontract {\n")

	for _, fn := range c.Functions {
		buf.WriteString(fn.String() + "\n")
	}
	buf.WriteString("}")

	return buf.String()
}

// Represent identifier
type Identifier struct {
	Name string
}

func (i *Identifier) String() string {
	return i.Name
}

func (i *Identifier) produce() {}

// Operator represent operator between expression
type Operator int

const (
	_        Operator = iota
	Plus              // +
	Minus             // -
	Bang              // !
	Asterisk          // *
	Slash             // /
	Mod               // %
	LT                // <
	GT                // >
	LTE               // <=
	GTE               // >=
	EQ                // ==
	NOT_EQ            // !=
	LAND              // &&
	LOR               // ||
)

var OperatorMap = map[Operator]string{
	Plus:     "+",
	Minus:    "-",
	Bang:     "!",
	Asterisk: "*",
	Slash:    "/",
	Mod:      "%",
	LT:       "<",
	GT:       ">",
	LTE:      "<=",
	GTE:      ">=",
	EQ:       "==",
	NOT_EQ:   "!=",
	LAND:     "&&",
	LOR:      "||",
}

func (o Operator) String() string {
	return OperatorMap[o]
}

// DataStructure represent identifier's data structure
// e.g. string, int, bool
type DataStructure int

const (
	_ DataStructure = iota
	IntType
	StringType
	BoolType
	VoidType
)

var DataStructureMap = map[DataStructure]string{
	IntType:    "int",
	StringType: "string",
	BoolType:   "bool",
	VoidType:   "void",
}

func (ds DataStructure) String() string {
	return DataStructureMap[ds]
}

// Represent assign statement
type AssignStatement struct {
	Type     DataStructure
	Variable Identifier
	Value    Expression
}

func (a *AssignStatement) do() {}

func (a *AssignStatement) String() string {
	var out bytes.Buffer
	out.WriteString(a.Type.String() + " ")
	out.WriteString(a.Variable.Name + " = ")
	out.WriteString(a.Value.String())
	return out.String()
}

// ReassignStatement is used when we want re-assign value to variable
type ReassignStatement struct {
	Variable *Identifier
	Value    Expression
}

func (r *ReassignStatement) do() {}
func (r *ReassignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.Variable.String() + " = ")
	out.WriteString(r.Value.String())

	return out.String()
}

// Represent return statement
type ReturnStatement struct {
	ReturnValue Expression
}

func (r *ReturnStatement) do() {}

func (r *ReturnStatement) String() string {
	if r.ReturnValue == nil {
		return "return"
	}
	return fmt.Sprintf("return %s", r.ReturnValue.String())
}

// Represent if statement
type IfStatement struct {
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfStatement) do() {}

func (i *IfStatement) String() string {
	if i.Alternative == nil {
		return fmt.Sprintf("if ( %s ) { %s }", i.Condition.String(), i.Consequence.String())
	}
	return fmt.Sprintf("if ( %s ) { %s } else { %s }", i.Condition.String(), i.Consequence.String(),
		i.Alternative.String())
}

// FunctionLiteral represents function definition
// e.g. func foo(int a) { ... }
type FunctionLiteral struct {
	Name       *Identifier
	Parameters []*ParameterLiteral
	Body       *BlockStatement
	ReturnType DataStructure
}

func (f *FunctionLiteral) do() {}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("func " + f.Name.String() + "(")

	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.ReturnType.String() + " {\n")
	out.WriteString(f.Body.String() + "\n")
	out.WriteString("}")

	return out.String()
}

func (f *FunctionLiteral) Signature() string {

	paramTypes := []string{}
	for _, p := range f.Parameters {
		paramTypes = append(paramTypes, p.Type.String())
	}

	return fmt.Sprintf("%s(%s)", f.Name.String(), strings.Join(paramTypes, ","))
}

// Represent block statement
type BlockStatement struct {
	Statements []Statement
}

func (b *BlockStatement) do() {}

func (b *BlockStatement) String() string {
	if b.Statements == nil {
		return ""
	}

	str := make([]string, 0)
	for _, s := range b.Statements {
		str = append(str, s.String())
	}
	return strings.Join(str, "\n")
}

// Represent function statement
type ExpressionStatement struct {
	Expr Expression
}

func (e *ExpressionStatement) do() {}

func (e *ExpressionStatement) String() string {
	return e.Expr.String()
}

// Represent string literal
type StringLiteral struct {
	Value string
}

func (s *StringLiteral) produce() {}

func (s *StringLiteral) String() string {
	return s.Value
}

// Represent integer literal
type IntegerLiteral struct {
	Value int64
}

func (i *IntegerLiteral) produce() {}

func (i *IntegerLiteral) String() string {
	return strconv.FormatInt(i.Value, 10)
}

// Represent Boolean expression
type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) produce() {}

func (b *BooleanLiteral) String() string {
	return strconv.FormatBool(b.Value)
}

// Represent Function Parameter expression
type ParameterLiteral struct {
	Identifier *Identifier
	Type       DataStructure
}

func (p *ParameterLiteral) produce() {}

func (p *ParameterLiteral) String() string {
	return fmt.Sprintf("Parameter : (Identifier: %s, Type: %s)", p.Identifier.String(), p.Type.String())
}

// Represent prefix expression
type PrefixExpression struct {
	Operator
	Right Expression
}

func (p *PrefixExpression) produce() {}

func (p *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator.String(), p.Right.String())
}

// Repersent Infix expression
type InfixExpression struct {
	Left Expression
	Operator
	Right Expression
}

func (i *InfixExpression) produce() {}

func (i *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left.String(), i.Operator.String(), i.Right.String())
}

// Represent Call expression
type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) produce() {}

func (c *CallExpression) String() string {
	strs := make([]string, 0)
	for _, exps := range c.Arguments {
		strs = append(strs, exps.String())
	}
	return fmt.Sprintf("function %s( %s )", c.Function.String(), strings.Join(strs, ", "))
}
