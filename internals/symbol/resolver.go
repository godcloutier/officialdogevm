package symbol

import (
	"../ast"
)

// Resolver traverse AST and resolve symbols, (1) check symbol's scope
// (2) check type of symbol
type Resolver struct {
	scope *Scope

	// types manage expression its own type
	types map[ast.Expression]SymbolType

	// defs manage identifier its own object
	defs map[*ast.Identifier]Symbol
}

func NewResolver() *Resolver {
	return &Resolver{
		scope: NewScope(),
		types: make(map[ast.Expression]SymbolType),
		defs:  make(map[*ast.Identifier]Symbol),
	}
}

// typeof returns the type of expression exp,
// or return InvalidSymbol if not found
func (r *Resolver) typeOf(exp ast.Expression) SymbolType {
	return ""
}

// objectOf returns the object denoted by the specified identifier
// or return nil if not found
func (r *Resolver) objectOf(id *ast.Identifier) Symbol {
	return nil
}

func (r *Resolver) ResolveContract(c *ast.Contract) error {
	return nil
}