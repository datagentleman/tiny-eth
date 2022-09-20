package rlp

import (
	"bytes"
	"fmt"
	"reflect"
)

type Decoder struct {
	Buf *bytes.Buffer
}

func NewDecoder(encodings []byte) *Decoder {
	d := &Decoder{
		Buf: bytes.NewBuffer(encodings),
	}

	return d
}

func (d *Decoder) Next() *Decoder {
	return NewDecoder(d.Buf.Bytes())
}

func (d *Decoder) Decode(v interface{}) {
	_, list := nextEncoding(d.Buf)
	buf := bytes.NewBuffer(list)

	encoding := func() []byte {
		prefix, encoding := nextEncoding(buf)
		return append(prefix, encoding...)
	}

	switch v := v.(type) {
	case *[]int8:
		var elem int8

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]int16:
		var elem int16

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]int32:
		var elem int32

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]int64:
		var elem int64

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]uint16:
		var elem uint16

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]uint32:
		var elem uint32

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]uint64:
		var elem uint64

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]float32:
		var elem float32

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	case *[]float64:
		var elem float64

		for buf.Len() > 0 {
			Decode(encoding(), &elem)
			*v = append(*v, elem)
		}

	default:
		fmt.Println("Unknown type:")
		fmt.Println(reflect.TypeOf(v))
	}
}
