package leveldb

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	db, err := Open(os.Getenv("ETH_CHAIN_PATH"))
	if err != nil {
		t.Error(err)
	}

	blockKey := []byte{98, 0, 0, 0, 0, 0, 223, 251, 28, 91, 115, 205, 108, 168, 21, 65, 30, 170, 239, 84, 55, 90, 43, 148, 193, 150, 161, 79, 119, 129, 187, 112, 150, 20, 47, 13, 160, 136, 89, 98, 86}
	v, _ := db.Get(blockKey)

	if len(v) == 0 {
		t.Errorf("levelDB error. Expected data got nothing")
	}
}
