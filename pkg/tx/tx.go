package tx

import (
	"math/big"
	"sync/atomic"
	"time"
)

const (
	HashLength    = 32
	AddressLength = 20
)

type AccessTuple struct {
	Address     Address
	StorageKeys []Hash
}

type AccessList []AccessTuple

type Hash [HashLength]byte
type Address [AddressLength]byte

type Tx struct {
	ChainID    *big.Int   // destination chain ID
	Nonce      uint64     // nonce of sender account
	GasTipCap  *big.Int   // maxPriorityFeePerGas
	GasFeeCap  *big.Int   // maxFeePerGas
	GasPrice   *big.Int   // wei per gas
	Gas        uint64     // gas limit
	To         *Address   // destination address
	Value      *big.Int   // wei amount
	Data       []byte     // contract invocation input data
	AccessList AccessList // EIP-2930 access list
	V, R, S    *big.Int   // signature values

	firstSeen time.Time // Time first seen locally
	hash      atomic.Value
	size      atomic.Value
	from      atomic.Value
}
