package dogewire

import "encoding/json"

// BlockJSON is a raw Bitcoin block (with verbosity == 1).
type BlockJSON struct {
	Hash              string  `json:"hash"`
	Height            int64   `json:"height"`
	PreviousBlockHash string  `json:"previousblockhash"`
	Time              int64   `json:"time"`
	MedianTime        int64   `json:"mediantime"`
	Nonce             int64   `json:"nonce"`
	MerkleRoot        string  `json:"merkleroot"`
	Version           int32   `json:"version"`
	Size              int64   `json:"size"`
	Weight            int64   `json:"weight"`
	Bits              string  `json:"bits"`
	Difficulty        float64 `json:"difficulty"`

	Txs []interface{} `json:"tx"`
}

func (b *BlockJSON) UnmarshalJSON(data []byte) error {

	var res BlockJSON
	if err := json.Unmarshal(data, &res); err != nil {
		return nil
	}

	b.Hash = res.Hash
	b.Height = res.Height
	b.PreviousBlockHash = res.PreviousBlockHash
	b.Time = res.Time
	b.MedianTime = res.MedianTime
	b.Nonce = res.Nonce
	b.MerkleRoot = res.MerkleRoot
	b.Version = res.Version
	b.Size = res.Size
	b.Weight = res.Weight
	b.Bits = res.Bits
	b.Difficulty = res.Difficulty

	return nil
}
