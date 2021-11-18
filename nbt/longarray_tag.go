package nbt

import (
	"errors"
	"fmt"
)

type longArrayType struct{}

func (_ longArrayType) Read(reader Reader) (Tag, error) {
	length, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	data := make([]int64, length)

	for i, _ := range data {
		value, err := reader.readInt64()

		if err != nil {
			return nil, fmt.Errorf("unable to read long array at index %d. Reason: %w", i, err)
		}

		data[i] = value
	}

	return LongArrayTag{
		value: data,
	}, nil
}

func (_ longArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONGARRAY")
	}

	if err := writer.writeInt32(int32(len(data.value))); err != nil {
		return err
	}

	for i, value := range data.value {
		if err := writer.writeInt64(value); err != nil {
			return fmt.Errorf("unable to write long array at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

type LongArrayTag struct {
	value []int64
}

func (_ LongArrayTag) dataType() dataType {
	return longArrayType{}
}

func (_ LongArrayTag) Type() int8 {
	return TagLongArray
}

func (arr LongArrayTag) Raw() []int64 {
	return arr.value
}

func (arr LongArrayTag) Get(index int) int64 {
	return arr.value[index]
}

func (arr LongArrayTag) Set(index int, value int64) {
	arr.value[index] = value
}

func (arr LongArrayTag) Length() int {
	return len(arr.value)
}

func NewLongArray(size int) *LongArrayTag {
	return &LongArrayTag{
		value: make([]int64, size),
	}
}
