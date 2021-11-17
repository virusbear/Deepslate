package nbt

import (
	"fmt"
)

type Tag interface{
	dataType() dataType
}

type dataType interface {
	Read(reader Reader) (Tag, error)
	Write(writer Writer, tag Tag) error
	GetId() int8
}

func getDataType(dtype int8) (dataType, error) {
	switch dtype {
	case int8(endTypeId):
		return endTypeId, nil
	case int8(byteTypeId):
		return byteTypeId, nil
	case int8(shortTypeId):
		return shortTypeId, nil
	case int8(intTypeId):
		return intTypeId, nil
	case int8(longTypeId):
		return longTypeId, nil
	case int8(floatTypeId):
		return floatTypeId, nil
	case int8(doubleTypeId):
		return doubleTypeId, nil
	case int8(byteArrayTypeId):
		return byteArrayTypeId, nil
	case int8(stringTypeId):
		return stringTypeId, nil
	case int8(listTypeId):
		return listTypeId, nil
	case int8(compoundTypeId):
		return compoundTypeId, nil
	case int8(intArrayTypeId):
		return intArrayTypeId, nil
	case int8(longArrayTypeId):
		return longArrayTypeId, nil
	default:
		return nil, fmt.Errorf("unknown NBT datatype %d", dtype)
	}
}
