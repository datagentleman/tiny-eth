package block

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/datagentleman/tiny-eth/pkg/common"
	"github.com/datagentleman/tiny-eth/pkg/config"
	"github.com/datagentleman/tiny-eth/pkg/db"
)

func TestFindHeader(t *testing.T) {
	conf, _ := config.Get("database", "test")
	db.Configure(conf)

	headerHash := []byte{35, 32, 71, 217, 213, 95, 254, 81, 128, 185, 134, 176, 57, 191, 244, 155, 28, 152, 253, 196, 86, 253, 50, 46, 210, 205, 255, 116, 94, 211, 232, 226}
	h, err := FindHeader(common.NewHash(headerHash))
	if err != nil {
		fmt.Println(err)
	}

	h1 := []byte{115, 75, 181, 67, 250, 79, 27, 103, 109, 14, 125, 22, 39, 201, 197, 110, 147, 69, 82, 77, 112, 56, 18, 23, 189, 253, 52, 123, 9, 91, 128, 95}
	if !bytes.Equal(h.ParentHash.Bytes(), h1) {
		t.Errorf("wrong hash, expected %x, got %x", h1, h.ParentHash)
	}

	v1 := big.NewInt(14595348028527168)
	res := h.Difficulty.Cmp(v1)
	if res != 0 {
		t.Errorf("wrong difficulty, expected %d, got %d", v1, h.Difficulty)
	}

	v1 = big.NewInt(14768770)
	res = h.Number.Cmp(v1)
	if res != 0 {
		t.Errorf("wrong number, expected %, got %d", v1, h.Number)
	}
}
