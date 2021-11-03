package nbt

import (
	"encoding/binary"
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) Writer {
	return Writer {
		writer: writer,
	}
}

func (nbt Writer) write(data interface{}) error {
	return binary.Write(nbt.writer, binary.LittleEndian, data)
}

func (nbt Writer) WriteInt8(data int8) error {
	return nbt.write(data)
}

func (nbt Writer) WriteInt16(data int16) error {
	return nbt.write(data)
}

func (nbt Writer) WriteInt32(data int32) error {
	return nbt.write(data)
}

func (nbt Writer) WriteInt64(data int64) error {
	return nbt.write(data)
}

func (nbt Writer) WriteString(data string) error {
	buf := []byte(data)

	if err := nbt.WriteInt32(int32(len(buf))); err != nil {
		return err
	}

	return nbt.write(buf)
}

func (nbt Writer) WriteFloat32(data float32) error {
	return nbt.write(data)
}

func (nbt Writer) WriteFloat64(data float64) error {
	return nbt.write(data)
}

func (nbt Writer) WriteByteArray(data []int8) error {
	if err := nbt.WriteInt32(int32(len(data))); err != nil {
		return err
	}

	return binary.Write(nbt.writer, binary.LittleEndian, data)
}

func (nbt Writer) WriteInt32Array(data []int32) error {
	if err := nbt.WriteInt32(int32(len(data))); err != nil {
		return err
	}

	return binary.Write(nbt.writer, binary.LittleEndian, data)
}

func (nbt Writer) WriteInt64Array(data []int64) error {
	if err := nbt.WriteInt32(int32(len(data))); err != nil {
		return err
	}

	return binary.Write(nbt.writer, binary.LittleEndian, data)
}
