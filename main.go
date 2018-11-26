package main

import (
	"flag"
	"log"
	"os"
	"topogopher/agent/messages"
)

func main() {
	start := flag.String("start", "server", "start server|client")
	port := flag.String("port", "8080", "port <port number>")
	flag.Parse()
	if *start == "server" {
		log.Printf("Launching %s on port %s \n", *start, *port)
	} else {
		log.Printf("Launching client %v\n", os.Args[2:])
	}
	switch *start {
	case "client":

		messages.StartClient(*port, os.Args[2:])
	default:
		messages.StartServer(*port)
	}

}
