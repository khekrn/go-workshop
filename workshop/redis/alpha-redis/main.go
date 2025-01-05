package main

import (
	"log"

	"coding2fun.in/alpha-redis/server"
)

func main() {
	red := "\033[31m"
	reset := "\033[0m"
	alpha := red + "𝛂" + "𝛂" + "𝛂" + "𝛂" + "𝛂" + reset
	log.Println("starting " + alpha + " redis")
	server.RunTcpServer()
}
