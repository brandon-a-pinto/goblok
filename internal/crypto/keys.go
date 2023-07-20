package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

const (
	privateKeyLength = 64
	publicKeyLength  = 32
	seedLength       = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

func (p *PrivateKey) Public() *PublicKey {
	buffer := make([]byte, publicKeyLength)
	copy(buffer, p.key[:32])
	return &PublicKey{
		key: buffer,
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLength)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}
