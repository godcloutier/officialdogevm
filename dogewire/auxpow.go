package dogewire

import (
	"encoding/binary"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"io"
)

const ERR_INVALID_HEADER = 10050
const ERR_COINBASE_INDEX = 10060 // coinbase tx index within Litecoin merkle isn't 0
const ERR_NOT_MERGE_MINED = 10070 // trying to check AuxPoW on a block that wasn't merge mined
const ERR_FOUND_TWICE = 10080 // 0xfabe6d6d found twice
const ERR_NO_MERGE_HEADER = 10090 // 0xfabe6d6d not found
const ERR_NOT_IN_FIRST_20 = 10100 // chain Merkle root isn't in the first 20 bytes of coinbase tx
const ERR_CHAIN_MERKLE = 10110
const ERR_PARENT_MERKLE = 10120
const ERR_PROOF_OF_WORK = 10130

const (
	NETWORK_MAIN = "MAINNET"
	NETWORK_TEST = "TESTNET"
	NETWORK_REG = "REGNET"
)

// Copyright 2021 Dogecoin Rosetta Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

const vAuxPoW = 0x0100


type MerkleBranch struct {
	BranchHashes   []*chainhash.Hash
	BranchSideMask int32
}

type AuxHeader struct {
	CoinbaseTx wire.MsgTx
	BlockHash chainhash.Hash
	CoinbaseBranch MerkleBranch
	BlockchainBranch MerkleBranch
	ParentHeader wire.BlockHeader

}

type AuxBlock struct {
	Header       wire.BlockHeader
	AuxPoW       AuxHeader
	Transactions []*wire.MsgTx
}


// Deserialize decodes a merkel branch from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.
func (mb *MerkleBranch) Deserialize(r io.Reader) error {
	branchCount, err := wire.ReadVarInt(r, 0)
	if err != nil {
		return err
	}

	mb.BranchHashes = make([]*chainhash.Hash, 0, branchCount)
	for i := uint64(0); i < branchCount; i++ {
		hash := chainhash.Hash{}
		if _, err := io.ReadFull(r, hash[:]); err != nil {
			return err
		}
		mb.BranchHashes = append(mb.BranchHashes, &hash)
	}

	if err := binary.Read(r, binary.LittleEndian, &mb.BranchSideMask); err != nil {
		return err
	}

	return nil
}

// Deserialize decodes a AuxPow header from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.
func (aux *AuxHeader) Deserialize(r io.Reader) error {
	if err := aux.CoinbaseTx.Deserialize(r); err != nil {
		return err
	}
	if _, err := io.ReadFull(r, aux.BlockHash[:]); err != nil {
		return err
	}
	if err := aux.CoinbaseBranch.Deserialize(r); err != nil {
		return err
	}
	if err := aux.BlockchainBranch.Deserialize(r); err != nil {
		return err
	}
	if err := aux.ParentHeader.Deserialize(r); err != nil {
		return err
	}
	return nil
}

// Deserialize decodes a AuxPoW block from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.
func (b *AuxBlock) Deserialize(r io.Reader) error {
	if err := b.Header.Deserialize(r); err != nil {
		return err
	}

	if (b.Header.Version & vAuxPoW) != 0 {
		if err := b.AuxPoW.Deserialize(r); err != nil {
			return err
		}
	}

	txCount, err := wire.ReadVarInt(r, 0)
	if err != nil {
		return err
	}

	b.Transactions = make([]*wire.MsgTx, 0, txCount)
	for i := uint64(0); i < txCount; i++ {
		tx := wire.MsgTx{}
		if err := tx.Deserialize(r); err != nil {
			return err
		}
		b.Transactions = append(b.Transactions, &tx)
	}

	return nil
}




/*
	END OF CODE FROM ROSETTA API AUXPOW
*/
