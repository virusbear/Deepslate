package nbt

import (
	"errors"
	"fmt"
)

const LongArrayTypeId LongArrayType = 12

type LongArrayType int8

type LongArrayTag struct {
	value []int64
}

func (_ LongArrayType) Read(reader Reader) (Tag, error) {
	length, err := reader.ReadInt32()

	if err != nil {
		return nil, err
	}

	data := make([]int64, length)

	for i, _ := range data {
		value, err := reader.ReadInt64()

		if err != nil {
			return nil, fmt.Errorf("unable to read long array at index %d. Reason: %w", i, err)
		}

		data[i] = value
	}

	return data, nil
}

func (_ LongArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONGARRAY")
	}

	if err := writer.WriteInt32(int32(len(data.value))); err != nil {
		return err
	}

	for i, value := range data.value {
		if err := writer.WriteInt64(value); err != nil {
			return fmt.Errorf("unable to write long array at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

func (_ LongArrayType) GetId() int8 {
	return int8(LongArrayTypeId)
}