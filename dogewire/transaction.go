package dogewire

// ScriptPubKey is a script placed on the output operations
// of a Bitcoin transaction that must be satisfied to spend
// the output.
type ScriptPubKey struct {
	ASM          string   `json:"asm"`
	Hex          string   `json:"hex"`
	RequiredSigs int64    `json:"reqSigs,omitempty"`
	Type         string   `json:"type"`
	Addresses    []string `json:"addresses,omitempty"`
}

// ScriptSig is a script on the input operations of a
// Bitcoin transaction that satisfies the ScriptPubKey
// on an output being spent.
type ScriptSig struct {
	ASM string `json:"asm"`
	Hex string `json:"hex"`
}
