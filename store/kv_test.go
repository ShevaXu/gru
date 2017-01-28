package store

import (
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/ShevaXu/gru/utils"
)

const testBoltDBUri = "/tmp/test.db"

func TestMain(m *testing.M) {
	flag.Parse()

	// if not exist, will return "no such file or directory" error
	os.Remove(testBoltDBUri)

	// between set-up and clean-up
	ret := m.Run()

	// this one should do its job
	err := os.Remove(testBoltDBUri)
	if err != nil {
		// TODO
	}

	os.Exit(ret)
}

func TestNewBoltStore(t *testing.T) {
	assert := utils.NewAssert(t)

	s, err := NewBoltStore(testBoltDBUri)

	assert.NoError(err, "New BoltStore no error")
	assert.NotNil(s, "New BoltStore not nil")

	var (
		k0 = []byte("k0")
		v0 = []byte("v0")
		k1 = []byte("k1")
		v1 = []byte("v1")
	)

	// TestBoltStore_Get
	val, err := s.Get(k0)
	assert.NoError(err, "Get no error even non-exist")
	assert.Nil(val, "Get non-exist returns nil")

	// TestBoltStore_Set
	err = s.Set(k0, v0)
	assert.NoError(err, "Set no error")

	_ = s.Set(k1, v1)
	val, err = s.Get(k1)
	assert.NoError(err, "Get no error")
	assert.Equal(v1, val, "Get returns checked")

	// TestBoltStore_Peek
	err = s.Peek(k0, func(v []byte) error {
		assert.Equal(v0, v, "Peek access")
		return errors.New("Peek error")
	})
	assert.NotNil(err, "Peek error comes through")
}

//func TestNewBoltStore(t *testing.T) {
//	assert := utils.NewAssert(t)
//
//	s, err := NewBoltStore(testBoltDBUri)
//
//	assert.NoError(t, err, "New BoltStore no error")
//	assert.NotNil(t, s, "New BoltStore not nil")
//
//	var (
//		k0 = []byte("k0")
//		v0 = []byte("v0")
//		k1 = []byte("k1")
//		v1 = []byte("v1")
//	)
//
//	// TestBoltStore_Get
//	val, err := s.Get(k0)
//	assert.NoError(t, err, "Get no error even non-exist")
//	assert.Nil(t, val, "Get non-exist returns nil")
//
//	// TestBoltStore_Set
//	err = s.Set(k0, v0)
//	assert.NoError(t, err, "Set no error")
//
//	_ = s.Set(k1, v1)
//	val, err = s.Get(k1)
//	assert.NoError(t, err, "Get no error")
//	assert.Equal(t, v1, val, "Get returns checked")
//
//	// TestBoltStore_Peek
//	err = s.Peek(k0, func(v []byte) error {
//		assert.Equal(t, v0, v, "Peek access")
//		return errors.New("Peek error")
//	})
//	assert.Error(t, err, "Peek error comes through")
//}