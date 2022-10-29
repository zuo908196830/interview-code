package server

import (
	"bufio"
	"net"
	"tcp-server/utils"
)

type Connection struct {
	Conn     net.Conn
	Server   *Server
	ExitChan chan bool
	MsgChan  chan string
}

func NewConnection(server *Server, conn net.Conn) *Connection {
	connection := &Connection{
		Conn:     conn,
		Server:   server,
		ExitChan: make(chan bool),
		MsgChan:  make(chan string),
	}
	return connection
}

func (c *Connection) Start() {
	go c.Read()
	go c.Write()
}

func (c *Connection) Read() {
	defer c.Stop()
	reader := bufio.NewReader(c.Conn)
	for {
		msg, err := utils.Decode(reader)
		if err != nil {
			break
		}
		go c.Server.Router.Handler(msg, c.MsgChan)
	}
}

func (c *Connection) Write() {
	for {
		select {
		case msg := <-c.MsgChan:
			_, err := c.Conn.Write(utils.Encode(msg))
			if err != nil {
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	c.ExitChan <- true
	close(c.ExitChan)
	close(c.MsgChan)
	c.Conn.Close()
}
