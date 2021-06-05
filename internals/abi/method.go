package abi

type Method struct {
	Name      string
	Arguments Arguments
	Output    Argument
}

// Signature returns function's signature according to the ABI spec.
//
// Example
// function foo(int64 a, int64 b) = "foo(int64,int64)"
func (method Method) Signature() string {
	return method.Name + "(" + method.Arguments.Pack() + ")"
}

// ID return function's id using function selector
func (method Method) ID() []byte {
	return Selector(method.Name + "(" + method.Arguments.Pack() + ")")
}