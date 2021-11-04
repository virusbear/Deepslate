package nbt

import "errors"

const CompoundTypeId  CompoundType  = 10

type CompoundType int8

type CompoundTag struct {
	value []BaseTag
}

func (_ CompoundType) Read(reader Reader) (Tag, error) {
	return nil, errors.New("Not Implemented")
}

func (_ CompoundType) Write(writer Writer, tag Tag) error {
	return errors.New("Not Implemented")
}

func (_ CompoundType) GetId() int8 {
	return int8(CompoundTypeId)
}