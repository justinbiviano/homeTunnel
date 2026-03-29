package crypto

import (
	"crypto/rand"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
)

func GeneratePrivatePublicKeys() ([32]byte, [32]byte, error) {
	var privateKey, publicKey [32]byte
	_, err := rand.Read(privateKey[:])
	if err != nil {
		return publicKey, privateKey, fmt.Errorf("Failed to Generate Private Key: %w", err)
	}

	publicKeyStr, err := curve25519.X25519(privateKey[:], curve25519.Basepoint)
	if err != nil {
		return publicKey, privateKey, fmt.Errorf("Failed to Derive the Public Key: %w", err)
	}

	copy(publicKey[:], publicKeyStr)
	return publicKey, privateKey, nil
}

func hashKeys(secret []byte) ([32]byte, [32]byte, error) {
	reader := hkdf.New(func() hash.Hash {
		h, _ := blake2s.New256(nil)
		return h
	}, secret, nil, []byte("homeTunnel"))

	var clientKey, serverKey [32]byte
	reader.Read(clientKey[:])
	reader.Read(serverKey[:])

	return clientKey, serverKey, nil
}
