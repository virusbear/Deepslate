package nbt

import "errors"

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
}

type EndType int8
type ByteType int8
type ShortType int8
type IntType int8
type LongType int8
type FloatType int8
type DoubleType int8
type ByteArrayType int8
type StringType int8
type ListType int8
type CompoundType int8
type IntArrayType int8
type LongArrayType int8

type EndTag struct{}
type ByteTag struct {
	value int8
}
type ShortTag struct {
	value int16
}
type IntTag struct {
	value int32
}
type LongTag struct {
	value int64
}
type FloatTag struct {
	value float32
}
type DoubleTag struct {
	value float64
}
type ByteArrayTag struct {
	value []byte
}
type StringTag struct {
	value string
}
type ListTag struct {
	dataType DataType
	value    []Tag
}
type CompoundTag struct {
	value []BaseTag
}
type IntArrayTag struct {
	value []int32
}
type LongArrayTag struct {
	value []int64
}

func (end EndType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt8()

	if err != nil {
		return nil, err
	}

	if data != 0 {
		return nil, errors.New("invalid end tag")
	}

	return EndTag{}, nil
}

func (_ EndType) Write(writer Writer, tag Tag) error {
	if _, ok := tag.(EndTag); !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.WriteInt8(0)
}

func (_ ByteType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt8()

	if err != nil {
		return nil, err
	}

	return ByteTag{value: data}, nil
}

func (_ ByteType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteTag)
	if !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.WriteInt8(data.value)
}

func (_ ShortType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt16()

	if err != nil {
		return nil, err
	}

	return ShortTag{
		value: data,
	}, nil
}

func (_ ShortType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ShortTag)

	if !ok {
		return errors.New("incompatible tag. Expected SHORT")
	}

	return writer.WriteInt16(data.value)
}

const (
	EndTypeId       EndType       = 0
	ByteTypeId      ByteType      = 1
	ShortTypeId     ShortType     = 2
	IntTypeId       IntType       = 3
	LongTypeId      LongType      = 4
	FloatTypeId     FloatType     = 5
	DoubleTypeId    DoubleType    = 6
	ByteArrayTypeId ByteArrayType = 7
	StringTypeId    StringType    = 8
	ListTypeId      ListType      = 9
	CompoundTypeId  CompoundType  = 10
	IntArrayTypeId  IntArrayType  = 11
	LongArrayTypeId LongArrayType = 12
)

func getDataType(dtype int8) DataType {
	switch dtype {
	case int8(EndTypeId):
		return EndTypeId
	case int8(ByteTypeId):
		return ByteTypeId
	case int8(ShortTypeId):
		return ShortTypeId
	default:
		return nil
	}
}
