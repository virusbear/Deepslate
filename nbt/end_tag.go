package nbt

import "errors"

const endTypeId endType = 0

type endType int8

func (end endType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt8()

	if err != nil {
		return nil, err
	}

	if data != 0 {
		return nil, errors.New("invalid end tag")
	}

	return endTag{}, nil
}

func (_ endType) Write(writer Writer, tag Tag) error {
	if _, ok := tag.(endTag); !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.writeInt8(0)
}

func (_ endType) GetId() int8 {
	return int8(endTypeId)
}

type endTag struct{}

func (_ endTag) dataType() dataType {
	return endTypeId
}