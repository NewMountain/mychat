package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const port = "8080"

// RunGuest takes an ip conection and does cool stuff.
func RunGuest(ip string) {

}

// RunHost is a function that takes an ip as an argument and listens
// for connections on that IP.
func RunHost(ip string) {
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
