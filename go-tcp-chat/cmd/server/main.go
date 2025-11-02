package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	fmt.Println("Server is listening on port 8080")

	if err != nil {
		panic(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		conn.Write([]byte("Echo: " + message))
	}
}
