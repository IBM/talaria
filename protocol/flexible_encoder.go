package protocol

import uuid "github.com/google/uuid"

type flexibleEncoder struct {
	parent packetEncoder
}

func (fe *flexibleEncoder) putInt8(in int8) {
	fe.parent.putInt8(in)
}

func (fe *flexibleEncoder) putInt16(in int16) {
	fe.parent.putInt16(in)
}

func (fe *flexibleEncoder) putUint16(in uint16) {
	fe.parent.putUint16(in)
}

func (fe *flexibleEncoder) putInt32(in int32) {
	fe.parent.putInt32(in)
}

func (fe *flexibleEncoder) putUint32(in uint32) {
	fe.parent.putUint32(in)
}

func (fe *flexibleEncoder) putInt64(in int64) {
	fe.parent.putInt64(in)
}

func (fe *flexibleEncoder) putFloat64(in float64) {
	fe.parent.putFloat64(in)
}

func (fe *flexibleEncoder) putVarint(in int64) {
	fe.parent.putVarint(in)
}

func (fe *flexibleEncoder) putUVarint(in uint64) {
	fe.parent.putUVarint(in)
}

func (fe *flexibleEncoder) putCompactArrayLength(in int) error {
	return fe.parent.putCompactArrayLength(in)
}

func (fe *flexibleEncoder) putArrayLength(in int) error {
	return fe.parent.putCompactArrayLength(in)
}

func (fe *flexibleEncoder) putBool(in bool) {
	fe.parent.putBool(in)
}

// Collections
func (fe *flexibleEncoder) putBytes(in []byte) error {
	return fe.parent.putCompactBytes(in)
}

func (fe *flexibleEncoder) putVarintBytes(in []byte) error {
	return fe.parent.putVarintBytes(in)
}

func (fe *flexibleEncoder) putCompactBytes(in []byte) error {
	return fe.parent.putCompactBytes(in)
}

func (fe *flexibleEncoder) putRawBytes(in []byte) error {
	return fe.parent.putRawBytes(in)
}

func (fe *flexibleEncoder) putCompactString(in string) error {
	return fe.parent.putCompactString(in)
}

func (fe *flexibleEncoder) putNullableCompactString(in *string) error {
	return fe.parent.putNullableCompactString(in)
}

func (fe *flexibleEncoder) putString(in string) error {
	return fe.parent.putCompactString(in)
}

func (fe *flexibleEncoder) putNullableString(in *string) error {
	return fe.parent.putNullableString(in)
}

func (fe *flexibleEncoder) putStringArray(in []string) error {
	return fe.parent.putStringArray(in)
}

func (fe *flexibleEncoder) putCompactInt32Array(in []int32) error {
	return fe.parent.putCompactInt32Array(in)
}

func (fe *flexibleEncoder) putNullableCompactInt32Array(in []int32) error {
	return fe.parent.putNullableCompactInt32Array(in)
}

func (fe *flexibleEncoder) putInt32Array(in []int32) error {
	return fe.parent.putCompactInt32Array(in)
}

func (fe *flexibleEncoder) putInt64Array(in []int64) error {
	return fe.parent.putInt64Array(in)
}

func (fe *flexibleEncoder) putEmptyTaggedFieldArray() {
	fe.parent.putEmptyTaggedFieldArray()
}

func (fe *flexibleEncoder) putUUID(in uuid.UUID) error {
	return fe.parent.putUUID(in)
}

// Provide the current offset to record the batch size metric
func (fe *flexibleEncoder) offset() int {
	return fe.parent.offset()
}

// Stacks, see PushEncoder
func (fe *flexibleEncoder) push(in pushEncoder) {
	fe.parent.push(in)
}

func (fe *flexibleEncoder) pop() error {
	return fe.parent.pop()
}
