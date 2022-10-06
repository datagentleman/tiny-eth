package rlp

import (
	"bytes"
	"fmt"
	"reflect"
)

type Decoder struct {
	Encodings *bytes.Buffer
}

type ByteSetter interface {
	SetBytes([]byte)
}

func NewDecoder(encodings []byte) *Decoder {
	return &Decoder{Encodings: bytes.NewBuffer(encodings)}
}

func (d *Decoder) DecodeList(v interface{}) {
	switch v := v.(type) {

	// Special case - rlp treats []byte as string so we don't have to iterate it.
	case *[]uint8:
		*v = d.nextEncoding()

	case *[]interface{}:
		dec := NewDecoder(d.nextEncoding())

		for _, val := range *v {
			dec.Decode(val)
		}

	case *[]int8:
		decodeList(d, v)
	case *[]int16:
		decodeList(d, v)
	case *[]int32:
		decodeList(d, v)
	case *[]int64:
		decodeList(d, v)
	case *[]uint16:
		decodeList(d, v)
	case *[]uint32:
		decodeList(d, v)
	case *[]uint64:
		decodeList(d, v)
	case *[]float32:
		decodeList(d, v)
	case *[]float64:
		decodeList(d, v)
	default:
		bs, ok := v.(ByteSetter)
		if ok {
			bs.SetBytes(d.nextEncoding())
			return
		}

		fmt.Println("Unknown type:")
		fmt.Println(reflect.TypeOf(v))
	}
}

func decodeList[T any](d *Decoder, v *[]T) {
	dec := NewDecoder(d.nextEncoding())

	var elem T
	for dec.Encodings.Len() > 0 {
		dec.Decode(&elem)
		*v = append(*v, elem)
	}
}

func decodeInterfaceList[T any](d *Decoder, v *any, a []T) {
	d.Decode(&a)
	*v = a
}

func decodeInterfaceString[T any](d *Decoder, v *any, a T) {
	d.Decode(&a)
	*v = a
}
