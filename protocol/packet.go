package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	MaxStringLength = 32767
	MaxChatLength = 262144
	MaxIdentifierLength = 32767
)

type Packet struct {
	Id int32
	data *bytes.Buffer
}

func (p Packet) read(v interface{}) error {
	return binary.Read(p.data, binary.BigEndian, v)
}

func (p Packet) ReadBool() (bool, error) {
	var value byte
	if err := p.read(value); err != nil {
		return false, err
	}

	return value != 0, nil
}

func (p Packet) ReadByte() (int8, error) {
	var value int8
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadUByte() (uint8, error) {
	var value uint8
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadShort() (int16, error) {
	var value int16
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadUShort() (uint16, error) {
	var value uint16
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadInt() (int32, error) {
	var value int32
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadLong() (int64, error) {
	var value int64
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadFloat() (float32, error) {
	var value float32
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadDouble() (float64, error) {
	var value float64
	if err := p.read(value); err != nil {
		return 0, err
	}

	return value, nil
}

func (p Packet) ReadString() (string, error) {
	return p.readString(MaxStringLength)
}

func (p Packet) ReadChat() (string, error) {
	return p.readString(MaxChatLength)
}

func (p Packet) ReadIdentifier() (string, error) {
	return p.readString(MaxIdentifierLength)
}

func (p Packet) readString(maxLength int) (string, error) {
	length, err := p.ReadVarInt()
	if err != nil {
		return "", err
	}
	if int(length) > maxLength * 4 {
		return "", fmt.Errorf("the received encoded string buffer length is longer than maximum allowed (%d > %d)", length, maxLength * 4)
	}
	if length < 0 {
		return "", errors.New("the received encoded string buffer length is less than zero! Weird string!")
	}

	value := make([]rune, 0)
	for size := 0; size < int(length); {
		r, s, err := p.data.ReadRune()
		if err != nil {
			return "", err
		}
		size += s

		value = append(value, r)
	}

	return string(value), nil
}

func (p Packet) ReadVarInt() (int32, error) {
	var value int32
	var offset int
	var current byte

	for {
		v, err := p.data.ReadByte()
		if err != nil {
			return 0, err
		}

		value |= (v & 0x7f) << offset * 7
		offset++

		if offset > 5 {
			return 0, errors.New("VarInt too big")
		}

		if (current & 0x80) != 0x80 {
			break
		}
	}

	return value, nil
}