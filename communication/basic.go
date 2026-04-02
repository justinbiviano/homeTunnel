package communication

import "net"

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
