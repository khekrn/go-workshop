package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error while connecting to tcp server : ", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error while accepting the connection = " + err.Error())
		}

		defer conn.Close()
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error while reading from connection : ", err.Error())
			return
		}

		message := buf[:size]
		fmt.Println("Received = ", string(message))
		conn.Write(message)
	}
}
