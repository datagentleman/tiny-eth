package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Level struct {
	db *leveldb.DB
}

func Open(file string) (*Level, error) {
	db, err := leveldb.OpenFile(file, nil)
	if err != nil {
		return nil, err
	}

	return &Level{db: db}, nil
}

func (lvl *Level) Close() error {
	return lvl.db.Close()
}

func (lvl *Level) Get(key []byte) ([]byte, error) {
	val, err := lvl.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (lvl *Level) First(n uint64, prefix []byte) [][]byte {
	iter := lvl.db.NewIterator(util.BytesPrefix(prefix), nil)

	num := uint64(0)
	res := [][]byte{}

	for iter.Next() {
		res = append(res, iter.Key())

		num++
		if num == n {
			break
		}
	}

	iter.Release()
	return res
}
