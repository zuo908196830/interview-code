package main

import (
	"tcp-server/router"
	"tcp-server/server"
)

func main() {
	Server := server.NewServer("tcp4", "127.0.0.1", 8001)
	Server.AddRouter(&router.AddRouter{})
	Server.Server()
}
