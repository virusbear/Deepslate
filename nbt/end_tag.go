package nbt

import "errors"

type endType struct{}

func (end endType) Read(reader Reader) (Tag, error) {
	return endTag{}, nil
}

func (_ endType) Write(writer Writer, tag Tag) error {
	if _, ok := tag.(endTag); !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return nil
}

type endTag struct{}

func (_ endTag) dataType() dataType {
	return endType{}
}

func (_ endTag) Type() int8 {
	return TagEnd
}
