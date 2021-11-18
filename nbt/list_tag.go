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
		dType: dataType,
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
		if err := data.dType.Write(writer, value); err != nil {
			return fmt.Errorf("unable to write list at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

func (_ listType) GetId() int8 {
	return int8(listTypeId)
}

type ListTag struct {
	dType dataType
	value    []Tag
}

func (_ ListTag) dataType() dataType {
	return listTypeId
}