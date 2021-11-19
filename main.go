package main

import (
	"Deepslate/server"
)

func main() {

	server, err := server.Listen(25565)
	if err != nil {
		panic(err)
	}

	server.Start()
}