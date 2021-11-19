package server

import (
	"Deepslate/protocol"
	"io"
	"net"
)

type Connection struct {
	conn net.Conn
	state protocol.State
	reader protocol.PacketReader
	onClose func(conn *Connection)
}

func NewConnection(conn net.Conn, onClose func(conn *Connection)) *Connection {
	return &Connection{conn: conn, state: protocol.Handshaking{}, reader: protocol.NewPacketReader(), onClose: onClose }
}

func (conn Connection) Handle() {
	defer conn.Close()

	for {
		packet, err := conn.reader.ReadPacket(conn.conn)
		if err == io.EOF {
			//TODO logging
			break
		}

		if err != nil {
			//TODO logging
			break
		}

		conn.state.Handle(packet)
	}
}

func (conn Connection) Close() error {
	conn.onClose(&conn)
	return conn.conn.Close()
}