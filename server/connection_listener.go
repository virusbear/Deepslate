package server

import (
	"Deepslate/log"
	"Deepslate/protocol"
	"fmt"
	"net"
)

var logger = log.NewLogger("ConnectionListener")

type ConnectionListener struct {
	listener *net.TCPListener
	connections []*protocol.Connection
}

func Listen(port int) (*ConnectionListener, error) {
	logger.Info(fmt.Sprintf("Listening on port %d", port))

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IP{0, 0, 0, 0}, Port: port})
	if err != nil {
		return nil, err
	}

	server := ConnectionListener{
		listener: listener,
		connections: []*protocol.Connection{},
	}

	return &server, nil
}

func (server ConnectionListener) Start() {
	logger.Info("Starting tcp connection listener")
	defer logger.Info("tcp connection listener stopped")
	defer func() {
		err := server.listener.Close()
		logger.Warn(fmt.Errorf("error occured closing TCP listener: %w", err))
	}()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			logger.Error(fmt.Errorf("error accepting tcp connection. Ignoring connection. %w", err))
		}

		connection := protocol.NewConnection(conn, func(conn *protocol.Connection) { server.Remove(conn) })
		server.Put(connection)
		go connection.Handle()
	}
}

func (server ConnectionListener) Stop() {
	panic("not implemented!")
}

func (server ConnectionListener) Put(conn *protocol.Connection) {
	server.connections = append(server.connections, conn)
}

func (server ConnectionListener) Remove(conn *protocol.Connection) {
	if idx := server.IndexOf(conn); idx >= 0 {
		server.connections = append(server.connections[:idx], server.connections[idx + 1:]...)
	}
}

func (server ConnectionListener) IndexOf(conn *protocol.Connection) int {
	for i, c := range server.connections {
		if c == conn {
			return i
		}
	}

	return -1
}