package nbt

import (
	"errors"
)

const compoundTypeId compoundType = 10

type compoundType int8

func (_ compoundType) Read(reader Reader) (Tag, error) {
	compound := CompoundTag{
		tags: map[string]Tag{},
	}

	for {
		name, tag, err := reader.Read()

		if err != nil {
			return nil, err
		}

		if tag.dataType() == endTypeId {
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

	err := writer.writeInt8(endTypeId.GetId())
	if err != nil {
		return err
	}

	err = endTypeId.Write(writer, endTag{})
	return err
}

func (_ compoundType) GetId() int8 {
	return int8(compoundTypeId)
}

type CompoundTag struct {
	tags map[string]Tag
}

func (_ CompoundTag) dataType() dataType {
	return compoundTypeId
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