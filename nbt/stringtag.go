package nbt

import "errors"

const StringTypeId StringType = 8

type StringType int8

type StringTag struct {
	value string
}

func (_ StringType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadString()

	if err != nil {
		return nil, err
	}

	return StringTag{
		value: data,
	}, nil
}

func (_ StringType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(StringTag)

	if !ok {
		return errors.New("incompatible tag. Expected STRING")
	}

	return writer.WriteString(data.value)
}

func (_ StringType) GetId() int8 {
	return int8(StringTypeId)
}
