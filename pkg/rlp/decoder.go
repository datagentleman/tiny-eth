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

func (d *Decoder) Any() bool {
	return d.Buf.Len() > 0
}

func (d *Decoder) Next() *Decoder {
	return NewDecoder(d.Buf.Bytes())
}

func (d *Decoder) Decode(v interface{}) {
	// TODO: For now I'm just playing with decoding slices. More will come.

	_, list := nextEncoding(d.Buf)
	buf := bytes.NewBuffer(list)

	switch v := v.(type) {
	case *[]int32:
		var elem int32
		var l []int32

		for buf.Len() > 0 {
			prefix, encoding := nextEncoding(buf)
			encoding = append(prefix, encoding...)

			Decode(encoding, &elem)
			l = append(l, elem)
		}

		*v = l

	default:
		fmt.Println("Unknown type:")
		fmt.Println(reflect.TypeOf(v))
	}

	fmt.Println()
}
