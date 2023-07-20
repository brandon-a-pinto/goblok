package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	assert.Equal(t, len(privateKey.Bytes()), privateKeyLength)
	publicKey := privateKey.Public()
	assert.Equal(t, len(publicKey.Bytes()), publicKeyLength)
}
