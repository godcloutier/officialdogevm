package dogewire

// OperationMetadata is a collection of useful
// metadata from Bitcoin inputs and outputs.
type OperationMetadata struct {
	// Coinbase Metadata
	Coinbase string `json:"coinbase,omitempty"`

	// Input Metadata
	ScriptSig   *ScriptSig `json:"scriptsig,omitempty"`
	Sequence    int64      `json:"sequence,omitempty"`
	TxInWitness []string   `json:"txinwitness,omitempty"`

	// Output Metadata
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey,omitempty"`
}

// request represents the JSON-RPC request body
type request struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func (r request) GetVersion() string       { return r.JSONRPC }
func (r request) GetID() int               { return r.ID }
func (r request) GetMethod() string        { return r.Method }
func (r request) GetParams() []interface{} { return r.Params }
