// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type BrokerRegistrationResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// BrokerEpoch contains the broker's assigned epoch, or -1 if none was assigned.
	BrokerEpoch int64
}

func (r *BrokerRegistrationResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	pe.putInt64(r.BrokerEpoch)

	pe.putUVarint(0)
	return nil
}

func (r *BrokerRegistrationResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.BrokerEpoch, err = pd.getInt64(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *BrokerRegistrationResponse) GetKey() int16 {
	return 62
}

func (r *BrokerRegistrationResponse) GetVersion() int16 {
	return r.Version
}

func (r *BrokerRegistrationResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *BrokerRegistrationResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *BrokerRegistrationResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *BrokerRegistrationResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
