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

// Transaction is a raw Bitcoin transaction.
type Transaction struct {
	Hex      string `json:"hex"`
	Hash     string `json:"txid"`
	Size     int64  `json:"size"`
	Vsize    int64  `json:"vsize"`
	Version  int32  `json:"version"`
	Locktime int64  `json:"locktime"`
	Weight   int64  `json:"weight"`

	Inputs  []*Input  `json:"vin"`
	Outputs []*Output `json:"vout"`
}

// Metadata returns the metadata for a transaction.
func (t Transaction) Metadata() (map[string]interface{}, error) {
	m := &TransactionMetadata{
		Size:     t.Size,
		Vsize:    t.Vsize,
		Version:  t.Version,
		Locktime: t.Locktime,
		Weight:   t.Weight,
	}

	return types.MarshalMap(m)
}

// TransactionMetadata is a collection of useful
// metadata in a transaction.
type TransactionMetadata struct {
	Size     int64 `json:"size,omitempty"`
	Vsize    int64 `json:"vsize,omitempty"`
	Version  int32 `json:"version,omitempty"`
	Locktime int64 `json:"locktime,omitempty"`
	Weight   int64 `json:"weight,omitempty"`
}

// Input is a raw input in a Bitcoin transaction.
type Input struct {
	TxHash      string     `json:"txid"`
	Vout        int64      `json:"vout"`
	ScriptSig   *ScriptSig `json:"scriptSig"`
	Sequence    int64      `json:"sequence"`
	TxInWitness []string   `json:"txinwitness"`

	// Relevant when the input is the coinbase input
	Coinbase string `json:"coinbase"`
}

// Metadata returns the metadata for an input.
func (i Input) Metadata() (map[string]interface{}, error) {
	m := &OperationMetadata{
		ScriptSig:   i.ScriptSig,
		Sequence:    i.Sequence,
		TxInWitness: i.TxInWitness,
		Coinbase:    i.Coinbase,
	}

	return types.MarshalMap(m)
}

// Output is a raw output in a Bitcoin transaction.
type Output struct {
	Value        float64       `json:"value"`
	Index        int64         `json:"n"`
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey"`
}

// Metadata returns the metadata for an output.
func (o Output) Metadata() (map[string]interface{}, error) {
	m := &OperationMetadata{
		ScriptPubKey: o.ScriptPubKey,
	}

	return types.MarshalMap(m)
}
