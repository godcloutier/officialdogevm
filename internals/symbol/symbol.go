package symbol

import (
	"fmt"

	"../ast"
)

type SymbolType string

const (
	IntegerSymbol  = "INTEGER"
	BooleanSymbol  = "BOOLEAN"
	StringSymbol   = "STRING"
	FunctionSymbol = "FUNCTION"
)

type Symbol interface {
	Type() SymbolType
	String() string
}

// Represent Integer symbol
type Integer struct {
	Name *ast.Identifier
}

func (i *Integer) Type() SymbolType {
	return IntegerSymbol
}

// String() returns symbol's name
func (i *Integer) String() string {
	return fmt.Sprintf("%s", i.Name.String())
}

// Represent Boolean Object
type Boolean struct {
	Name *ast.Identifier
}

func (b *Boolean) Type() SymbolType {
	return BooleanSymbol
}

func (b *Boolean) String() string {
	return fmt.Sprintf("%s", b.Name.String())
}

// Represent String Object
type String struct {
	Name *ast.Identifier
}

func (s *String) Type() SymbolType {
	return StringSymbol
}

func (s *String) String() string {
	return fmt.Sprintf("%s", s.Name.String())
}

// Represent Function symbol
// Name represents function's name.
// Scope represents function value's scope.
type Function struct {
	Name  string
	Scope *Scope
}

func (f *Function) Type() SymbolType {
	return FunctionSymbol
}

func (f *Function) String() string {
	return fmt.Sprintf("%s", f.Name)
}