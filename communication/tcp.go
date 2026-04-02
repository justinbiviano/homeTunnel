package communication

import (
	"fmt"
	"net"
)

func tcpNetDial(address string) (net.Conn, error) {
	// Dialing Conn From Client
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return conn, fmt.Errorf("Failed Dialing Connection: %w", err)
	}
	return conn, nil
}

func tcpNetListen(port string) (net.Listener, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("Failed Listening: %w", err)
	}
	return ln, nil
}

func tcpNetAccept(ln net.Listener) (net.Conn, error) {
	conn, err := ln.Accept()
	if err != nil {
		return nil, fmt.Errorf("Failed Accpeting Conn: %w", err)
	}

	return conn, nil
}
