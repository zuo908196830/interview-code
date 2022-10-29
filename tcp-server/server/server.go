package server

import (
	"errors"
	"fmt"
	"net"
	"tcp-server/router"
	"tcp-server/utils"
)

type Server struct {
	TCPVersion string
	Host       string
	Port       int
	Nagle      bool
	Router     router.Router
}

func NewServer() *Server {
	s := &Server{
		TCPVersion: "tcp4",
		Host:       utils.GlobalObject.Host,
		Port:       utils.GlobalObject.Port,
		Nagle:      utils.GlobalObject.Nagle,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(s.TCPVersion, fmt.Sprintf("%s:%d", s.Host, s.Port))
		if err != nil {
			return
		}
		listener, err := net.ListenTCP(s.TCPVersion, addr)
		if err != nil {
			return
		}
		for {
			//conn, err := listener.Accept()
			conn, err := listener.AcceptTCP()
			conn.SetNoDelay(s.Nagle)
			if err != nil {
				return
			}
			connection := NewConnection(s, conn)
			connection.Start()
		}
	}()
}

func (s *Server) Server() {
	s.Start()
	select {}
}

func (s *Server) AddRouter(router router.Router) error {
	if s.Router == nil {
		s.Router = router
		return nil
	}
	return errors.New("router exits")
}
