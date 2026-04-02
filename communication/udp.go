package communication

import (
	"fmt"
	"net"
)

// Dial a connection to Server ONLY from client (Client Side)
func UdpDial(address string) (net.Conn, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to %s over udp: %w", address, err)
	}
	return conn, nil
}

// Binds Port (Server Side)
func UdpListen(port string) (net.PacketConn, error) {
	conn, err := net.ListenPacket("udp", port)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect on port %s over udp: %w", port, err)
	}
	return conn, nil
}

// Reads Packets (Serverside)
func UdpRead(conn net.PacketConn) ([]byte, net.Addr, error) {
	buf := make([]byte, 1420)
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to read packet: %w", err)
	}
	return buf[:n], addr, nil
}

// Sends packets (serverside)
func UdpSend(conn net.PacketConn, data []byte, addr net.Addr) error {
	_, err := conn.WriteTo(data, addr)
	return err
}
