package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("server or client?")
		return
	}

	if os.Args[1] == "server" {
		StartServer()
	} else if os.Args[1] == "client" {
		StartClient()
	}
}
