package nbt

import (
	"errors"
	"fmt"
)

const intArrayTypeId intArrayType = 11

type intArrayType int8

type IntArrayTag struct {
	value []int32
}

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

	return data, nil
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

func (_ intArrayType) GetId() int8 {
	return int8(intArrayTypeId)
}