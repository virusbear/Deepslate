package nbt

import "errors"

const ByteArrayTypeId ByteArrayType = 7

type ByteArrayType int8

type ByteArrayTag struct {
	value []int8
}

func (_ ByteArrayType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadByteArray()

	if err != nil {
		return nil, err
	}

	return ByteArrayTag{
		value: data,
	}, nil
}

func (_ ByteArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected BYTE_ARRAY")
	}

	return writer.WriteByteArray(data.value)
}

func (_ ByteArrayType) GetId() int8 {
	return int8(ByteArrayTypeId)
}
