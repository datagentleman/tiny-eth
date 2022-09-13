package rlp

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"reflect"
)

func Encode(values ...interface{}) []byte {
	encoding := []byte{}

	for _, v := range values {
		if isList(v) && !isByteArray(v) {
			encodedList := Encode(toList(v)...)
			firstByte := 0

			encodedList = bytes.TrimLeft(encodedList, "\x00")

			if len(encodedList) <= 55 {
				firstByte = 0xc0 + len(encodedList)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, encodedList...)
			} else {
				lenInBytes, _ := toBytes(len(encodedList))
				lenInBytes = bytes.TrimLeft(lenInBytes, "\x00")

				firstByte = 0xf7 + len(lenInBytes)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, lenInBytes...)
				encoding = append(encoding, encodedList...)
			}
		} else {
			buf, _ := toBytes(v)

			if !isByteArray(v) {
				buf = bytes.TrimLeft(buf, "\x00")
			}

			if len(buf) == 1 && buf[0] <= 0x7f {
				encoding = append(encoding, buf[0])
			} else if len(buf) <= 55 {
				firstByte := 0x80 + len(buf)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, buf...)
			} else {
				lenInBytes, _ := toBytes(len(buf))
				lenInBytes = bytes.TrimLeft(lenInBytes, "\x00")

				firstByte := 0xb7 + len(lenInBytes)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, lenInBytes...)
				encoding = append(encoding, buf...)
			}
		}
	}

	return encoding
}

func isList(v interface{}) bool {
	slice := reflect.TypeOf(v).Kind() == reflect.Slice
	array := reflect.TypeOf(v).Kind() == reflect.Array

	if isPointer(v) {
		ptr := reflect.TypeOf(v)
		return ptr.Elem().Kind() == reflect.Slice || ptr.Elem().Kind() == reflect.Array
	}

	return slice || array
}

func isPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Pointer
}

func isByteArray(v interface{}) bool {
	if isList(v) {
		typ := reflect.TypeOf(v)

		if isPointer(v) && typ.Elem().Elem().Kind() == reflect.Uint8 {
			return true
		}

		if typ.Elem().Kind() == reflect.Uint8 {
			return true
		}
	}

	return false
}

func toList(v interface{}) []interface{} {
	if !isList(v) {
		return nil
	}

	list := []interface{}{}
	ref := reflect.ValueOf(v)

	if ref.IsNil() {
		return list
	}

	if isPointer(v) {
		ref = ref.Elem()
	}

	for i := 0; i < ref.Len(); i++ {
		list = append(list, ref.Index(i).Interface())
	}

	return list
}

func toBytes(v interface{}) ([]byte, error) {
	if isPointer(v) && reflect.ValueOf(v).IsNil() {
		return []byte{}, nil
	}

	switch t := v.(type) {
	case big.Int:
		return t.Bytes(), nil
	case *big.Int:
		return t.Bytes(), nil
	case int:
		return toBytes(int64(t))
	case uint:
		return toBytes(uint64(t))
	case string:
		return []byte(t), nil
	default:
		var buf bytes.Buffer

		err := binary.Write(&buf, binary.BigEndian, v)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}
}
