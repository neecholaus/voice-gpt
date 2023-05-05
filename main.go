package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Remove("./socket.sock")

	socket, err := net.Listen("unix", "./socket.sock")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer socket.Close()

	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}

		// Read from the connection
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		// Print the received message
		fmt.Println(string(buf[:n]))

		// Close the connection
		conn.Close()
	}

}
