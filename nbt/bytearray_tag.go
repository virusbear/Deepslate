package nbt

import "errors"

type byteArrayType struct{}

func (_ byteArrayType) Read(reader Reader) (Tag, error) {
	data, err := reader.readByteArray()

	if err != nil {
		return nil, err
	}

	return ByteArrayTag{
		value: data,
	}, nil
}

func (_ byteArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected BYTE_ARRAY")
	}

	return writer.writeByteArray(data.value)
}

type ByteArrayTag struct {
	value []int8
}

func (_ ByteArrayTag) dataType() dataType {
	return byteArrayType{}
}

func (_ ByteArrayTag) Type() int8 {
	return TagByteArray
}

func (arr ByteArrayTag) Raw() []int8 {
	return arr.value
}

func (arr ByteArrayTag) Get(index int) int8 {
	return arr.value[index]
}

func (arr ByteArrayTag) Set(index int, value int8) {
	arr.value[index] = value
}

func (arr ByteArrayTag) Length() int {
	return len(arr.value)
}

func NewByteArray(size int) *ByteArrayTag {
	return &ByteArrayTag{
		value: make([]int8, size),
	}
}