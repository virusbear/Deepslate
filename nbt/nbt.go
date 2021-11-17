package nbt

import (
	"io"
)


func Read(reader io.Reader) (Tag, error) {
	return NewReader(reader).Read()
}

func Write(writer io.Writer, tag Tag) error {
	return NewWriter(writer).Write(tag)
}

// NewByteTag TODO: Add Tag "Constructors"
func NewByteTag(name string, value int8) Tag {
	return BaseTag{
		name: name,
		dataType: byteTypeId,
		tag: ByteTag{
			value: value,
		},
	}
}