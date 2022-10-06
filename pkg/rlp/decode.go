package rlp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
)

func (d *Decoder) Decode(values ...interface{}) (int, error) {
	var encoding []byte

	for _, v := range values {
		if isList(v) {
			d.DecodeList(v)
			continue
		}

		// String
		switch v := v.(type) {

		case *interface{}:
			d.decodeInterface(v)

		case *bool:
			encoding = d.nextEncoding()
			*v = len(encoding) > 0

		case *uint8:
			encoding = d.nextEncoding()
			*v = uint8(encoding[0])

		case *uint16:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 2)
			*v = binary.BigEndian.Uint16(encoding)

		case *uint32:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 4)
			*v = binary.BigEndian.Uint32(encoding)

		case *uint64:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 8)
			*v = binary.BigEndian.Uint64(encoding)

		case *int8:
			encoding = d.nextEncoding()
			*v = int8(encoding[0])

		case *int16:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 2)
			*v = (int16)(binary.BigEndian.Uint16(encoding))

		case *int32:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 4)
			*v = (int32)(binary.BigEndian.Uint32(encoding))

		case *int64:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 8)
			*v = (int64)(binary.BigEndian.Uint64(encoding))

		case *float32:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 4)
			buf := bytes.NewReader(encoding)
			binary.Read(buf, binary.BigEndian, v)

		case *float64:
			encoding = d.nextEncoding()
			ensureLen(&encoding, 8)
			buf := bytes.NewReader(encoding)
			binary.Read(buf, binary.BigEndian, v)

		case *string:
			encoding = d.nextEncoding()
			*v = string(encoding)

		case *big.Int:
			encoding = d.nextEncoding()
			v.SetBytes(encoding)

		default:
			fmt.Println("Unknown type 2:")
			fmt.Println(reflect.TypeOf(v))
		}
	}

	return len(encoding), nil
}

func (d *Decoder) decodeInterface(v interface{}) {
	switch v := v.(type) {
	case *interface{}:
		if isList(*v) {
			kind := reflect.TypeOf(*v).Elem().Kind()

			if kind == reflect.Array || kind == reflect.Slice {
				kind = reflect.TypeOf(*v).Elem().Elem().Kind()
			}

			switch kind {
			case reflect.Uint8:
				decodeInterfaceList(d, v, []byte{})
			case reflect.Int8:
				decodeInterfaceList(d, v, []int8{})
			case reflect.Int16:
				decodeInterfaceList(d, v, []int16{})
			case reflect.Int32:
				decodeInterfaceList(d, v, []int32{})
			case reflect.Int64:
				decodeInterfaceList(d, v, []int64{})
			case reflect.Uint16:
				decodeInterfaceList(d, v, []uint16{})
			case reflect.Uint32:
				decodeInterfaceList(d, v, []uint32{})
			case reflect.Uint64:
				decodeInterfaceList(d, v, []uint64{})
			case reflect.Float32:
				decodeInterfaceList(d, v, []float32{})
			case reflect.Float64:
				decodeInterfaceList(d, v, []float64{})
			default:
				fmt.Println("Unknown type:")
				fmt.Println(v)
			}
		} else {
			a := reflect.TypeOf(*v)

			switch a.Kind() {
			case reflect.Bool:
				decodeInterfaceString(d, v, bool(false))
			case reflect.Uint8:
				decodeInterfaceString(d, v, uint8(0))
			case reflect.Uint16:
				decodeInterfaceString(d, v, uint16(0))
			case reflect.Uint32:
				decodeInterfaceString(d, v, uint32(0))
			case reflect.Uint64:
				decodeInterfaceString(d, v, uint64(0))
			case reflect.Int8:
				decodeInterfaceString(d, v, int8(0))
			case reflect.Int16:
				decodeInterfaceString(d, v, int16(0))
			case reflect.Int32:
				decodeInterfaceString(d, v, int32(0))
			case reflect.Int64:
				decodeInterfaceString(d, v, int64(0))
			case reflect.Float32:
				decodeInterfaceString(d, v, float32(0))
			case reflect.Float64:
				decodeInterfaceString(d, v, float64(0))
			case reflect.String:
				decodeInterfaceString(d, v, string(""))
			default:
				fmt.Println("Unknown type:")
				fmt.Println(a)
			}
		}
	}
}

func (d *Decoder) nextEncoding() []byte {
	firstByte := d.Encodings.Next(1)
	size := firstByte[0]

	if size <= 0x7f {
		return firstByte
	}

	if size <= 0xb7 {
		len := int(size - 0x80)
		return d.Encodings.Next(len)
	}

	if size <= 0xbf {
		len := int(size - 0xb7)
		buf := d.Encodings.Next(len)
		ensureLen(&buf, 8)

		size := (int)(binary.BigEndian.Uint64(buf))
		return d.Encodings.Next(size)
	}

	if size <= 0xf7 {
		len := int(size - 0xc0)
		return d.Encodings.Next(len)
	}

	if size <= 0xff {
		len := int(size - 0xf7)
		buf := d.Encodings.Next(len)
		ensureLen(&buf, 8)

		size := (int)(binary.BigEndian.Uint64(buf))
		return d.Encodings.Next(size)
	}

	fmt.Println("Unknown type")
	return nil
}

func ensureLen(buf *[]byte, length int) {
	bufLen := len(*buf)

	if bufLen < length {
		l := length - bufLen
		b := make([]byte, length)

		*buf = append(b[:l], *buf...)
	}
}
