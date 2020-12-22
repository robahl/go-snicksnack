package main

import (
	"flag"
)

func main() {
	server := flag.Bool("s", false, "Start a server")
	port := flag.String("p", "8080", "Port number")
	flag.Parse()

	if *server {
		StartServer(*port)
	} else {
		StartClient(*port)
	}
}
