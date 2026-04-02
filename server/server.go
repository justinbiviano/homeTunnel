package main

import (
	"fmt"
	"net"

	"github.com/justinbiviano/homeTunnel/communication"
	"github.com/justinbiviano/homeTunnel/crypto"
	"golang.org/x/crypto/curve25519"
)

func StartUpServer(ln net.Listener) ([32]byte, [32]byte, error) {
	public, private, err := crypto.GeneratePrivatePublicKeys()
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	conn, err := communication.TcpNetAccept(ln)

	recived, err := communication.KeyRead(conn)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	err = communication.KeySend(conn, public)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	secret, err := curve25519.X25519(private[:], recived[:])
	if err != nil {
		return [32]byte{}, [32]byte{}, fmt.Errorf("Failed finding secrete from keys: %w", err)
	}

	conn.Close()

	clientKey, serverKey, err := crypto.HashKeys(secret)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	return clientKey, serverKey, nil
}

func main() {
	// Creates Net Listener
	var port string = ":8080"
	ln, err := communication.TcpNetListen(port)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Handels and loops to all connections
	for {
		clientKey, serverKey, err := StartUpServer(ln)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(clientKey)
		fmt.Println(serverKey)
	}
}
