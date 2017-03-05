package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listens on the specified IP address")

	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIP)
	}
}

const port = "8080"

func runGuest(ip string) {

}

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		// Print with output and exit with statcode 1
		log.Fatal("Error", listenErr)
	}
	conn, acceptErr := listener.Accept()

	if acceptErr != nil {
		log.Fatal("Error", acceptErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error", readErr)
	}

	fmt.Println("Message recevied:", message)
}
