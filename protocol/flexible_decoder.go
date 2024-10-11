package protocol

import uuid "github.com/google/uuid"

// flexibleDecoder implements a shim for packetDecoder that calls the compact methods for getString and getArray
type flexibleDecoder struct {
	parent packetDecoder
}

// Primitives
func (fd *flexibleDecoder) getInt8() (int8, error) {
	return fd.parent.getInt8()
}

func (fd *flexibleDecoder) getInt16() (int16, error) {
	return fd.parent.getInt16()
}

func (fd *flexibleDecoder) getUint16() (uint16, error) {
	return fd.parent.getUint16()
}

func (fd *flexibleDecoder) getInt32() (int32, error) {
	return fd.parent.getInt32()
}

func (fd *flexibleDecoder) getUint32() (uint32, error) {
	return fd.parent.getUint32()
}

func (fd *flexibleDecoder) getInt64() (int64, error) {
	return fd.parent.getInt64()
}

func (fd *flexibleDecoder) getFloat64() (float64, error) {
	return fd.parent.getFloat64()
}

func (fd *flexibleDecoder) getVarint() (int64, error) {
	return fd.parent.getVarint()
}

func (fd *flexibleDecoder) getUVarint() (uint64, error) {
	return fd.parent.getUVarint()
}

func (fd *flexibleDecoder) getArrayLength() (int, error) {
	return fd.parent.getCompactArrayLength()
}

func (fd *flexibleDecoder) getCompactArrayLength() (int, error) {
	return fd.parent.getCompactArrayLength()
}

func (fd *flexibleDecoder) getBool() (bool, error) {
	return fd.parent.getBool()
}

func (fd *flexibleDecoder) getEmptyTaggedFieldArray() (int, error) {
	return fd.parent.getEmptyTaggedFieldArray()
}

// Collections
func (fd *flexibleDecoder) getBytes() ([]byte, error) {
	return fd.parent.getCompactBytes()
}

func (fd *flexibleDecoder) getCompactBytes() ([]byte, error) {
	return fd.parent.getCompactBytes()
}

func (fd *flexibleDecoder) getVarintBytes() ([]byte, error) {
	return fd.parent.getVarintBytes()
}

func (fd *flexibleDecoder) getRawBytes(length int) ([]byte, error) {
	return fd.parent.getRawBytes(length)
}

func (fd *flexibleDecoder) getString() (string, error) {
	return fd.parent.getCompactString()
}

func (fd *flexibleDecoder) getNullableString() (*string, error) {
	return fd.parent.getNullableString()
}

func (fd *flexibleDecoder) getCompactString() (string, error) {
	return fd.parent.getCompactString()
}

func (fd *flexibleDecoder) getCompactNullableString() (*string, error) {
	return fd.parent.getCompactNullableString()
}

func (fd *flexibleDecoder) getCompactInt32Array() ([]int32, error) {
	return fd.parent.getCompactInt32Array()
}

func (fd *flexibleDecoder) getInt32Array() ([]int32, error) {
	return fd.parent.getCompactInt32Array()
}

func (fd *flexibleDecoder) getInt64Array() ([]int64, error) {
	return fd.parent.getInt64Array()
}

func (fd *flexibleDecoder) getStringArray() ([]string, error) {
	return fd.parent.getStringArray()
}

func (fd *flexibleDecoder) getUUID() (uuid.UUID, error) {
	return fd.parent.getUUID()
}

// Subsets
func (fd *flexibleDecoder) remaining() int {
	return fd.parent.remaining()
}

func (fd *flexibleDecoder) getSubset(length int) (packetDecoder, error) {
	return fd.parent.getSubset(length)
}

func (fd *flexibleDecoder) peek(offset, length int) (packetDecoder, error) {
	return fd.parent.peek(offset, length)
}

func (fd *flexibleDecoder) peekInt8(offset int) (int8, error) {
	return fd.parent.peekInt8(offset)
}

// Stacks, see PushDecoder
func (fd *flexibleDecoder) push(in pushDecoder) error {
	return fd.parent.push(in)
}

func (fd *flexibleDecoder) pop() error {
	return fd.parent.pop()
}
