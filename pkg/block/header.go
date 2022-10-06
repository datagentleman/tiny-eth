package block

import (
	"math/big"
)

type Hash [32]byte

func (h *Hash) SetBytes(data []byte) {
	if len(h) == len(data) {
		_ = append(h[:0], data...)
	}
}

type Address [20]byte

func (a *Address) SetBytes(data []byte) {
	if len(a) == len(data) {
		_ = append(a[:0], data...)
	}
}

type Bloom [256]byte

func (b *Bloom) SetBytes(data []byte) {
	if len(b) == len(data) {
		_ = append(b[:0], data...)
	}
}

type BlockNonce [8]byte

func (n *BlockNonce) SetBytes(data []byte) {
	if len(n) == len(data) {
		_ = append(n[:0], data...)
	}
}

type Header struct {
	ParentHash  Hash
	UncleHash   Hash
	Coinbase    Address
	Root        Hash
	TxHash      Hash
	ReceiptHash Hash
	Bloom       Bloom
	Difficulty  *big.Int
	Number      *big.Int
	GasLimit    uint64
	GasUsed     uint64
	Time        uint64
	Extra       []byte
	MixDigest   Hash
	Nonce       BlockNonce
}
