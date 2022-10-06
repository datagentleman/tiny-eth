package db

import (
	"fmt"

	"github.com/datagentleman/tiny-eth/pkg/db/leveldb"
)

var db DB

type DB interface {
	Close() error

	Get(key []byte) ([]byte, error)
	First(n uint64, prefix []byte) [][]byte
}

func Configure(config map[string]interface{}) {
	d, err := leveldb.Open(config["path"].(string))
	if err != nil {
		fmt.Println(err)
	}

	db = d

}

func Get(key []byte) ([]byte, error) {
	return db.Get(key)
}
