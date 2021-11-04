package nbt

import "errors"

const ShortTypeId ShortType = 2

type ShortType int8

type ShortTag struct {
	value int16
}

func (_ ShortType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt16()

	if err != nil {
		return nil, err
	}

	return ShortTag{
		value: data,
	}, nil
}

func (_ ShortType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ShortTag)

	if !ok {
		return errors.New("incompatible tag. Expected SHORT")
	}

	return writer.WriteInt16(data.value)
}

func (_ ShortType) GetId() int8 {
	return int8(ShortTypeId)
}