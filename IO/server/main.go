// Sample code for handling connections in the server

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read file name and extension
	fileNameWithExtension, err := readString(conn)
	fmt.Println(fileNameWithExtension)
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}

	file, err := os.Create(fileNameWithExtension)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Receive file content
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error receiving file content:", err)
		return
	}

}

// Read a string from the connection
func readString(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func main() {
	listenAddr := "localhost:8080"
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
