package server

import (
	"io"
	"log"
	"net"
)

const (
	Host = "0.0.0.0"
	Port = "7379"
)

func RunTcpServer() {
	log.Println("starting redis tcp server")

	listener, err := net.Listen("tcp", Host+":"+Port)
	if err != nil {
		log.Println("error on listening : ", err.Error())
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error while accepting the connection : ", err.Error())
			panic(err)
		}
		go handleCommand(conn)
	}
}

func handleCommand(conn net.Conn) {
	for {
		cmd, err := readCommand(conn)
		if err != nil {
			conn.Close()
			log.Println("client disconnected : ", conn.RemoteAddr())
			if err == io.EOF {
				break
			}
			log.Println("error :", err.Error())
		}
		log.Println("command = ", cmd)
		if err = respond(cmd, conn); err != nil {
			log.Println("error while writing : ", err.Error())
		}
	}
}

func readCommand(conn net.Conn) (string, error) {
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, conn net.Conn) error {
	if _, err := conn.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
