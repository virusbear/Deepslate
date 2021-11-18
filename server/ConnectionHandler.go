package server

import (
	"Deepslate/log"
	"Deepslate/protocol"
	"net"
)

type ConnectionHandler struct {
	connections []net.Conn
}

func (handler ConnectionHandler) Handle(conn net.Conn) {
	logger := log.NewLogger("connectionHandler(conn)")
	defer logger.Info("Connection closed")
	defer conn.Close()

	for {
		logger.Info(*protocol.ReadPacket(conn))
	}
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{connections: []net.Conn{}}
}