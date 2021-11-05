package nbt

import "errors"

const doubleTypeId doubleType = 6

type doubleType int8

type DoubleTag struct {
	value float64
}

func (_ doubleType) Read(reader Reader) (Tag, error) {
	data, err := reader.readFloat64()

	if err != nil {
		return nil, err
	}

	return DoubleTag{
		value: data,
	}, nil
}

func (_ doubleType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(DoubleTag)

	if !ok {
		return errors.New("incompatible tag. Expected DOUBLE")
	}

	return writer.writeFloat64(data.value)
}

func (_ doubleType) GetId() int8 {
	return int8(doubleTypeId)
}
