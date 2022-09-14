package rlp

import (
	"testing"
)

func TestDecode(t *testing.T) {
	// 0x7f
	src1, dst1 := int16(127), int16(0)
	Decode(Encode(src1), &dst1)
	if src1 != dst1 {
		t.Errorf("Rlp decoding error. Expecter %d got %d\n", src1, dst1)
	}

	src2, dst2 := int32(127), int32(0)
	Decode(Encode(src2), &dst2)
	if src2 != dst2 {
		t.Errorf("Rlp decoding error. Expecter %d got %d\n", src2, dst2)
	}

	src3, dst3 := int64(127), int64(0)
	Decode(Encode(src3), &dst3)
	if src3 != dst3 {
		t.Errorf("Rlp decoding error. Expecter %d got %d\n", src3, dst3)
	}

	src4, dst4 := int8(127), int8(0)
	Decode(Encode(src4), &dst4)
	if src4 != dst4 {
		t.Errorf("Rlp decoding error. Expecter %d got %d\n", src4, dst4)
	}
}
