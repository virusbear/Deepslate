package nbt

import "errors"

const DoubleTypeId DoubleType = 6

type DoubleType int8

type DoubleTag struct {
	value float64
}

func (_ DoubleType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadFloat64()

	if err != nil {
		return nil, err
	}

	return DoubleTag{
		value: data,
	}, nil
}

func (_ DoubleType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(DoubleTag)

	if !ok {
		return errors.New("incompatible tag. Expected DOUBLE")
	}

	return writer.WriteFloat64(data.value)
}

func (_ DoubleType) GetId() int8 {
	return int8(DoubleTypeId)
}
