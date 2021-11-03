package nbt

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

func Unmarshal(reader io.Reader, v interface{}) error {
	for {
		dataType, err := readByte(reader)
		if err != nil {
			return err
		}

		if DataType(dataType) == End {
			return nil
		}

		name, err := readString(reader)
		if err != nil {
			return err
		}

		field, found := reflect.TypeOf(v).
		if !found {
			return fmt.Errorf("field %s not found", name)
		}

		field.Tag.Get("nbt")

		switch DataType(dataType) {
			case End: return nil
			case Byte: {
				readByte(reader)
			}

		}
	}
}

func Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}