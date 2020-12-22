package main

import (
	"bufio"
	"fmt"
	"net"
)

type user struct {
	connection net.Conn
	name       string
}

var users []user

// StartServer starts the Snicksnack server
func StartServer(port string) {
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
		scanner := bufio.NewScanner(conn)
		if scanner.Scan() {

			message := scanner.Text()

			if message == "exit" {
				break
			}

			for _, user := range users {
				if user.connection == conn {
					continue
				}

				_, err := fmt.Fprintf(user.connection, "%v: %v\n", name, message)
				if err != nil {
					fmt.Println("Error sending message to peers. ", err.Error())
				}
			}
		}

	}

	conn.Close()
}
