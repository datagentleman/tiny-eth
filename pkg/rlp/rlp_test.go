package rlp

import (
	"fmt"
	"testing"

	orig "github.com/ethereum/go-ethereum/rlp"
)

func TestEncode(t *testing.T) {
	v := []uint16{200, 300, 0}

	fmt.Println(orig.EncodeToBytes(v))
	fmt.Println(Encode(v))
}
