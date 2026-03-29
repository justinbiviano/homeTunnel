package main

import (
	"fmt"
	"net"
	"sync"
)

func listen(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Client:", string(buf[:n]))
}

func message(conn net.Conn) {
	var message string
	fmt.Scanln(&message)
	fmt.Print("\033[1A\033[2K")
	fmt.Println("Server:", message)
	fmt.Fprintf(conn, "%s", message)
}

func main() {
	var wg sync.WaitGroup

	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()
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
