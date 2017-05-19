package fosite

import (
	"testing"

	"github.com/lhecker/argon2"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestArgon2Hash(t *testing.T) {
	h := &Argon2{
		Config: argon2.DefaultConfig(),
	}
	password := []byte("foo")
	hash, err := h.Hash(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.NotEqual(t, hash, password)
}

func TestArgon2CompareEquals(t *testing.T) {
	h := &Argon2{
		Config: argon2.DefaultConfig(),
	}
	password := []byte("foo")
	hash, err := h.Hash(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	err = h.Compare(hash, password)
	assert.Nil(t, err)
}

func TestArgon2CompareDifferent(t *testing.T) {
	h := &Argon2{
		Config: argon2.DefaultConfig(),
	}
	password := []byte("foo")
	hash, err := h.Hash(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	err = h.Compare(hash, []byte(uuid.NewRandom()))
	assert.NotNil(t, err)
}
