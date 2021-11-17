package nbt

import (
	"errors"
	"fmt"
)

const longArrayTypeId longArrayType = 12

type longArrayType int8

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

func (_ longArrayType) GetId() int8 {
	return int8(longArrayTypeId)
}

type LongArrayTag struct {
	value []int64
}

func (_ LongArrayTag) dataType() dataType {
	return longArrayTypeId
}