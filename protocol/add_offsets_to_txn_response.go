// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type AddOffsetsToTxnResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the response error code, or 0 if there was no error.
	ErrorCode int16
}

func (r *AddOffsetsToTxnResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AddOffsetsToTxnResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AddOffsetsToTxnResponse) GetKey() int16 {
	return 25
}

func (r *AddOffsetsToTxnResponse) GetVersion() int16 {
	return r.Version
}

func (r *AddOffsetsToTxnResponse) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 1
	}
	return 0
}

func (r *AddOffsetsToTxnResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *AddOffsetsToTxnResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AddOffsetsToTxnResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
