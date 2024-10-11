// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type HeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (r *HeartbeatResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	pe.putInt16(r.ErrorCode)

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *HeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *HeartbeatResponse) GetKey() int16 {
	return 12
}

func (r *HeartbeatResponse) GetVersion() int16 {
	return r.Version
}

func (r *HeartbeatResponse) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 1
	}
	return 0
}

func (r *HeartbeatResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *HeartbeatResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *HeartbeatResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
