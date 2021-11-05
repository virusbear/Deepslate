package nbt

import "errors"

const compoundTypeId compoundType = 10

type compoundType int8

type CompoundTag struct {
	value []BaseTag
}

func (_ compoundType) Read(reader Reader) (Tag, error) {
	compound := CompoundTag{
		value: []BaseTag{},
	}

	for {
		tag, err := reader.Read()

		if err != nil {
			return nil, err
		}

		baseTag, ok := tag.(BaseTag)
		if !ok {
			return nil, errors.New("unable to read CompoundTag. Expected reader.Read() to return BASE_TAG")
		}

		if baseTag.dataType.GetId() == endTypeId.GetId() {
			break
		}

		compound.value = append(compound.value, baseTag)
	}

	return compound, nil
}

func (_ compoundType) Write(writer Writer, tag Tag) error {
	return errors.New("Not Implemented")
}

func (_ compoundType) GetId() int8 {
	return int8(compoundTypeId)
}