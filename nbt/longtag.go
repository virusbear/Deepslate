package nbt

import "errors"

const LongTypeId LongType = 4

type LongType int8

type LongTag struct {
	value int64
}

func (_ LongType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt64()

	if err != nil {
		return nil, err
	}

	return LongTag{
		value: data,
	}, nil
}

func (_ LongType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONG")
	}

	return writer.WriteInt64(data.value)
}

func (_ LongType) GetId() int8 {
	return int8(LongTypeId)
}