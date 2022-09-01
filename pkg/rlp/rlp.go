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

			if len(bytes.TrimLeft(encodedList, "\x00")) <= 55 {
				firstByte = 0xc0 + len(bytes.TrimLeft(encodedList, "\x00"))

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, bytes.TrimLeft(encodedList, "\x00")...)
			} else {
				lenInBytes, _ := toBytes(len(encodedList))
				lenInBytes = bytes.TrimLeft(lenInBytes, "\x00")

				firstByte = 0xf7 + len(lenInBytes)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, lenInBytes...)
				encoding = append(encoding, bytes.TrimLeft(encodedList, "\x00")...)
			}
		} else {
			bs, _ := toBytes(v)

			if !isByteArray(v) {
				bs = bytes.TrimLeft(bs, "\x00")
			}

			if len(bs) == 1 && bs[0] <= 0x7f {
				encoding = append(encoding, bs[0])
			} else if len(bs) <= 55 {
				firstByte := 0x80 + len(bs)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, bs...)
			} else {
				lenInBytes, _ := toBytes(len(bs))
				lenInBytes = bytes.TrimLeft(lenInBytes, "\x00")

				firstByte := 0xb7 + len(lenInBytes)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, lenInBytes...)
				encoding = append(encoding, bs...)
			}
		}
	}

	return encoding
}

func isList(v interface{}) bool {
	slice := reflect.TypeOf(v).Kind() == reflect.Slice
	array := reflect.TypeOf(v).Kind() == reflect.Array

	pttr := reflect.TypeOf(v)
	if pttr.Kind() == reflect.Pointer {
		aa := pttr.Elem().Kind() == reflect.Slice || pttr.Elem().Kind() == reflect.Array
		return aa
	}

	return slice || array
}

func isPointer(v interface{}) bool {
	pttr := reflect.TypeOf(v)
	return pttr.Kind() == reflect.Pointer
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

	s := []interface{}{}
	l := reflect.ValueOf(v)

	ll := reflect.ValueOf(v)
	if ll.IsNil() {
		return s
	}

	for i := 0; i < l.Len(); i++ {
		s = append(s, l.Index(i).Interface())
	}

	return s
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
