package rlp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type Decoder struct {
	Encodings *bytes.Buffer
}

func NewDecoder(encodings []byte) *Decoder {
	return &Decoder{Encodings: bytes.NewBuffer(encodings)}
}

func (d *Decoder) Decode(v interface{}) error {
	val := getValue(v)

	val1 := val
	if val.Kind() == reflect.Interface {
		val1 = val.Elem()
	}

	// Struct || Lists
	if isStruct(v) || isList2(val1) {
		encode(d, val)
		return nil
	}

	// Strings
	d.set(val)
	return nil
}

func getValue(elem interface{}) reflect.Value {
	val := reflect.ValueOf(elem)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func encode(d *Decoder, val reflect.Value) error {
	// create and set element
	if val.Kind() == reflect.Ptr && val.IsZero() {
		val.Set(createValue(val.Type()))
		val = val.Elem()
	}

	tmp := val

	if tmp.Kind() == reflect.Interface {
		tmp = val.Elem()
	}

	switch tmp.Kind() {

	case reflect.Struct:
		d := NewDecoder(d.nextEncoding())

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			encode(d, field)
		}

	case reflect.Slice, reflect.Array:
		d := NewDecoder(d.nextEncoding())
		tmp := createValue(reflect.TypeOf(val.Interface()).Elem()).Elem()

		list := val
		if list.Kind() == reflect.Interface {
			list = val.Elem()
		}

		// Iterate all elements from encoding
		for d.Encodings.Len() > 0 {
			d.set(tmp)
			list = reflect.Append(list, tmp)
		}

		val.Set(list)

	default:
		// set basic types
		d.set(val)
	}

	return nil
}

func createValue(t reflect.Type) reflect.Value {
	typ := reflect.PtrTo(t).Elem()

	// iterate until element won't be a pointer
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return reflect.New(typ)
}

func (d *Decoder) set(val reflect.Value) {
	isPointer := val.Kind() == reflect.Ptr

	tmp := val.Type()

	if isPointer {
		tmp = tmp.Elem()
	}

	if tmp.Kind() == reflect.Interface {
		tmp = val.Elem().Type()
	}

	switch tmp.Kind() {

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
		buf := bytes.NewReader(encoding)
		tmp := float32(0)
		binary.Read(buf, binary.BigEndian, &tmp)
		val.Set(reflect.ValueOf(tmp))

	case reflect.Float64:
		encoding := d.nextEncoding()
		ensureLen(&encoding, 8)
		buf := bytes.NewReader(encoding)
		tmp := float64(0)
		binary.Read(buf, binary.BigEndian, &tmp)
		val.Set(reflect.ValueOf(tmp))

	default:
		fmt.Println("UNKNOWN TYPE")
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
