package rlp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"reflect"
)

type Decoder struct {
	Encodings *bytes.Buffer
}

func NewDecoder(encodings []byte) *Decoder {
	return &Decoder{Encodings: bytes.NewBuffer(encodings)}
}

func (d *Decoder) Decode(v interface{}) error {
	val := reflect.ValueOf(v)

	decode(d, val)
	return nil
}

func kind(v reflect.Value) reflect.Kind {
	kind := v

	for kind.Kind() == reflect.Ptr || kind.Kind() == reflect.Interface {
		kind = kind.Elem()
	}

	return kind.Kind()
}

func decode(d *Decoder, val reflect.Value) error {
	// create and set element
	if val.Kind() == reflect.Ptr && val.IsZero() {
		val.Set(createValue(val.Type()))
		val = val.Elem()
	}

	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Check exotic cases not supported by reflect.Kind
	switch val.Interface().(type) {
	case big.Int, *big.Int:
		d.set(val)
		return nil
	}

	switch kind(val) {
	case reflect.Struct:
		d := NewDecoder(d.nextEncoding())

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			decode(d, field)
		}

	case reflect.Slice:
		d := NewDecoder(d.nextEncoding())
		tmp := createValue(reflect.TypeOf(val.Interface()).Elem()).Elem()

		slice := val
		if slice.Kind() == reflect.Interface {
			slice = val.Elem()
		}

		// Iterate all elements
		for d.Encodings.Len() > 0 {
			d.set(tmp)
			slice = reflect.Append(slice, tmp)
		}

		val.Set(slice)

	case reflect.Array:
		// TODO: For now this only handle []byte array
		dec := NewDecoder(d.nextEncoding())

		array := reflect.New(reflect.TypeOf(val.Interface())).Elem()
		reflect.Copy(array, reflect.ValueOf(dec.Encodings.Bytes()))

		val.Set(array)

	default:
		d.set(val)
	}

	return nil
}

func createValue(t reflect.Type) reflect.Value {
	typ := reflect.PtrTo(t).Elem()

	// iterate until element no longer be pointer
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return reflect.New(typ)
}

func (d *Decoder) set(val reflect.Value) {
	switch kind(val) {

	case reflect.String:
		encoding := d.nextEncoding()
		tmp := string(encoding)
		val.Set(reflect.ValueOf(tmp))

	case reflect.Bool:
		encoding := d.nextEncoding()
		tmp := len(encoding) > 0
		val.Set(reflect.ValueOf(tmp))

	case reflect.Uint8:
		encoding := d.nextEncoding()
		tmp := uint8(encoding[0])
		val.Set(reflect.ValueOf(tmp))

	case reflect.Uint16:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 2)
		tmp := binary.BigEndian.Uint16(encoding)
		val.Set(reflect.ValueOf(tmp))

	case reflect.Uint32:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 4)
		tmp := binary.BigEndian.Uint32(encoding)
		val.Set(reflect.ValueOf(tmp))

	case reflect.Uint64:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 8)
		tmp := binary.BigEndian.Uint64(encoding)
		val.Set(reflect.ValueOf(tmp))

	case reflect.Uint:
		encoding := d.nextEncoding()
		tmp := uint(encoding[0])
		val.Set(reflect.ValueOf(tmp))

	case reflect.Int8:
		encoding := d.nextEncoding()
		tmp := int8(encoding[0])
		val.Set(reflect.ValueOf(tmp))

	case reflect.Int16:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 2)
		tmp := (int16)(binary.BigEndian.Uint16(encoding))
		val.Set(reflect.ValueOf(tmp))

	case reflect.Int32:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 4)
		tmp := (int32)(binary.BigEndian.Uint32(encoding))
		val.Set(reflect.ValueOf(tmp))

	case reflect.Int64:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 8)
		tmp := (int64)(binary.BigEndian.Uint64(encoding))
		val.Set(reflect.ValueOf(tmp))

	case reflect.Float32:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 4)
		tmp := math.Float32frombits(binary.BigEndian.Uint32(encoding))
		val.Set(reflect.ValueOf(tmp))

	case reflect.Float64:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 8)
		tmp := math.Float64frombits(binary.BigEndian.Uint64(encoding))
		val.Set(reflect.ValueOf(tmp))

	case reflect.Array, reflect.Slice:
		decode(d, val)

	default:
		switch val.Interface().(type) {
		case big.Int:
			encoding := d.nextEncoding()
			b := new(big.Int).SetBytes(encoding)
			val.Set(reflect.ValueOf(b).Elem())
			return

		case *big.Int:
			encoding := d.nextEncoding()
			b := new(big.Int).SetBytes(encoding)
			val.Set(reflect.ValueOf(b))
			return
		}

		fmt.Printf("UNKNOWN TYPE: %s\n", kind(val))
	}
}

func (d *Decoder) nextEncoding() []byte {
	if len(d.Encodings.Bytes()) == 0 {
		return []byte{}
	}

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

	fmt.Println("UNSUPPORTED TYPE")
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
