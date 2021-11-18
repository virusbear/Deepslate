package nbt

import (
	"encoding/binary"
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) Writer {
	return Writer{
		writer: writer,
	}
}

func (nbt Writer) Write(name string, tag Tag) error {
	err := nbt.writeInt8(tag.Type())
	if err != nil {
		return err
	}

	err = nbt.writeString(name)
	if err != nil {
		return err
	}

	return tag.dataType().Write(nbt, tag)
}

func (nbt Writer) write(data interface{}) error {
	return binary.Write(nbt.writer, binary.BigEndian, data)
}

func (nbt Writer) writeInt8(data int8) error {
	return nbt.write(data)
}

func (nbt Writer) writeInt16(data int16) error {
	return nbt.write(data)
}

func (nbt Writer) writeInt32(data int32) error {
	return nbt.write(data)
}

func (nbt Writer) writeInt64(data int64) error {
	return nbt.write(data)
}

func (nbt Writer) writeString(data string) error {
	buf := []byte(data)

	if err := nbt.writeInt32(int32(len(buf))); err != nil {
		return err
	}

	return nbt.write(buf)
}

func (nbt Writer) writeFloat32(data float32) error {
	return nbt.write(data)
}

func (nbt Writer) writeFloat64(data float64) error {
	return nbt.write(data)
}

func (nbt Writer) writeByteArray(data []int8) error {
	if err := nbt.writeInt32(int32(len(data))); err != nil {
		return err
	}

	return binary.Write(nbt.writer, binary.BigEndian, data)
}

func (nbt Writer) writeInt32Array(data []int32) error {
	if err := nbt.writeInt32(int32(len(data))); err != nil {
		return err
	}

	return binary.Write(nbt.writer, binary.BigEndian, data)
}

func (nbt Writer) writeInt64Array(data []int64) error {
	if err := nbt.writeInt32(int32(len(data))); err != nil {
		return err
	}

	return binary.Write(nbt.writer, binary.BigEndian, data)
}