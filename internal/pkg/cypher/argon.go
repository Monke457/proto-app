package cypher 

import (
	"bytes"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
)

type HashSalt struct {
	Hash []byte
	Salt []byte
}

type Argon2ID struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

func New(time, memory, keyLen, saltLen uint32, threads uint8) *Argon2ID {
	return &Argon2ID{
		time: time,
		memory: memory,
		keyLen: keyLen,
		saltLen: saltLen,
		threads: threads,
	}
}

func (a *Argon2ID) Encrypt(password []byte, salt []byte) (*HashSalt, error) {
	var err error
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}
	if err != nil {
		return nil, err
	}
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
	return &HashSalt{ Hash: hash, Salt: salt }, nil
}

func (a *Argon2ID) Compare(hash, salt, password []byte) error {
	hashSalt, err := a.Encrypt(password, salt)
	if err != nil {
		return err
	}
	if !bytes.Equal(hash, hashSalt.Hash) {
		return errors.New("Hash doesn't match")
	}
	return nil
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}
