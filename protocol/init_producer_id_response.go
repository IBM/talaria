// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type InitProducerIdResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ProducerID contains the current producer id.
	ProducerID int64
	// ProducerEpoch contains the current epoch associated with the producer id.
	ProducerEpoch int16
}

func (r *InitProducerIdResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	pe.putInt64(r.ProducerID)

	pe.putInt16(r.ProducerEpoch)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *InitProducerIdResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ProducerEpoch, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *InitProducerIdResponse) GetKey() int16 {
	return 22
}

func (r *InitProducerIdResponse) GetVersion() int16 {
	return r.Version
}

func (r *InitProducerIdResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *InitProducerIdResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *InitProducerIdResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *InitProducerIdResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
