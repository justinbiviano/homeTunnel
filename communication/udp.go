package communication

import (
	"fmt"
	"net"
)

// Dial a connection to Server ONLY from client (Client Side)
func udpDial(address string) (net.Conn, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return conn, fmt.Errorf("Failed to connect to %s over udp: %w", address, err)
	}
	return conn, nil
}

// Binds Port (Server Side)
func udpListen(port string) (net.PacketConn, error) {
	conn, err := net.ListenPacket("udp", port)
	if err != nil {
		return conn, fmt.Errorf("Failed to connect on port %s over udp: %w", port, err)
	}
	return conn, nil
}

// Reads Packets (Serverside)
func udpRead(conn net.PacketConn, buf []byte) (int, net.Addr, error) {

}

// Sends packets (serverside)
func udpSend(conn net.PacketConn, buf []byte, addr net.Addr) error {

}
