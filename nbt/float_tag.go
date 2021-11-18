package nbt

import "errors"

type floatType struct{}

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

type FloatTag struct {
	value float32
}

func (_ FloatTag) dataType() dataType {
	return floatTypeId
}

func (tag FloatTag) Get() float32 {
	return tag.value
}

func (tag FloatTag) Set(value float32) {
	tag.value = value
}

func NewFloat(value float32) *FloatTag {
	return &FloatTag{
		value: value,
	}
}