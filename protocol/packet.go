package protocol

import (
	"Deepslate/math"
	"Deepslate/nbt"
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
)

const (
	MaxStringLength = 32767
	MaxChatLength = 262144
	MaxIdentifierLength = 32767
	MaxCompressedPacketSize = 8388608
)

type Packet struct {
	Id int32
	data *bytes.Buffer
}

type PacketReader interface {
	ReadPacket(reader io.Reader) (*Packet, error)
}

type DefaultPacketReader struct {}
type CompressedPacketReader struct {
	Threshold int
}

func NewPacketReader() PacketReader {
	return &DefaultPacketReader{}
}

func NewCompressedPacketReader(threshold int) PacketReader {
	return &CompressedPacketReader{Threshold: threshold}
}

func readPacketPayload(reader io.Reader) (*bytes.Buffer, error) {
	length, err := readVarInt(reader)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, length)
	n, err := reader.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != len(buf) {
		return nil, fmt.Errorf("unable to read complete packet. Packetsize: %d, read bytes: %d", len(buf), n)
	}

	return bytes.NewBuffer(buf), nil
}

func readPacket(payload *bytes.Buffer) (*Packet, error) {
	packetId, err := readVarInt(payload)
	if err != nil {
		return nil, err
	}

	return &Packet{
		Id: packetId,
		data: bytes.NewBuffer(payload.Bytes()),
	}, nil
}

func (_ DefaultPacketReader) ReadPacket(reader io.Reader) (*Packet, error) {
	payload, err := readPacketPayload(reader)
	if err != nil {
		return nil, err
	}

	return readPacket(payload)
}

func (r CompressedPacketReader) ReadPacket(reader io.Reader) (*Packet, error) {
	payload, err := readPacketPayload(reader)
	if err != nil {
		return nil, err
	}

	dataLength, err := readVarInt(payload)
	if err != nil {
		return nil, err
	}

	if dataLength == 0 {
		return readPacket(payload)
	}

	if r.Threshold > 0 {
		if int(dataLength) < r.Threshold {
			return nil, fmt.Errorf("badly compressed packet - size of %d is below server threshold of %d", dataLength, r.Threshold)
		} else if dataLength > MaxCompressedPacketSize {
			return nil, fmt.Errorf("badly compressed packet - size of %d is larger than protocol maximum of %d", dataLength, MaxCompressedPacketSize)
		}
	}

	uncompressed, err := zlib.NewReader(payload)
	data, err := io.ReadAll(uncompressed)
	if err != nil {
		return nil, err
	}

	return readPacket(bytes.NewBuffer(data))
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

func (p Packet) ReadChat() (string, error) {
	return p.ReadString(MaxChatLength)
}

func (p Packet) ReadIdentifier() (string, error) {
	return p.ReadString(MaxIdentifierLength)
}

func (p Packet) ReadString(maxLength int) (string, error) {
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
	return readVarInt(p.data)
}

func (p Packet) ReadVarLong() (int64, error) {
	var value int64
	var offset int
	var current byte

	for {
		v, err := p.data.ReadByte()
		if err != nil {
			return 0, err
		}

		value |= int64(v & 0x7f) << (offset * 7)
		offset++

		if offset > 10 {
			return 0, errors.New("VarLong too big")
		}

		if (current & 0x80) != 0x80 {
			break
		}
	}

	return value, nil
}

func (p Packet) ReadNbt() (nbt.Tag, error) {
	_, tag, err := nbt.NewReader(p.data).Read()

	return tag, err
}

func (p Packet) ReadPosition() (*math.IVec3, error) {
	value, err := p.ReadLong()
	if err != nil {
		return nil, err
	}

	vec := math.IVec3{
		X: int(value >> 38),
		Y: int(value & 0xFFF),
		Z: int(value << 26 >> 38),
	}

	if vec.X >= 1 << 26 {
		vec.X -= 1 << 26
	}
	if vec.Y >= 1 << 12 {
		vec.Y -= 1 << 12
	}
	if vec.Z >= 1 << 26 {
		vec.Z -= 1 << 26
	}

	return &vec, nil
}

func (p Packet) ReadAngle() (byte, error) {
	return p.data.ReadByte()
}

func (p Packet) ReadUuid() (*uuid.UUID, error) {
	var buf [16]byte
	err := p.read(buf)
	if err != nil {
		return nil, err
	}

	value, err := uuid.FromBytes(buf[:])
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (p Packet) ReadOptionalPresent() (bool, error) {
	return p.ReadBool()
}

func (p Packet) ReadArrayLength() (int, error) {
	length, err := p.ReadVarInt()
	return int(length), err
}

func (p Packet) ReadByteArray() ([]int8, error) {
	length, err := p.ReadVarInt()
	if err != nil {
		return nil, err
	}

	array := make([]int8, length)
	err = p.read(&array)
	if err != nil {
		return nil, err
	}

	return array, nil
}

func readVarInt(reader io.Reader) (int32, error) {
	var value int32
	var offset int
	var current byte
	buf := make([]byte, 1)

	for {
		_, err := reader.Read(buf)
		if err != nil {
			return 0, err
		}

		value |= int32(buf[0] & 0x7f) << (offset * 7)
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