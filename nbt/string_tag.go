package nbt

import "errors"

const stringTypeId stringType = 8

type stringType int8

func (_ stringType) Read(reader Reader) (Tag, error) {
	data, err := reader.readString()

	if err != nil {
		return nil, err
	}

	return StringTag{
		value: data,
	}, nil
}

func (_ stringType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(StringTag)

	if !ok {
		return errors.New("incompatible tag. Expected STRING")
	}

	return writer.writeString(data.value)
}

func (_ stringType) GetId() int8 {
	return int8(stringTypeId)
}

type StringTag struct {
	value string
}

func (_ StringTag) dataType() dataType {
	return stringTypeId
}