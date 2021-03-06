package nbt

import (
	"errors"
)

type compoundType struct{}

func (_ compoundType) Read(reader Reader) (Tag, error) {
	compound := CompoundTag{
		tags: map[string]Tag{},
	}

	for {
		name, tag, err := reader.Read()

		if err != nil {
			return nil, err
		}

		if tag.Type() == TagEnd {
			break
		}

		compound.tags[name] = tag
	}

	return compound, nil
}

func (_ compoundType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(CompoundTag)

	if !ok {
		return errors.New("incompatible tag. Expected COMPOUND")
	}

	for name, value := range data.tags {
		err := writer.Write(name, value)

		if err != nil {
			return err
		}
	}

	err := writer.writeInt8(TagEnd)
	if err != nil {
		return err
	}

	end := endTag{}
	err = end.dataType().Write(writer, end)
	return err
}

type CompoundTag struct {
	tags map[string]Tag
}

func (_ CompoundTag) dataType() dataType {
	return compoundType{}
}

func (_ CompoundTag) Type() int8 {
	return TagCompound
}

func (tag CompoundTag) GetTag(name string) Tag {
	return tag.tags[name]
}

func (t CompoundTag) SetTag(name string, tag Tag) {
	t.tags[name] = tag
}

func (tag CompoundTag) GetByte(name string) int8 {
	result, ok := tag.GetTag(name).(ByteTag)

	if !ok {
		panic("unable to get byte")
	}

	return result.Get()
}

func (tag CompoundTag) SetByte(name string, value int8) {
	tag.SetTag(name, NewByte(value))
}

func (tag CompoundTag) GetByteArray(name string) []int8 {
	result, ok := tag.GetTag(name).(ByteArrayTag)

	if !ok {
		panic("unable to get bytearray")
	}

	return result.value
}

func (tag CompoundTag) SetByteArray(name string, value []int8) {
	arr := NewByteArray(0)
	arr.value = value
	tag.SetTag(name, arr)
}

func (tag CompoundTag) GetCompound(name string) *CompoundTag {
	result, ok := tag.GetTag(name).(CompoundTag)

	if !ok {
		panic("unable to get compound")
	}

	return &result
}

func (tag CompoundTag) SetCompound(name string, compound *CompoundTag) {
	tag.SetTag(name, compound)
}

func (tag CompoundTag) GetDouble(name string) float64 {
	result, ok := tag.GetTag(name).(DoubleTag)

	if !ok {
		panic("unable to get double")
	}

	return result.Get()
}

func (tag CompoundTag) SetDouble(name string, value float64) {
	tag.SetTag(name, NewDouble(value))
}

func (tag CompoundTag) GetFloat(name string) float32 {
	result, ok := tag.GetTag(name).(FloatTag)

	if !ok {
		panic("unable to get float")
	}

	return result.Get()
}

func (tag CompoundTag) SetFloat(name string, value float32) {
	tag.SetTag(name, NewFloat(value))
}

func (tag CompoundTag) GetInt(name string) int32 {
	result, ok := tag.GetTag(name).(IntTag)

	if !ok {
		panic("unable to get int")
	}

	return result.Get()
}

func (tag CompoundTag) SetInt(name string, value int32) {
	tag.SetTag(name, NewInt(value))
}

func (tag CompoundTag) GetIntArray(name string) []int32 {
	result, ok := tag.GetTag(name).(IntArrayTag)

	if !ok {
		panic("unable to get intarray")
	}

	return result.value
}

func (tag CompoundTag) SetIntArray(name string, value []int32) {
	arr := NewIntArray(0)
	arr.value = value
	tag.SetTag(name, arr)
}

func (tag CompoundTag) GetLong(name string) int64 {
	result, ok := tag.GetTag(name).(LongTag)

	if !ok {
		panic("unable to get long")
	}

	return result.Get()
}

func (tag CompoundTag) SetLong(name string, value int64) {
	tag.SetTag(name, NewLong(value))
}

func (tag CompoundTag) GetLongArray(name string) []int64 {
	result, ok := tag.GetTag(name).(LongArrayTag)

	if !ok {
		panic("unable to get longarray")
	}

	return result.value
}

func (tag CompoundTag) SetLongArray(name string, value []int64) {
	arr := NewLongArray(0)
	arr.value = value
	tag.SetTag(name, arr)
}

func (tag CompoundTag) GetShort(name string) int16 {
	result, ok := tag.GetTag(name).(ShortTag)

	if !ok {
		panic("unable to get short")
	}

	return result.Get()
}

func (tag CompoundTag) SetShort(name string, value int16) {
	tag.SetTag(name, NewShort(value))
}

func (tag CompoundTag) GetString(name string) string {
	result, ok := tag.GetTag(name).(StringTag)

	if !ok {
		panic("unable to get string")
	}

	return result.Get()
}

func (tag CompoundTag) SetString(name string, value string) {
	tag.SetTag(name, NewString(value))
}

func (tag CompoundTag) SetList(name string, value *ListTag) {
	tag.SetTag(name, value)
}

func (tag CompoundTag) GetList(name string) *ListTag {
	result, ok := tag.GetTag(name).(ListTag)

	if !ok {
		panic("unable to get short")
	}

	return &result
}
