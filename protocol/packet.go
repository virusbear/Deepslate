package protocol

import (
	"encoding/binary"
	"io"
)

type PacketReader struct {
	io.Reader
}

type Packet struct {
	Id int32
	Data []byte
}

func NewPacketReader(reader io.Reader) *PacketReader {
	return &PacketReader{reader}
}

func (reader PacketReader) Read(threshold int32) *Packet {

}

func (reader PacketReader) readBool() (bool, error) {
	var value byte
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return false, err
	}

	return value != 0, nil
}

func (reader PacketReader) readByte() (int8, error) {
	var value int8
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (reader PacketReader) readUByte() (uint8, error) {
	var value uint8
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (reader PacketReader) readShort() (int16, error) {
	var value int16
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (reader PacketReader) readUShort() (uint16, error) {
	var value uint16
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (reader PacketReader) readInt() (int32, error) {
	var value int32
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (reader PacketReader) readLong() (int64, error) {
	var value int64
	if err := binary.Read(reader.Reader, binary.BigEndian, &value); err != nil {
		return 0, err
	}

	return value, nil
}

