package communication

import (
	"fmt"
	"net"
)

func netDial(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return conn, fmt.Errorf("Failed Dialing Connection: %w", err)
	}
	defer conn.Close()
	return conn, nil
}

func read(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return string(buf), err
	}
	return string(buf[:n]), nil
}

func send(conn net.Conn, message string) {
	fmt.Fprintf(conn, "%s", message)
}

func netListen(port string) (net.Conn, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("Failed Listening: %w", err)
	}

	conn, err := ln.Accept()
	if err != nil {
		return nil, fmt.Errorf("Failed Listening: %w", err)
	}

	defer conn.Close()
	return conn, nil
}
