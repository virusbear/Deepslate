package nbt

import "errors"

const floatTypeId floatType = 5

type floatType int8

type FloatTag struct {
	value float32
}

func (_ floatType) Read(reader Reader) (Tag, error) {
	data, err := reader.readFloat32()

	if err != nil {
		return nil, err
	}

	return FloatTag{
		value: data,
	}, nil
}

func (_ floatType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(FloatTag)

	if !ok {
		return errors.New("incompatible tag. Expected FLOAT")
	}

	return writer.writeFloat32(data.value)
}

func (_ floatType) GetId() int8 {
	return int8(floatTypeId)
}
