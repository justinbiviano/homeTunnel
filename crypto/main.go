package main

import (
	"crypto/rand"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
)

type Keypair struct {
	PrivateKey [32]byte
	PublicKey  [32]byte
}

func generateKeypair() (Keypair, error) {
	var kp Keypair

	_, err := rand.Read(kp.PrivateKey[:])
	if err != nil {
		return Keypair{}, fmt.Errorf("failed to generate private key: %w", err)
	}

	pub, err := curve25519.X25519(kp.PrivateKey[:], curve25519.Basepoint)
	if err != nil {
		return Keypair{}, fmt.Errorf("failed to derive public key: %w", err)
	}

	copy(kp.PublicKey[:], pub)
	return kp, nil
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

func main() {
	client, _ := generateKeypair()
	server, _ := generateKeypair()

	fmt.Printf("Client public key: %x\n", client.PublicKey)
	fmt.Printf("Server public key: %x\n", server.PublicKey)

	secret, _ := curve25519.X25519(client.PrivateKey[:], server.PublicKey[:])
	fmt.Printf("Shared Secret:%x\n", secret)

	clientKey, serverKey, err := hashKeys(secret)
	if err != nil {
		fmt.Println("failed to derive keys:", err)
		return
	}

	fmt.Printf("Client key: %x\n", clientKey)
	fmt.Printf("Server key: %x\n", serverKey)

}
