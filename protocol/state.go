package protocol

type State interface {
	GetId() int
	Handle(packet *Packet)
}

type Handshaking struct {}

func (_ Handshaking) GetId() int {
	return -1
}