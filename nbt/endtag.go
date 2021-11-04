package nbt

import "errors"

const EndTypeId EndType = 0

type EndType int8

type EndTag struct{}

func (end EndType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt8()

	if err != nil {
		return nil, err
	}

	if data != 0 {
		return nil, errors.New("invalid end tag")
	}

	return EndTag{}, nil
}

func (_ EndType) Write(writer Writer, tag Tag) error {
	if _, ok := tag.(EndTag); !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.WriteInt8(0)
}

func (_ EndType) GetId() int8 {
	return int8(EndTypeId)
}