package db

import "github.com/datagentleman/tiny-eth/pkg/db/leveldb"

var db DB

type DB interface {
	Close() error

	Get(key []byte) ([]byte, error)
	First(n uint64, prefix []byte) [][]byte
}

func Open(file string) error {
	d, err := leveldb.Open(file)

	db = d
	return err
}

func Get(key []byte) ([]byte, error) {
	return db.Get(key)
}
