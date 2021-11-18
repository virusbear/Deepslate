package nbt

import "errors"

type shortType struct{}

func (_ shortType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt16()

	if err != nil {
		return nil, err
	}

	return ShortTag{
		value: data,
	}, nil
}

func (_ shortType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ShortTag)

	if !ok {
		return errors.New("incompatible tag. Expected SHORT")
	}

	return writer.writeInt16(data.value)
}

type ShortTag struct {
	value int16
}

func (_ ShortTag) dataType() dataType {
	return shortType{}
}

func (_ ShortTag) Type() int8 {
	return TagShort
}

func (tag ShortTag) Get() int16 {
	return tag.value
}

func (tag ShortTag) Set(value int16) {
	tag.value = value
}

func NewShort(value int16) *ShortTag {
	return &ShortTag{
		value: value,
	}
}
