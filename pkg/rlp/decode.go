package rlp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func Decode(encodings []byte, values ...interface{}) (int, error) {
	enc := bytes.NewBuffer(encodings)
	var encoding []byte

	for _, v := range values {
		if isList(v) && !isByteArray(v) {
			// Decode list
			decoder := NewDecoder(encodings)
			list := decoder.Next()
			list.Decode(v)
		} else {
			// Decode string
			switch v := v.(type) {
			case *interface{}:
				decodeInterface(v, enc, encodings)

			case *bool:
				_, encoding = nextEncoding(enc)
				*v = len(encoding) > 0

			case *uint8:
				_, encoding = nextEncoding(enc)
				*v = uint8(encoding[0])

			case *uint16:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 2)
				*v = binary.BigEndian.Uint16(encoding)

			case *uint32:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 4)
				*v = binary.BigEndian.Uint32(encoding)

			case *uint64:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 8)
				*v = binary.BigEndian.Uint64(encoding)

			case *int8:
				_, encoding = nextEncoding(enc)
				*v = int8(encoding[0])

			case *int16:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 2)
				*v = (int16)(binary.BigEndian.Uint16(encoding))

			case *int32:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 4)
				*v = (int32)(binary.BigEndian.Uint32(encoding))

			case *int64:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 8)
				*v = (int64)(binary.BigEndian.Uint64(encoding))

			case *float32:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 4)
				buf := bytes.NewReader(encoding)
				binary.Read(buf, binary.BigEndian, v)

			case *float64:
				_, encoding = nextEncoding(enc)
				ensureLen(&encoding, 8)
				buf := bytes.NewReader(encoding)
				binary.Read(buf, binary.BigEndian, v)

			case *string:
				_, encoding = nextEncoding(enc)
				*v = string(encoding)

			case *[]uint8:
				_, encoding = nextEncoding(enc)
				*v = append(*v, encoding...)

			default:
				fmt.Println("Unknown type:")
				fmt.Println(reflect.TypeOf(v))
			}
		}
	}

	return len(encoding), nil
}

func decodeInterface(v interface{}, enc *bytes.Buffer, encodings []byte) {
	switch v := v.(type) {
	case *interface{}:
		if isList(*v) {
			switch reflect.TypeOf(*v).Elem().Kind() {

			case reflect.Uint8:
				tmp := []byte{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int8:
				tmp := []int8{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int16:
				tmp := []int16{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int32:
				tmp := []int32{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int64:
				tmp := []int64{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint16:
				tmp := []uint16{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint32:
				tmp := []uint32{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint64:
				tmp := []uint64{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Float32:
				tmp := []float32{}
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Float64:
				tmp := []float64{}
				Decode(encodings, &tmp)
				*v = tmp
			}
		} else {
			a := reflect.TypeOf(*v)

			switch a.Kind() {
			case reflect.Bool:
				var tmp bool
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint8:
				var tmp uint8
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint16:
				var tmp uint16
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint32:
				var tmp uint32
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Uint64:
				var tmp uint64
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int8:
				var tmp int8
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int16:
				var tmp int16
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int32:
				var tmp int32
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Int64:
				var tmp int64
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Float32:
				var tmp float32
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.Float64:
				var tmp float64
				Decode(encodings, &tmp)
				*v = tmp

			case reflect.String:
				var tmp string
				Decode(encodings, &tmp)
				*v = tmp

			default:
				fmt.Println("Unknown type:")
				fmt.Println(a)
			}
		}
	}
}

func nextEncoding(encodings *bytes.Buffer) (prefix, encoding []byte) {
	firstByte := encodings.Next(1)
	size := firstByte[0]

	if size <= 0x7f {
		return nil, firstByte
	}

	if size <= 0xb7 {
		len := int(size - 0x80)
		return firstByte, encodings.Next(len)
	}

	if size <= 0xbf {
		len := int(size - 0xb7)
		buf := encodings.Next(len)
		ensureLen(&buf, 8)

		size := (int)(binary.BigEndian.Uint64(buf))
		return firstByte, encodings.Next(size)
	}

	if size <= 0xf7 {
		len := int(size - 0xc0)
		return firstByte, encodings.Next(len)
	}

	fmt.Println("unknown encoding type")
	return nil, nil
}

func ensureLen(buf *[]byte, length int) {
	bufLen := len(*buf)

	if bufLen < length {
		l := length - bufLen
		b := make([]byte, length)

		*buf = append(b[:l], *buf...)
	}
}

func isEmptyList(v interface{}) bool {
	if !isList(v) {
		return false
	}

	return reflect.ValueOf(v).Len() <= 0
}
