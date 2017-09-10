package keys

import (
	"chain/crypto/ed25519"
	"crypto/rand"
)

func GenerateED25519Key() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}
