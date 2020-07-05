package db

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v2"
)

// KeyValueStore is an abstract interface so Trippl can support more than one storage method
type KeyValueStore interface {
	GetObject([]byte, interface{}) error
	SetObject([]byte, interface{}) error
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	Close()
}

// EmbeddedKeyValueStoreStruct holds implementation of embedded key value store
type EmbeddedKeyValueStoreStruct struct {
	db *badger.DB
}

// EmbeddedKeyValueStore is an embedded Key Value storage (typically deployed with Docker)
func EmbeddedKeyValueStore(storageDir *string) (*EmbeddedKeyValueStoreStruct, error) {
	var opt badger.Options
	if storageDir == nil {
		opt = badger.DefaultOptions("").WithInMemory(true)
	} else {
		opt = badger.DefaultOptions(*storageDir)
	}
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	return &EmbeddedKeyValueStoreStruct{
		db: db,
	}, nil
}

func (r EmbeddedKeyValueStoreStruct) Close() {
	r.db.Close()
}

func (r EmbeddedKeyValueStoreStruct) Set(key []byte, value []byte) error {
	return r.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
}

func (r EmbeddedKeyValueStoreStruct) Get(key []byte) ([]byte, error) {
	var val []byte
	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(v []byte) error {
			val = append([]byte{}, v...)
			return nil
		})
	})
	return val, err
}

func (r EmbeddedKeyValueStoreStruct) GetObject(key []byte, out interface{}) error {
	res, err := r.Get(key)
	if err != nil {
		return err
	}
	if res == nil {
		return nil
	}

	return json.Unmarshal(res, out)
}

func (r EmbeddedKeyValueStoreStruct) SetObject(key []byte, d interface{}) error {
	jsonData, err := json.Marshal(&d)
	if err != nil {
		return err
	}

	return r.Set(key, jsonData)
}
