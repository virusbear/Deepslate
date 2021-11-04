package nbt

import (
	"errors"
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

type EndTag struct {
}
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
	value []int8
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

func (_ IntType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt32()

	if err != nil {
		return nil, err
	}

	return IntTag{
		value: data,
	}, nil
}

func (_ IntType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(IntTag)

	if !ok {
		return errors.New("incompatible tag. Expected INT")
	}

	return writer.WriteInt32(data.value)
}

func (_ LongType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadInt64()

	if err != nil {
		return nil, err
	}

	return LongTag{
		value: data,
	}, nil
}

func (_ LongType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONG")
	}

	return writer.WriteInt64(data.value)
}

func (_ FloatType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	return FloatTag{
		value: data,
	}, nil
}

func (_ FloatType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(FloatTag)

	if !ok {
		return errors.New("incompatible tag. Expected FLOAT")
	}

	return writer.WriteFloat32(data.value)
}

func (_ DoubleType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadFloat64()

	if err != nil {
		return nil, err
	}

	return DoubleTag{
		value: data,
	}, nil
}

func (_ DoubleType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(DoubleTag)

	if !ok {
		return errors.New("incompatible tag. Expected DOUBLE")
	}

	return writer.WriteFloat64(data.value)
}

func (_ ByteArrayType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadByteArray()

	if err != nil {
		return nil, err
	}

	return ByteArrayTag{
		value: data,
	}, nil
}

func (_ ByteArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected BYTE_ARRAY")
	}

	return writer.WriteByteArray(data.value)
}

func (_ StringType) Read(reader Reader) (Tag, error) {
	data, err := reader.ReadString()

	if err != nil {
		return nil, err
	}

	return StringTag{
		value: data,
	}, nil
}

func (_ StringType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(StringTag)

	if !ok {
		return errors.New("incompatible tag. Expected STRING")
	}

	return writer.WriteString(data.value)
}

func (_ ListType) Read(reader Reader) (Tag, error) {
	dtype, err := reader.ReadInt8()

	if err != nil {
		return nil, err
	}

	dataType, err := getDataType(dtype)

	if err != nil {
		return nil, err
	}

	length, err := reader.ReadInt32()

	if err != nil {
		return nil, err
	}

	list := make([]Tag, length)

	for i, _ := range list {
		data, err := dataType.Read(reader)

		if err != nil {
			return nil, fmt.Errorf("unable to read list at index %d. Reason: %w", i, err)
		}

		list[i] = data
	}

	return list, nil
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
