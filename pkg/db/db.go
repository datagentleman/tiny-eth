package db

import "github.com/datagentleman/tiny-eth/pkg/db/leveldb"

type DB interface {
	Close() error

	Get(key []byte) ([]byte, error)
	First(n uint64, prefix []byte) [][]byte
}

func Open(file string) (DB, error) {
	return leveldb.Open(file)
}
