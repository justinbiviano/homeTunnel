package communication

import (
	"fmt"
	"net"
)

func NetDial(address string) (net.Conn, error) {
	// Dialing Conn From Client
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return conn, fmt.Errorf("Failed Dialing Connection: %w", err)
	}
	return conn, nil
}

func Read(conn net.Conn) ([32]byte, error) {
	// Read TCP from Conn
	var buf [32]byte
	_, err := conn.Read(buf[:])
	if err != nil {
		return [32]byte{}, err
	}
	return buf, nil
}

func Send(conn net.Conn, message [32]byte) error {
	// Send bytes over TCP from CONN
	_, err := conn.Write(message[:])
	return err
}

func NetListen(port string) (net.Listener, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("Failed Listening: %w", err)
	}
	return ln, nil
}

func NetAccept(ln net.Listener) (net.Conn, error) {
	conn, err := ln.Accept()
	if err != nil {
		return nil, fmt.Errorf("Failed Accpeting Conn: %w", err)
	}

	return conn, nil
}
