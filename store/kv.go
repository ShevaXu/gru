package store

import (
	"time"

	"github.com/boltdb/bolt"
)

// KVStore is a simple key-value database.
type KVStore interface {
	// Get retrieves the value by certain key,
	// and returns nil if key is not existed.
	Get(key []byte) (value []byte, err error)

	// Set stores the key-value pair with
	// the update-if-existed fashion.
	Set(key, value []byte) error
}

// BoltStore implements the KVStore with boltDB as backend;
// it uses a single bucket to simplify interaction with DB.
type BoltStore struct {
	db         *bolt.DB
	bucketName []byte
}

// Set stores key-value pair in a boltDB update-transaction.
func (s *BoltStore) Set(key, value []byte) error {
	// update transaction is lock-protected
	return s.db.Update(func(tx *bolt.Tx) error {
		// bucket must exist
		b := tx.Bucket(s.bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Put(key, value)
	})
}

// Get copies the returning []byte from db.View() to
// make sure it keeps valid outside the transaction.
// Refer to https://github.com/boltdb/bolt#using-keyvalue-pairs.
func (s *BoltStore) Get(key []byte) (value []byte, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		// bucket must exist
		b := tx.Bucket(s.bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		buf := b.Get(key)
		if buf != nil {
			value = make([]byte, len(buf))
			copy(value, buf)
		}
		// no value is not error
		return nil
	})
	return
}

// PeekFunc provides the callback-style access to the value.
type PeekFunc func([]byte) error

// Peek allows access to the value at given key through the PeekFunc;
// note that values from BoltDB's View are only transaction-valid.
func (s *BoltStore) Peek(key []byte, f PeekFunc) error {
	return s.db.View(func(tx *bolt.Tx) error {
		// bucket must exist
		b := tx.Bucket(s.bucketName)
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return f(b.Get(key))
	})
}

// theBoltBucket is the fixed bucket name of each BoltStore instance.
const theBoltBucket = "single"

// NewBoltStore returns a NewBoltStore instance;
// it will open the BoltDB and create the bucket if needed.
// It is by design that only one bucket is used per DB
// and its name is fixed.
func NewBoltStore(path string) (*BoltStore, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(theBoltBucket))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &BoltStore{
		db:         db,
		bucketName: []byte(theBoltBucket),
	}, nil
}
