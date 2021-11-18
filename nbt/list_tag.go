package nbt

import (
	"errors"
	"fmt"
)

type listType struct{}

func (_ listType) Read(reader Reader) (Tag, error) {
	dtype, err := reader.readInt8()

	if err != nil {
		return nil, err
	}

	dataType, err := getDataType(dtype)

	if err != nil {
		return nil, err
	}

	length, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	list := make([]Tag, length)

	for i := range list {
		data, err := dataType.Read(reader)

		if err != nil {
			return nil, fmt.Errorf("unable to read list at index %d. Reason: %w", i, err)
		}

		list[i] = data
	}

	return ListTag{
		dType: dtype,
		value: list,
	}, nil
}

func (_ listType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ListTag)

	if !ok {
		return errors.New("incompatible tag. Expected LIST")
	}

	if err := writer.writeInt8(data.Type()); err != nil {
		return fmt.Errorf("unable to write list datatype. Reason: %w", err)
	}

	for i, value := range data.value {
		dataType, err := getDataType(data.dType)
		if err != nil {
			return fmt.Errorf("unable to get datatype for id %d. Reason: %w", data.dType, err)
		}

		if err := dataType.Write(writer, value); err != nil {
			return fmt.Errorf("unable to write list at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

type ListTag struct {
	dType int8
	value []Tag
}

func (_ ListTag) dataType() dataType {
	return listType{}
}

func (_ ListTag) Type() int8 {
	return TagList
}

func (tag ListTag) Raw() []Tag {
	return tag.value
}

func (tag ListTag) ContentType() int8 {
	return tag.dType
}

func (tag ListTag) Get(index int) *Tag {
	return &tag.value[index]
}

func (tag ListTag) Set(index int, value Tag) {
	tag.value[index] = value
}

func (tag ListTag) Length() int {
	return len(tag.value)
}

func NewList(size int, dtype int8) *ListTag {
	return &ListTag{
		dType: dtype,
		value: make([]Tag, size),
	}
}
