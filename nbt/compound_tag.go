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

		if baseTag.dataType == endTypeId {
			break
		}

		compound.value = append(compound.value, baseTag)
	}

	return compound, nil
}

func (_ compoundType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(CompoundTag)

	if !ok {
		return errors.New("incompatible tag. Expected COMPOUND")
	}

	for _, value := range data.value {
		err := writer.Write(value)

		if err != nil {
			return err
		}
	}

	return writer.Write(
		BaseTag{
			name: "",
			dataType: endTypeId,
			tag: EndTag{},
		},
	)
}

func (_ compoundType) GetId() int8 {
	return int8(compoundTypeId)
}