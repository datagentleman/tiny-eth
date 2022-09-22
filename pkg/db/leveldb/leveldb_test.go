package leveldb

import (
	"os"
	"testing"
)

var dbTest *Level

func TestMain(m *testing.M) {
	dbTest, _ = Open(os.Getenv("ETH_CHAIN_PATH"))

	code := m.Run()

	err := dbTest.Close()
	if err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestGet(t *testing.T) {
	blockKey := []byte{98, 0, 0, 0, 0, 0, 223, 251, 28, 91, 115, 205, 108, 168, 21, 65, 30, 170, 239, 84, 55, 90, 43, 148, 193, 150, 161, 79, 119, 129, 187, 112, 150, 20, 47, 13, 160, 136, 89, 98, 86}
	v, _ := dbTest.Get(blockKey)

	if len(v) == 0 {
		t.Errorf("levelDB error. Expected data got nothing")
	}
}

func TestFirst(t *testing.T) {
	res := dbTest.First(10, []byte("b"))

	if len(res) < 10 {
		t.Errorf("levelDB error. Expected 10 elements got %d", len(res))
	}
}
