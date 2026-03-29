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

	// Random Generation of Private Key
	_, err := rand.Read(privateKey[:])
	if err != nil {
		return publicKey, privateKey, fmt.Errorf("Failed to Generate Private Key: %w", err)
	}

	// Derives Public Keys From Private Keys
	publicKeyStr, err := curve25519.X25519(privateKey[:], curve25519.Basepoint)
	if err != nil {
		return publicKey, privateKey, fmt.Errorf("Failed to Derive the Public Key: %w", err)
	}

	copy(publicKey[:], publicKeyStr)
	return publicKey, privateKey, nil
}

func hashKeys(secret []byte) ([32]byte, [32]byte, error) {
	// Hashes Keys Using BLAKE2S.
	reader := hkdf.New(func() hash.Hash {
		h, _ := blake2s.New256(nil)
		return h
	}, secret, nil, []byte("homeTunnel")) // Uses the "homeTunnel as a 'mixing key.

	// Seperates Keys into Server and Client Keys
	var clientKey, serverKey [32]byte
	reader.Read(clientKey[:])
	reader.Read(serverKey[:])

	return clientKey, serverKey, nil
}
