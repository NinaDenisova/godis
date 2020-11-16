package main

import (
	"fmt"
	"net"
	"os"
	"random/godis/internal"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "6379"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func parseRequest(buf []byte) []string {
	toString := string(buf)
	return strings.Split(toString, "\r\n")
}

func makeResponse(message string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(message), message)
}
func makeError(err error) string {
	return fmt.Sprintf("-%s\r\n", err.Error())
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	parsedRequest := parseRequest(buf)

	// print(string(d))
	godis := new(internal.Godis)

	command := parsedRequest[2]

	output, err := godis.Request(command)

	if err != nil {
		conn.Write([]byte(makeError(err)))
	} else {
		conn.Write([]byte(makeResponse(output)))
	}

	// Close the connection when you're done with it.
	conn.Close()
}
