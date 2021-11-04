package nbt

import (
	"errors"
	"fmt"
)

const IntArrayTypeId  IntArrayType  = 11

type IntArrayType int8

type IntArrayTag struct {
	value []int32
}

func (_ IntArrayType) Read(reader Reader) (Tag, error) {
	length, err := reader.ReadInt32()

	if err != nil {
		return nil, err
	}

	data := make([]int32, length)

	for i, _ := range data {
		value, err := reader.ReadInt32()

		if err != nil {
			return nil, fmt.Errorf("unable to read int array at index %d. Reason: %w", i, err)
		}

		data[i] = value
	}

	return data, nil
}

func (_ IntArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(IntArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected INTARRAY")
	}

	if err := writer.WriteInt32(int32(len(data.value))); err != nil {
		return err
	}

	for i, value := range data.value {
		if err := writer.WriteInt32(value); err != nil {
			return fmt.Errorf("unable to write int array at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

func (_ IntArrayType) GetId() int8 {
	return int8(IntArrayTypeId)
}