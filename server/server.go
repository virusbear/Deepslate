package server

import (
	"Deepslate/log"
	"fmt"
	"net"
)

var logger = log.NewLogger("Server")

type Server struct {
	listener *net.TCPListener
	handler ConnectionHandler
}

func Listen(port int) (*Server, error) {
	logger.Info(fmt.Sprintf("Listening on port %d", port))

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IP{0, 0, 0, 0}, Port: port})
	if err != nil {
		return nil, err
	}

	server := Server{
		listener: listener,
	}

	return &server, nil
}

func (server Server) Run() {
	logger.Info("Starting server")
	defer logger.Info("Server stopped")
	defer server.listener.Close()

	for {
		conn, err := server.listener.Accept()

		if err != nil {
			logger.Error(err)
		}

		go server.handler.Handle(conn)
	}
}