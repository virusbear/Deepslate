package nbt

import (
	"errors"
)

type intType struct{}

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

type IntTag struct {
	value int32
}

func (_ IntTag) dataType() dataType {
	return intType{}
}

func (_ IntTag) Type() int8 {
	return TagInt
}

func (tag IntTag) Get() int32 {
	return tag.value
}

func (tag IntTag) Set(value int32) {
	tag.value = value
}

func NewInt(value int32) *IntTag {
	return &IntTag{
		value: value,
	}
}
