package rlp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
)

type decodeTest struct {
	src interface{}
	dst interface{}
}

type TestCustomType []byte

type TestCustomClass struct {
	A1   string
	A2   *string
	B1   uint
	B2   *uint
	C1   uint16
	C2   *uint16
	D1   []byte
	D2   *[]byte
	CT1  TestCustomType
	CT11 *TestCustomType
}

var decodeTests = []decodeTest{
	// Strings
	{src: int8(127), dst: int8(0)},
	{src: int8(-127), dst: int8(0)},

	{src: int16(32767), dst: int16(0)},
	{src: int16(32767), dst: int16(0)},

	{src: int32(2147483647), dst: int32(0)},
	{src: int32(-2147483647), dst: int32(0)},

	{src: int64(9223372036854775807), dst: int64(0)},
	{src: int64(-9223372036854775807), dst: int64(0)},

	{src: bool(false), dst: bool(true)},
	{src: bool(true), dst: bool(false)},

	{src: uint8(255), dst: uint8(0)},
	{src: uint16(65535), dst: uint16(0)},
	{src: uint32(4294967295), dst: uint32(0)},
	{src: uint64(18446744073709551615), dst: uint64(0)},

	{src: float32(3.4e+38), dst: float32(0)},
	{src: float32(1.2e-38), dst: float32(0)},

	{src: float64(1.7e+308), dst: float64(0)},
	{src: float64(2.2e-308), dst: float64(0)},

	{src: string("Lorem ipsum dolor sit amet"), dst: string("")},
	{src: string("Lorem ipsum dolor sit amet Lorem ipsum dolor sit amet Lorem ipsum dolor sit amet"), dst: string("")},
}

var decodeSlice = []decodeTest{
	// Lists
	{src: make([]byte, 55), dst: []byte{}},
	{src: make([]byte, 1024), dst: []byte{}},

	{src: []int8{127, -127}, dst: []int8{}},
	{src: []int16{32767, -32767}, dst: []int16{}},
	{src: []int32{2147483647, -2147483647}, dst: []int32{}},
	{src: []int64{9223372036854775807, -9223372036854775807}, dst: []int64{}},

	{src: []uint16{65535, 65535}, dst: []uint16{}},
	{src: []uint32{4294967295, 4294967295}, dst: []uint32{}},
	{src: []uint64{18446744073709551615, 18446744073709551615}, dst: []uint64{}},

	{src: []float32{3.4e+38, 1.2e-38}, dst: []float32{}},
	{src: []float64{1.7e+308, 2.2e-308}, dst: []float64{}},
}

func TestDecodeStrings(t *testing.T) {
	for _, example := range decodeTests {
		dec := NewDecoder(Encode(example.src))
		dec.Decode(&example.dst)
		if example.src != example.dst {
			t.Errorf("Rlp decoding error. Expected\n %d got\n %d\n", example.src, example.dst)
		}
	}
}

func TestDecodeLists(t *testing.T) {
	for _, example := range decodeSlice {
		dec := NewDecoder(Encode(example.src))
		dec.Decode(&example.dst)
		if !(reflect.DeepEqual(example.src, example.dst)) {
			t.Errorf("Rlp decoding error. Expected %d got %d\n", example.src, example.dst)
		}
	}
}

func TestDecodeStructs(t *testing.T) {
	s := string("Lorem ipsum dolor sit amet")
	u := uint(127)
	u16 := uint16(127)
	b := []byte{1, 2, 3}
	ct := TestCustomType{10, 20, 30}

	src := TestCustomClass{
		A1:   s,
		A2:   &s,
		B1:   u,
		B2:   &u,
		C1:   u16,
		C2:   &u16,
		D1:   b,
		D2:   &b,
		CT1:  ct,
		CT11: &ct,
	}

	dst := TestCustomClass{}

	encoding, err := rlp.EncodeToBytes(&src)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	dec := NewDecoder(encoding)
	dec.Decode(&dst)

	if !(reflect.DeepEqual(src, dst)) {
		t.Errorf("rlp struct decoding error.\n Expected: %v\n Received: %v", src, dst)
	}
}
