package protocol

const (
	IdServerHandshake = 0x00
	IdServerStatus = 0x00
	IdServerPing = 0x01
)

const (
	IdClientPong = 0x01
)

type State interface {
	GetId() int
	Handle(packet *Packet) error
	GetConnection() *Connection
}

type Handshaking struct {
	conn *Connection
}

func (_ Handshaking) GetId() int {
	return -1
}

func (state Handshaking) GetConnection() *Connection {
	return state.conn
}

func (state Handshaking) Handle(packet *Packet) error {
	switch packet.Id {
	case IdServerHandshake:
		packet, err := NewServerHandshakePacket(packet)
		if err != nil {
			return err
		}

		if packet.ProtocolVersion != VersionMc1171 {
			return IncompatibleProtocolVersion(packet.ProtocolVersion)
		}

		switch packet.NextState {
		case 1:
			state.conn.state = NewStatusState(state.conn)
		case 2:
			println("switching to login state")
			//TODO: state.conn.state = NewLoginState(state.conn)
		}
	}

	return nil
}

func NewHandshakingState(conn *Connection) State {
	return Handshaking{conn: conn}
}

type Status struct {
	conn *Connection
}

func (_ Status) GetId() int {
	return 1
}

func NewStatusState(conn *Connection) State {
	return Status{conn: conn}
}

func (state Status) Handle(packet *Packet) error {
	switch packet.Id {
	case IdServerStatus:

	case IdServerPing:
		request, err := NewServerPingPacket(packet)
		if err != nil {
			return err
		}
		state.conn.Send(NewClientPong(request.Payload))
	default:
		return UnknownPacket(packet.Id)
	}

	return nil
}

func (state Status) GetConnection() *Connection {
	return state.conn
}

type ServerPingPacket struct {
	Payload int64
}

func NewClientPong(payload int64) *Packet {
	packet := NewPacket(IdClientPong)
	packet.WriteLong(payload)
	return &packet
}

func NewServerPingPacket(packet *Packet) (*ServerPingPacket, error) {
	payload, err := packet.ReadLong()
	if err != nil {
		return nil, err
	}

	return &ServerPingPacket{
		Payload: payload,
	}, nil
}

type ServerHandshakePacket struct {
	ProtocolVersion int32
	ServerAddress string
	ServerPort uint16
	NextState int32
}

func NewServerHandshakePacket(packet *Packet) (*ServerHandshakePacket, error) {
	protocolVersion, err := packet.ReadVarInt()
	if err != nil {
		return nil, err
	}

	serverAddress, err := packet.ReadString(255)
	if err != nil {
		return nil, err
	}

	serverPort, err := packet.ReadUShort()
	if err != nil {
		return nil, err
	}

	nextState, err := packet.ReadVarInt()
	if err != nil {
		return nil, err
	}

	return &ServerHandshakePacket{
		ProtocolVersion: protocolVersion,
		ServerAddress: serverAddress,
		ServerPort: serverPort,
		NextState: nextState,
	}, nil
}