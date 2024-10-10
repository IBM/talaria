// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type AllocateProducerIdsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code
	ErrorCode int16
	// ProducerIdStart contains the first producer ID in this range, inclusive
	ProducerIdStart int64
	// ProducerIdLen contains the number of producer IDs in this range
	ProducerIdLen int32
}

func (r *AllocateProducerIdsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	pe.putInt64(r.ProducerIdStart)

	pe.putInt32(r.ProducerIdLen)

	pe.putUVarint(0)
	return nil
}

func (r *AllocateProducerIdsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ProducerIdStart, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ProducerIdLen, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *AllocateProducerIdsResponse) GetKey() int16 {
	return 67
}

func (r *AllocateProducerIdsResponse) GetVersion() int16 {
	return r.Version
}

func (r *AllocateProducerIdsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *AllocateProducerIdsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AllocateProducerIdsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AllocateProducerIdsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
