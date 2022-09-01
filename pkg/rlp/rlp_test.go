package rlp

import (
	"bytes"
	"math/big"
	"testing"

	orig "github.com/ethereum/go-ethereum/rlp"
)

type test struct {
	val interface{}
}

type namedByteType byte

var (
	veryBigInt     = new(big.Int).Add(new(big.Int).Lsh(big.NewInt(0xFFFFFFFFFFFFFF), 16), big.NewInt(0xFFFF))
	veryVeryBigInt = new(big.Int).Exp(veryBigInt, big.NewInt(8), nil)
)

var tests = []test{
	// booleans
	{val: true},
	{val: false},

	// integers
	{val: uint32(0)},
	{val: uint32(127)},
	{val: uint32(128)},
	{val: uint32(256)},
	{val: uint32(1024)},
	{val: uint32(0xFFFFFF)},
	{val: uint32(0xFFFFFFFF)},
	{val: uint64(0xFFFFFFFF)},
	{val: uint64(0xFFFFFFFFFF)},
	{val: uint64(0xFFFFFFFFFFFF)},
	{val: uint64(0xFFFFFFFFFFFFFF)},
	{val: uint64(0xFFFFFFFFFFFFFFFF)},

	// big integers (should match uint for small values)
	{val: big.NewInt(0)},
	{val: big.NewInt(1)},
	{val: big.NewInt(127)},
	{val: big.NewInt(128)},
	{val: big.NewInt(256)},
	{val: big.NewInt(1024)},
	{val: big.NewInt(0xFFFFFF)},
	{val: big.NewInt(0xFFFFFFFF)},
	{val: big.NewInt(0xFFFFFFFFFF)},
	{val: big.NewInt(0xFFFFFFFFFFFF)},
	{val: big.NewInt(0xFFFFFFFFFFFFFF)},

	{val: veryBigInt},
	{val: veryVeryBigInt},

	// non-pointer big.Int
	{val: *big.NewInt(0)},
	{val: *big.NewInt(0xFFFFFF)},

	// byte arrays
	{val: [0]byte{}},
	{val: [1]byte{0}},
	{val: [1]byte{1}},
	{val: [1]byte{0x7F}},
	{val: [1]byte{0x80}},
	{val: [1]byte{0xFF}},
	{val: [3]byte{1, 2, 3}},
	{val: [57]byte{1, 2, 3}},

	// byte slices
	{val: []byte{}},
	{val: []byte{0}},
	{val: []byte{0x7E}},
	{val: []byte{0x7F}},
	{val: []byte{0x80}},
	{val: []byte{1, 2, 3}},

	// strings
	{val: ""},
	{val: "\x7E"},
	{val: "\x7F"},
	{val: "\x80"},
	{val: "dog"},
	{val: "Lorem ipsum dolor sit amet, consectetur adipisicing eli"},
	{val: "Lorem ipsum dolor sit amet, consectetur adipisicing elit"},
	{val: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur mauris magna, suscipit sed vehicula non, iaculis faucibus tortor. Proin suscipit ultricies malesuada. Duis tortor elit, dictum quis tristique eu, ultrices at risus. Morbi a est imperdiet mi ullamcorper aliquet suscipit nec lorem. Aenean quis leo mollis, vulputate elit varius, consequat enim. Nulla ultrices turpis justo, et posuere urna consectetur nec. Proin non convallis metus. Donec tempor ipsum in mauris congue sollicitudin. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Suspendisse convallis sem vel massa faucibus, eget lacinia lacus tempor. Nulla quis ultricies purus. Proin auctor rhoncus nibh condimentum mollis. Aliquam consequat enim at metus luctus, a eleifend purus egestas. Curabitur at nibh metus. Nam bibendum, neque at auctor tristique, lorem libero aliquet arcu, non interdum tellus lectus sit amet eros. Cras rhoncus, metus ac ornare cursus, dolor justo ultrices metus, at ullamcorper volutpat"},

	// named byte type arrays
	{val: [0]namedByteType{}},
	{val: [1]namedByteType{0}},
	{val: [1]namedByteType{1}},
	{val: [1]namedByteType{0x7F}},
	{val: [1]namedByteType{0x80}},
	{val: [1]namedByteType{0xFF}},
	{val: [3]namedByteType{1, 2, 3}},
	{val: [57]namedByteType{1, 2, 3}},

	// named byte type slices
	{val: []namedByteType{}},
	{val: []namedByteType{0}},
	{val: []namedByteType{0x7E}},
	{val: []namedByteType{0x7F}},
	{val: []namedByteType{0x80}},
	{val: []namedByteType{1, 2, 3}},

	// slices
	{val: []uint{}},
	{val: []uint{1, 2, 3}},
	{val: []interface{}{[]interface{}{}, [][]interface{}{{}}, []interface{}{[]interface{}{}, [][]interface{}{{}}}}},
	{val: []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii", "jjj", "kkk", "lll", "mmm", "nnn", "ooo"}},
	{val: []interface{}{uint(1), uint(0xFFFFFF), []interface{}{[]uint{4, 5, 5}}, "abc"}},
	{
		val: [][]string{
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
			{"asdf", "qwer", "zxcv"},
		},
	},

	// nil
	{val: (*uint)(nil)},
	{val: (*string)(nil)},
	{val: (*[]byte)(nil)},
	{val: (*[10]byte)(nil)},
	{val: (*big.Int)(nil)},
	{val: (*[]string)(nil)},
	{val: (*[10]string)(nil)},
	{val: (*[]interface{})(nil)},
	{val: (*[]struct{ uint })(nil)},
	{val: (*[]interface{})(nil)},
}

func TestEncode(t *testing.T) {
	for i, example := range tests {
		original, _ := orig.EncodeToBytes(example.val)
		tiny := Encode(example.val)

		if !bytes.Equal(tiny, original) {
			t.Errorf("\ntest %d: output mismatch:\ngot   %X\nwant  %X\n", i, tiny, original)
		}
	}
}
