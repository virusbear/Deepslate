package nbt

import "errors"

type longType struct{}

func (_ longType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt64()

	if err != nil {
		return nil, err
	}

	return LongTag{
		value: data,
	}, nil
}

func (_ longType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONG")
	}

	return writer.writeInt64(data.value)
}

type LongTag struct {
	value int64
}

func (_ LongTag) dataType() dataType {
	return longType{}
}

func (_ LongTag) Type() int8 {
	return TagLong
}

func (tag LongTag) Get() int64 {
	return tag.value
}

func (tag LongTag) Set(value int64) {
	tag.value = value
}

func NewLong(value int64) *LongTag {
	return &LongTag{
		value: value,
	}
}
