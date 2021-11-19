package protocol

import (
	"Deepslate/log"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	conn    net.Conn
	state   State
	reader  PacketReader
	writer  PacketWriter
	onClose func(conn *Connection)
}

func NewConnection(conn net.Conn, onClose func(conn *Connection)) *Connection {
	c := &Connection{conn: conn, state: nil, reader: NewPacketReader(), writer: NewPacketWriter(), onClose: onClose }
	c.state = NewHandshakingState(c)
	return c
}

func (conn Connection) Send(packet *Packet) {
	if err := conn.writer.WritePacket(conn.conn, packet); err != nil {
		//TODO logging
		conn.Close()
	}
}

func (conn Connection) Handle() {
	logger := log.NewLogger("connection")
	defer conn.Close()

	for {
		packet, err := conn.reader.ReadPacket(conn.conn)
		if err == io.EOF {
			logger.Info("Connection closed")
			//TODO logging
			break
		}

		if err != nil {
			logger.Error(err)
			//TODO logging
			break
		}

		err = conn.state.Handle(packet)
		if err != nil {
			logger.Error(fmt.Errorf("error handling packet: %w", err))
			//TODO logging
			break
		}
	}
}

func (conn Connection) Close() error {
	conn.onClose(&conn)
	return conn.conn.Close()
}