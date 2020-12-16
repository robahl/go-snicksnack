package main

import (
	"fmt"
	"net"
)

const (
	port = "8080"
)

type user struct {
	connection net.Conn
	name       string
}

var users []user

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

		go handleConnection(connection)
		fmt.Println("Users ", users)
	}
}

func handleConnection(conn net.Conn) {
	_, _ = fmt.Fprint(conn, "Name: ")
	var name string
	_, _ = fmt.Fscanln(conn, &name)
	users = append(users, user{conn, name})

	for _, user := range users {
		_, _ = fmt.Fprintf(user.connection, "%v connected!\n", name)
	}

	for {
		var incomingMessage string
		fmt.Fscanf(conn, "%v\n", &incomingMessage)

		if incomingMessage == "exit" {
			break
		}

		// TODO - Create a []byte chunk in memory and read into it

		fmt.Println("LOG: ", incomingMessage)

		for _, user := range users {
			if user.connection == conn {
				continue
			}

			_, _ = fmt.Fprintf(user.connection, "%v: %v\n", name, incomingMessage)
		}
	}

	conn.Close()

}
