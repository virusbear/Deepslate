package nbt

import "errors"

const FloatTypeId FloatType = 5

type FloatType int8

type FloatTag struct {
	value float32
}

func (_ FloatType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	return FloatTag{
		value: data,
	}, nil
}

func (_ FloatType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(FloatTag)

	if !ok {
		return errors.New("incompatible tag. Expected FLOAT")
	}

	return writer.WriteFloat32(data.value)
}

func (_ FloatType) GetId() int8 {
	return int8(FloatTypeId)
}
