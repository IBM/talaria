package protocol

import (
	"encoding/binary"
	"errors"
	"math"

	uuid "github.com/google/uuid"
)

type realEncoder struct {
	raw   []byte
	off   int
	stack []pushEncoder
}

// primitives

func (re *realEncoder) putInt8(in int8) {
	re.raw[re.off] = byte(in)
	re.off++
}

func (re *realEncoder) putInt16(in int16) {
	re.putUint16(uint16(in))
}

func (re *realEncoder) putUint16(in uint16) {
	binary.BigEndian.PutUint16(re.raw[re.off:], in)
	re.off += 2
}

func (re *realEncoder) putInt32(in int32) {
	re.putUint32(uint32(in))
}

func (re *realEncoder) putUint32(in uint32) {
	binary.BigEndian.PutUint32(re.raw[re.off:], in)
	re.off += 4
}

func (re *realEncoder) putInt64(in int64) {
	binary.BigEndian.PutUint64(re.raw[re.off:], uint64(in))
	re.off += 8
}

func (re *realEncoder) putVarint(in int64) {
	re.off += binary.PutVarint(re.raw[re.off:], in)
}

func (re *realEncoder) putUVarint(in uint64) {
	re.off += binary.PutUvarint(re.raw[re.off:], in)
}

func (re *realEncoder) putFloat64(in float64) {
	binary.BigEndian.PutUint64(re.raw[re.off:], math.Float64bits(in))
	re.off += 8
}

func (re *realEncoder) putArrayLength(in int) error {
	re.putInt32(int32(in))
	return nil
}

func (re *realEncoder) putCompactArrayLength(in int) error {
	// 0 represents a null array, so +1 has to be added
	re.putUVarint(uint64(in + 1))
	return nil
}

func (re *realEncoder) putBool(in bool) {
	if in {
		re.putInt8(1)
		return
	}
	re.putInt8(0)
}

// collection

func (re *realEncoder) putRawBytes(in []byte) error {
	copy(re.raw[re.off:], in)
	re.off += len(in)
	return nil
}

func (re *realEncoder) putBytes(in []byte) error {
	if in == nil {
		re.putInt32(-1)
		return nil
	}
	re.putInt32(int32(len(in)))
	return re.putRawBytes(in)
}

func (re *realEncoder) putVarintBytes(in []byte) error {
	if in == nil {
		re.putVarint(-1)
		return nil
	}
	re.putVarint(int64(len(in)))
	return re.putRawBytes(in)
}

func (re *realEncoder) putCompactBytes(in []byte) error {
	re.putUVarint(uint64(len(in) + 1))
	return re.putRawBytes(in)
}

func (re *realEncoder) putCompactString(in string) error {
	re.putCompactArrayLength(len(in))
	return re.putRawBytes([]byte(in))
}

func (re *realEncoder) putNullableCompactString(in *string) error {
	if in == nil {
		re.putInt8(0)
		return nil
	}
	return re.putCompactString(*in)
}

func (re *realEncoder) putString(in string) error {
	re.putInt16(int16(len(in)))
	copy(re.raw[re.off:], in)
	re.off += len(in)
	return nil
}

func (re *realEncoder) putNullableString(in *string) error {
	if in == nil {
		re.putInt16(-1)
		return nil
	}
	return re.putString(*in)
}

func (re *realEncoder) putStringArray(in []string) error {
	err := re.putArrayLength(len(in))
	if err != nil {
		return err
	}

	for _, val := range in {
		if err := re.putString(val); err != nil {
			return err
		}
	}

	return nil
}

func (re *realEncoder) putCompactInt32Array(in []int32) error {
	if in == nil {
		return errors.New("expected int32 array to be non null")
	}
	// 0 represents a null array, so +1 has to be added
	re.putUVarint(uint64(len(in)) + 1)
	for _, val := range in {
		re.putInt32(val)
	}
	return nil
}

func (re *realEncoder) putNullableCompactInt32Array(in []int32) error {
	if in == nil {
		re.putUVarint(0)
		return nil
	}
	// 0 represents a null array, so +1 has to be added
	re.putUVarint(uint64(len(in)) + 1)
	for _, val := range in {
		re.putInt32(val)
	}
	return nil
}

func (re *realEncoder) putInt32Array(in []int32) error {
	err := re.putArrayLength(len(in))
	if err != nil {
		return err
	}
	for _, val := range in {
		re.putInt32(val)
	}
	return nil
}

func (re *realEncoder) putInt64Array(in []int64) error {
	err := re.putArrayLength(len(in))
	if err != nil {
		return err
	}
	for _, val := range in {
		re.putInt64(val)
	}
	return nil
}

func (re *realEncoder) putEmptyTaggedFieldArray() {
	re.putUVarint(0)
}

func (re *realEncoder) putUUID(in uuid.UUID) error {
	bytes, err := in.MarshalBinary()
	if err != nil {
		return err
	}

	return re.putRawBytes(bytes)
}

func (re *realEncoder) offset() int {
	return re.off
}

func (re *realEncoder) Bytes() []byte {
	return re.raw
}

// stacks

func (re *realEncoder) push(in pushEncoder) {
	in.saveOffset(re.off)
	re.off += in.reserveLength()
	re.stack = append(re.stack, in)
}

func (re *realEncoder) pop() error {
	// this is go's ugly pop pattern (the inverse of append)
	in := re.stack[len(re.stack)-1]
	re.stack = re.stack[:len(re.stack)-1]

	return in.run(re.off, re.raw)
}
