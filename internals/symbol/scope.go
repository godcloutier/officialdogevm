package symbol

import (
	"bytes"
	"fmt"
	"sort"
)

// NewScope() makes a new scope.
func NewScope() *Scope {
	s := make(map[string]Symbol)
	return &Scope{
		store: s,
		inner: make([]*Scope, 0),
		outer: nil,
	}
}

// NewEnclosedScope makes a new scope which has outer scope.
func NewEnclosedScope(outer *Scope) *Scope {
	scope := NewScope()
	scope.outer = outer
	return scope
}

// Scope represent a variable's scope.
type Scope struct {
	store map[string]Symbol

	inner []*Scope
	outer *Scope
}

// Getter return's variable's scope.
// If a scope doesn't have a variable,
// Program will search in outer scope.
func (s *Scope) Get(name string) Symbol {
	var scope = s
	for scope != nil {
		obj, ok := scope.store[name]
		if ok {
			return obj
		}
		scope = scope.outer
	}
	return nil
}

// Setter set a variable to target scope's map
func (s *Scope) Set(name string, val Symbol) Symbol {
	s.store[name] = val
	return val
}

func (s *Scope) SetOuter(outer *Scope) {
	s.outer = outer
}

func (s *Scope) GetOuter() *Scope {
	return s.outer
}

func (s *Scope) AppendInner(in *Scope) {
	s.inner = append(s.inner, in)
}

func (s *Scope) GetInner() []*Scope {
	return s.inner
}

func (s *Scope) String() string {
	var out bytes.Buffer
	scope := s

	for scope != nil {
		// sort scope.store's keys alphabetically
		keys := make([]string, 0)
		for k := range scope.store {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		out.WriteString("[ Scope ]\n")
		for _, k := range keys {
			out.WriteString(fmt.Sprintf("key:%s, symbol:%s\n",
				k, scope.store[k].String()))
		}

		scope = scope.outer
	}
	return out.String()
}