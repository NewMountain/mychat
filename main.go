package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listens on the specified IP address")

	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		os.Args
		fmt.Println("is host")
	} else {
		// go run main.go <ip>
		fmt.Println("is guest")
	}
}
