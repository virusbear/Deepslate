package nbt

import "errors"

const ByteTypeId ByteType = 1

type ByteType int8

type ByteTag struct {
	value int8
}

func (_ ByteType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt8()

	if err != nil {
		return nil, err
	}

	return ByteTag{value: data}, nil
}

func (_ ByteType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteTag)
	if !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.WriteInt8(data.value)
}

func (_ ByteType) GetId() int8 {
	return int8(ByteTypeId)
}