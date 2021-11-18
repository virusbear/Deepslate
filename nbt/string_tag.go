package nbt

import "errors"

type stringType struct{}

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

type StringTag struct {
	value string
}

func (_ StringTag) dataType() dataType {
	return stringType{}
}

func (_ StringTag) Type() int8 {
	return TagString
}

func (tag StringTag) Get() string {
	return tag.value
}

func (tag StringTag) Set(value string) {
	tag.value = value
}

func NewString(value string) *StringTag {
	return &StringTag{
		value: value,
	}
}
