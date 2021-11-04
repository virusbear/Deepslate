package nbt

import (
	"errors"
)

const IntTypeId IntType = 3

type IntType int8

type IntTag struct {
	value int32
}

func (_ IntType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt32()

	if err != nil {
		return nil, err
	}

	return IntTag{
		value: data,
	}, nil
}

func (_ IntType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(IntTag)

	if !ok {
		return errors.New("incompatible tag. Expected INT")
	}

	return writer.WriteInt32(data.value)
}

func (_ IntType) GetId() int8 {
	return int8(IntTypeId)
}