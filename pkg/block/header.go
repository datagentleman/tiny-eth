package block

import (
	"encoding/binary"
	"math/big"

	"github.com/datagentleman/tiny-eth/pkg/common"
	"github.com/datagentleman/tiny-eth/pkg/db"
	"github.com/datagentleman/tiny-eth/pkg/rlp"
)

type Header struct {
	ParentHash  common.Hash
	UncleHash   common.Hash
	Coinbase    common.Address
	Root        common.Hash
	TxHash      common.Hash
	ReceiptHash common.Hash
	Bloom       common.Bloom
	Difficulty  *big.Int
	Number      *big.Int
	GasLimit    uint64
	GasUsed     uint64
	Time        uint64
	Extra       []byte
	MixDigest   common.Hash
	Nonce       common.BlockNonce
}

func FindHeader(hash *common.Hash) (*Header, error) {
	headerNumberPrefix := []byte("H")
	headerNumberKey := append(headerNumberPrefix, hash[:]...)

	data, err := db.Get(headerNumberKey)
	if err != nil {
		return nil, err
	}

	number := binary.BigEndian.Uint64(data)

	headerPrefix := []byte("h")
	headerKey := append(append(headerPrefix, common.NumberToBytes(number)...), hash[:]...)

	data, _ = db.Get(headerKey)
	if err != nil {
		return nil, err
	}

	h := &Header{}
	h.Decode(data)

	return h, nil
}

func (h *Header) Decode(data []byte) {
	dec := rlp.NewDecoder(data)

	h.Number = new(big.Int)
	h.Difficulty = new(big.Int)

	l := []interface{}{
		&h.ParentHash,
		&h.UncleHash,
		&h.Coinbase,
		&h.Root,
		&h.TxHash,
		&h.ReceiptHash,
		&h.Bloom,
		h.Difficulty,
		h.Number,
		&h.GasLimit,
		&h.GasUsed,
		&h.Time,
		&h.Extra,
		&h.MixDigest,
		&h.Nonce,
	}

	dec.Decode(&l)
}
