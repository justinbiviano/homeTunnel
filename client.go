package main

import (
	"fmt"
	"net"
	"sync"
)

func listen(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Server:", string(buf[:n]))
}

func message(conn net.Conn) {
	var message string
	fmt.Scanln(&message)
	fmt.Print("\033[1A\033[2K")
	fmt.Println("Client:", message)
	fmt.Fprintf(conn, message)
}

func main() {
	var wg sync.WaitGroup

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			listen(conn)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			message(conn)
		}
	}()

	wg.Wait()
}
