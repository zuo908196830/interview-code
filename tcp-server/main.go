package main

import (
	"tcp-server/router"
	"tcp-server/server"
)

func main() {
	Server := server.NewServer()
	Server.AddRouter(&router.AddRouter{})
	Server.Server()
}
