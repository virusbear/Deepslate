package nbt

import (
	"errors"
)

const intTypeId intType = 3

type intType int8

func (_ intType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	return IntTag{
		value: data,
	}, nil
}

func (_ intType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(IntTag)

	if !ok {
		return errors.New("incompatible tag. Expected INT")
	}

	return writer.writeInt32(data.value)
}

func (_ intType) GetId() int8 {
	return int8(intTypeId)
}

type IntTag struct {
	value int32
}

func (_ IntTag) dataType() dataType {
	return intTypeId
}