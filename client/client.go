package main

import (
	"fmt"

	"github.com/justinbiviano/homeTunnel/communication"
	"github.com/justinbiviano/homeTunnel/crypto"
	"golang.org/x/crypto/curve25519"
)

func StartUpClient(address string) ([32]byte, [32]byte, error) {
	public, private, err := crypto.GeneratePrivatePublicKeys()
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	conn, err := communication.NetDial(address)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	err = communication.Send(conn, public)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	recived, err := communication.Read(conn)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	secret, err := curve25519.X25519(private[:], recived[:])
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	conn.Close()

	clientKey, serverKey, err := crypto.HashKeys(secret)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	return clientKey, serverKey, nil
}

func main() {
	clientKey, serverKey, err := StartUpClient("127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(clientKey)
	fmt.Println(serverKey)
}
