package nbt

import "errors"

const longTypeId longType = 4

type longType int8

func (_ longType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt64()

	if err != nil {
		return nil, err
	}

	return LongTag{
		value: data,
	}, nil
}

func (_ longType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONG")
	}

	return writer.writeInt64(data.value)
}

func (_ longType) GetId() int8 {
	return int8(longTypeId)
}

type LongTag struct {
	value int64
}

func (_ LongTag) dataType() dataType {
	return longTypeId
}