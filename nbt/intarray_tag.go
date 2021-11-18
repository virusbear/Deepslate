package nbt

import (
	"errors"
	"fmt"
)

type intArrayType struct{}

func (_ intArrayType) Read(reader Reader) (Tag, error) {
	length, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	data := make([]int32, length)

	for i, _ := range data {
		value, err := reader.readInt32()

		if err != nil {
			return nil, fmt.Errorf("unable to read int array at index %d. Reason: %w", i, err)
		}

		data[i] = value
	}

	return IntArrayTag{
		value: data,
	}, nil
}

func (_ intArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(IntArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected INTARRAY")
	}

	if err := writer.writeInt32(int32(len(data.value))); err != nil {
		return err
	}

	for i, value := range data.value {
		if err := writer.writeInt32(value); err != nil {
			return fmt.Errorf("unable to write int array at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

type IntArrayTag struct {
	value []int32
}

func (_ IntArrayTag) dataType() dataType {
	return intArrayType{}
}

func (_ IntArrayTag) Type() int8 {
	return TagIntArray
}

func (arr IntArrayTag) Raw() []int32 {
	return arr.value
}

func (arr IntArrayTag) Get(index int) int32 {
	return arr.value[index]
}

func (arr IntArrayTag) Set(index int, value int32) {
	arr.value[index] = value
}

func (arr IntArrayTag) Length() int {
	return len(arr.value)
}

func NewIntArray(size int) *IntArrayTag {
	return &IntArrayTag{
		value: make([]int32, size),
	}
}
