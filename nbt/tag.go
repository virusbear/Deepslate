package nbt

import (
	"fmt"
)

type Tag interface{}

type Wrapper interface {
	GetTag() Tag
}

type Named interface {
	GetName() string
}

type Typed interface {
	GetDataType() DataType
}

type BaseTag struct {
	name     string
	dataType DataType
	tag      Tag
}

func (tag BaseTag) GetName() string {
	return tag.name
}

func (tag BaseTag) GetDataType() DataType {
	return tag.dataType
}

func (tag BaseTag) GetTag() Tag {
	return tag.tag
}

type DataType interface {
	Read(reader Reader) (Tag, error)
	Write(writer Writer, tag Tag) error
	GetId() int8
}

func getDataType(dtype int8) (DataType, error) {
	switch dtype {
	case int8(EndTypeId):
		return EndTypeId, nil
	case int8(ByteTypeId):
		return ByteTypeId, nil
	case int8(ShortTypeId):
		return ShortTypeId, nil
	case int8(IntTypeId):
		return IntTypeId, nil
	case int8(LongTypeId):
		return LongTypeId, nil
	case int8(FloatTypeId):
		return FloatTypeId, nil
	case int8(DoubleTypeId):
		return DoubleTypeId, nil
	case int8(ByteArrayTypeId):
		return ByteArrayTypeId, nil
	case int8(StringTypeId):
		return StringTypeId, nil
	case int8(ListTypeId):
		return ListTypeId, nil
	case int8(CompoundTypeId):
		return CompoundTypeId, nil
	case int8(IntArrayTypeId):
		return IntArrayTypeId, nil
	case int8(LongArrayTypeId):
		return LongArrayTypeId, nil
	default:
		return nil, fmt.Errorf("unknown NBT datatype %d", dtype)
	}
}
