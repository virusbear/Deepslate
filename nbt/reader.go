package nbt

import (
	"encoding/binary"
	"errors"
	"io"
)

type Reader struct {
	reader io.Reader
}

func NewReader(reader io.Reader) Reader {
	return Reader {
		reader: reader,
	}
}

func (nbt Reader) ReadInt8() (int8, error) {
	var data int8
	err := binary.Read(nbt.reader, binary.LittleEndian, data)

	return data, err
}

func (nbt Reader) ReadInt16() (int16, error) {
	var data int16
	err := binary.Read(nbt.reader, binary.LittleEndian, data)

	return data, err
}

func (nbt Reader) ReadInt32() (int32, error) {
	var data int32
	err := binary.Read(nbt.reader, binary.LittleEndian, data)

	return data, err
}

func (nbt Reader) ReadInt64() (int64, error) {
	var data int64
	err := binary.Read(nbt.reader, binary.LittleEndian, data)

	return data, err
}

func (nbt Reader) ReadString() (string, error) {
	length, err := nbt.ReadInt16()
	if err != nil {
		return "", err
	}

	if length < 0 {
		return "", errors.New("invalid NBT string: Length negative")
	}

	buf := make([]byte, length)
	err = binary.Read(nbt.reader, binary.LittleEndian, buf)

	return string(buf), err
}

func (nbt Reader) ReadFloat32() (float32, error) {
	var data float32
	err := binary.Read(nbt.reader, binary.LittleEndian, data)

	return data, err
}

func (nbt Reader) ReadFloat64() (float64, error) {
	var data float64
	err := binary.Read(nbt.reader, binary.LittleEndian, data)

	return data, err
}

func (nbt Reader) ReadByteArray() ([]int8, error) {
	length, err := nbt.ReadInt32()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("invalid NBT bytearray: Length negative")
	}

	buf := make([]int8, length)
	err = binary.Read(nbt.reader, binary.LittleEndian, buf)

	return buf, err
}

func (nbt Reader) ReadInt32Array() ([]int32, error) {
	length, err := nbt.ReadInt32()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("invalid NBT bytearray: Length negative")
	}

	buf := make([]int32, length)
	err = binary.Read(nbt.reader, binary.LittleEndian, buf)

	return buf, err
}

func (nbt Reader) ReadInt64Array() ([]int64, error) {
	length, err := nbt.ReadInt32()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("invalid NBT bytearray: Length negative")
	}

	buf := make([]int64, length)
	err = binary.Read(nbt.reader, binary.LittleEndian, buf)

	return buf, err
}