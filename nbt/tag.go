package nbt

import (
	"fmt"
)

const (
	TagEnd int8 = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)

type Tag interface{
	dataType() dataType
	Type() int8
}

type dataType interface {
	Read(reader Reader) (Tag, error)
	Write(writer Writer, tag Tag) error
}

func getDataType(dtype int8) (dataType, error) {
	switch dtype {
	case TagEnd:
		return endType{}, nil
	case TagByte:
		return byteType{}, nil
	case TagShort:
		return shortType{}, nil
	case TagInt:
		return intType{}, nil
	case TagLong:
		return longType{}, nil
	case TagFloat:
		return floatType{}, nil
	case TagDouble:
		return doubleType{}, nil
	case TagByteArray:
		return byteArrayType{}, nil
	case TagString:
		return stringType{}, nil
	case TagList:
		return listType{}, nil
	case TagCompound:
		return compoundType{}, nil
	case TagIntArray:
		return intArrayType{}, nil
	case TagLongArray:
		return longArrayType{}, nil
	default:
		return nil, fmt.Errorf("unknown NBT datatype %d", dtype)
	}
}
