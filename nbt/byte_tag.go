package nbt

import "errors"

type byteType struct {}

func (_ byteType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt8()

	if err != nil {
		return nil, err
	}

	return ByteTag{value: data}, nil
}

func (_ byteType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteTag)
	if !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.writeInt8(data.value)
}

type ByteTag struct {
	value int8
}

func (_ ByteTag) dataType() dataType {
	return byteType{}
}

func (_ ByteTag) Type() int8 {
	return TagByte
}

func (tag ByteTag) Get() int8 {
	return tag.value
}

func (tag ByteTag) Set(value int8) {
	tag.value = value
}

func NewByte(value int8) *ByteTag {
	return &ByteTag{
		value: value,
	}
}