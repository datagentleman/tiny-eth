package rlp

import (
	"bytes"
	"reflect"
	"testing"
)

type decodeTest struct {
	src interface{}
	dst interface{}
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

var decodeBytes = []decodeTest{
	{src: make([]byte, 55), dst: []byte{}},
	{src: make([]byte, 1024), dst: []byte{}},
}

var decodeSlice = []decodeTest{
	// List
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

func TestDecode(t *testing.T) {
	for _, example := range decodeTests {
		Decode(Encode(example.src), &example.dst)
		if example.src != example.dst {
			t.Errorf("Rlp decoding error. Expected\n %d got\n %d\n", example.src, example.dst)
		}
	}

	for _, example := range decodeBytes {
		Decode(Encode(example.src), &example.dst)
		if !bytes.Equal(example.src.([]byte), example.dst.([]byte)) {
			t.Errorf("Rlp decoding error. Expected %d got %d\n", example.src, example.dst)
		}
	}

	for _, example := range decodeSlice {
		Decode(Encode(example.src), &example.dst)
		if !(reflect.DeepEqual(example.src, example.dst)) {
			t.Errorf("Rlp decoding error. Expected %d got %d\n", example.src, example.dst)
		}
	}
}
