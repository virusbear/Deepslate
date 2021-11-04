package nbt

import (
	"errors"
	"fmt"
)

const ListTypeId ListType = 9

type ListType int8

type ListTag struct {
	dataType DataType
	value    []Tag
}

func (_ ListType) Read(reader Reader) (Tag, error) {
	dtype, err := reader.ReadInt8()

	if err != nil {
		return nil, err
	}

	dataType, err := getDataType(dtype)

	if err != nil {
		return nil, err
	}

	length, err := reader.ReadInt32()

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

	return list, nil
}

func (_ ListType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ListTag)

	if !ok {
		return errors.New("incompatible tag. Expected LIST")
	}

	if err := writer.WriteInt8(data.dataType.GetId()); err != nil {
		return fmt.Errorf("unable to write list datatype. Reason: %w", err)
	}

	for i, value := range data.value {
		if err := data.dataType.Write(writer, value); err != nil {
			return fmt.Errorf("unable to write list at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

func (_ ListType) GetId() int8 {
	return int8(ListTypeId)
}
