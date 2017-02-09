package store_test

import (
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ShevaXu/gru/store"
	"github.com/ShevaXu/gru/utils"
)

var tempBoltDBUri string

func TestMain(m *testing.M) {
	flag.Parse()

	f, _ := ioutil.TempFile(os.TempDir(), "test")
	tempBoltDBUri = f.Name()

	// between set-up and clean-up
	ret := m.Run()

	// this one should do its job
	err := os.Remove(tempBoltDBUri)
	if err != nil {
		panic(err)
	}

	os.Exit(ret)
}

func TestBoltStore(t *testing.T) {
	assert := utils.NewAssert(t)

	s, err := store.NewBoltStore(tempBoltDBUri)

	assert.NoError(err, "New BoltStore no error")
	assert.NotNil(s, "New BoltStore not nil")

	var (
		k0 = []byte("k0")
		v0 = []byte("v0")
		k1 = []byte("k1")
		v1 = []byte("v1")
	)

	val, err := s.Get(k0)
	assert.NoError(err, "Get no error even non-exist")
	assert.Nil(val, "Get non-exist returns nil")

	err = s.Set(k0, v0)
	assert.NoError(err, "Set no error")

	_ = s.Set(k1, v1)
	val, err = s.Get(k1)
	assert.NoError(err, "Get no error")
	assert.Equal(v1, val, "Get returns checked")

	err = s.Peek(k0, func(v []byte) error {
		assert.Equal(v0, v, "Peek access")
		return errors.New("Peek error")
	})
	assert.NotNil(err, "Peek error comes through")
}
