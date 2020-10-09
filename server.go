package main

import (
	"fmt"
	"net"
)

const (
	port = "8080"
)

var connections []net.Conn

// StartServer starts the Snicksnack server
func StartServer() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("An error occured")
	}

	fmt.Println("Server started on port", port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("An error occured")
		}

		connections = append(connections, connection)
		go handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection incoming")
	fmt.Println("Connections: ", len(connections))

	data := make([]byte, 1024)
	_, _ = conn.Read(data)
	fmt.Println(string(data))

}
