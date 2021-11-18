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
	return Reader{
		reader: reader,
	}
}

func (nbt Reader) Read() (string, Tag, error) {
	dtype, err := nbt.readInt8()
	if err != nil {
		return "", nil, err
	}

	dataType, err := getDataType(dtype)
	if err != nil {
		return "", nil, err
	}

	if dtype == TagEnd {
		tag, err := dataType.Read(nbt)
		return "", tag, err
	}

	name, err := nbt.readString()
	if err != nil {
		return "", nil, err
	}

	tag, err := dataType.Read(nbt)
	return name, tag, err
}

func (nbt Reader) read(data interface{}) error {
	return binary.Read(nbt.reader, binary.BigEndian, data)
}

func (nbt Reader) readInt8() (int8, error) {
	var data int8
	err := nbt.read(&data)

	return data, err
}

func (nbt Reader) readInt16() (int16, error) {
	var data int16
	err := nbt.read(&data)

	return data, err
}

func (nbt Reader) readInt32() (int32, error) {
	var data int32
	err := nbt.read(&data)

	return data, err
}

func (nbt Reader) readInt64() (int64, error) {
	var data int64
	err := nbt.read(&data)

	return data, err
}

func (nbt Reader) readString() (string, error) {
	length, err := nbt.readInt16()
	if err != nil {
		return "", err
	}

	if length < 0 {
		return "", errors.New("invalid NBT string: Length negative")
	}

	buf := make([]byte, length)
	err = nbt.read(&buf)

	return string(buf), err
}

func (nbt Reader) readFloat32() (float32, error) {
	var data float32
	err := nbt.read(&data)

	return data, err
}

func (nbt Reader) readFloat64() (float64, error) {
	var data float64
	err := nbt.read(&data)

	return data, err
}

func (nbt Reader) readByteArray() ([]int8, error) {
	length, err := nbt.readInt32()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("invalid NBT bytearray: Length negative")
	}

	buf := make([]int8, length)
	err = nbt.read(&buf)

	return buf, err
}

func (nbt Reader) readInt32Array() ([]int32, error) {
	length, err := nbt.readInt32()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("invalid NBT bytearray: Length negative")
	}

	buf := make([]int32, length)
	err = nbt.read(&buf)

	return buf, err
}

func (nbt Reader) readInt64Array() ([]int64, error) {
	length, err := nbt.readInt32()
	if err != nil {
		return nil, err
	}

	if length < 0 {
		return nil, errors.New("invalid NBT bytearray: Length negative")
	}

	buf := make([]int64, length)
	err = nbt.read(&buf)

	return buf, err
}
