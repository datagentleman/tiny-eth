package db

import (
	"github.com/datagentleman/tiny-eth/pkg/config"
	"github.com/datagentleman/tiny-eth/pkg/db/leveldb"
	"github.com/datagentleman/tiny-eth/pkg/logger"
)

var db DB

type DB interface {
	Close() error

	Get(key []byte) ([]byte, error)
	First(n uint64, prefix []byte) [][]byte
}

func Configure(conf *config.Config) {
	if db == nil {
		d, err := leveldb.Open(conf)
		if err != nil {
			logger.Panic(err)
			panic(err)
		}

		db = d
	}
}

func Get(key []byte) ([]byte, error) {
	return db.Get(key)
}

func Close() {
	db.Close()
}
