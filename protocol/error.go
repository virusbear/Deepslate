package protocol

import "fmt"

type IncompatibleProtocolVersion int

func (err IncompatibleProtocolVersion) Error() string {
	return fmt.Sprintf("incompatible protocol version %d", err)
}

type UnknownPacket int

func (err UnknownPacket) Error() string {
	return fmt.Sprintf("packet id %d not known", err)
}