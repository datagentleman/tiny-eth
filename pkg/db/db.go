package db

import (
	"fmt"

	"github.com/datagentleman/tiny-eth/pkg/config"
	"github.com/datagentleman/tiny-eth/pkg/db/leveldb"
)

var db DB

type DB interface {
	Close() error

	Get(key []byte) ([]byte, error)
	First(n uint64, prefix []byte) [][]byte
}

func init() {
	conf, err := config.Load("../../../config/database.json")
	if err != nil {
		fmt.Println(err)
	}

	d, err := leveldb.Open(conf["path"].(string))
	if err != nil {
		fmt.Println(err)
	}

	db = d
}

func Get(key []byte) ([]byte, error) {
	return db.Get(key)
}
