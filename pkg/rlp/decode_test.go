package rlp

import (
	"bytes"
	"testing"
)

type decodeTest struct {
	src interface{}
	dst interface{}
}

var decodeTests = []decodeTest{
	// String

	// 0sx7f
	{src: int8(127), dst: int8(0)},
	{src: int16(127), dst: int16(0)},
	{src: int32(127), dst: int32(0)},
	{src: int64(127), dst: int64(0)},

	// 0xb7
	{src: int16(-127), dst: int16(0)},
	{src: int64(999999999999999999), dst: int64(0)},
}

// {src: string("Lorem ipsum dolor sit amet"), dst: string("")},

var decodeSlice = []decodeTest{
	// List

	// 0xf7
	{src: []int32{200, 10, 10}, dst: []int32{}},
}

var decodeBytes = []decodeTest{
	// 0xb7
	{src: make([]byte, 55), dst: []byte{}},
}

func TestDecode(t *testing.T) {
	for _, example := range decodeTests {
		Decode(Encode(example.src), &example.dst)
		if example.src != example.dst {
			t.Errorf("Rlp decoding error. Expected %d got %d\n", example.src, example.dst)
		}
	}

	for _, example := range decodeSlice {
		Decode(Encode(example.src), &example.dst)
		if len(example.src.([]int32)) != len(example.dst.([]int32)) {
			t.Errorf("Rlp decoding error. Expected %d got %d\n", example.src, example.dst)
		}
	}

	for _, example := range decodeBytes {
		Decode(Encode(example.src), &example.dst)
		if !bytes.Equal(example.src.([]byte), example.dst.([]byte)) {
			t.Errorf("Rlp decoding error. Expected %d got %d\n", example.src, example.dst)
		}
	}
}
