package nbt

import "errors"

type doubleType struct{}

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

type DoubleTag struct {
	value float64
}

func (_ DoubleTag) dataType() dataType {
	return doubleTypeId
}

func (tag DoubleTag) Get() float64 {
	return tag.value
}

func (tag DoubleTag) Set(value float64) {
	tag.value = value
}

func NewDouble(value float64) *DoubleTag {
	return &DoubleTag{
		value: value,
	}
}