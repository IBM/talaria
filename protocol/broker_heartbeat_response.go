// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type BrokerHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// IsCaughtUp contains a True if the broker has approximately caught up with the latest metadata.
	IsCaughtUp bool
	// IsFenced contains a True if the broker is fenced.
	IsFenced bool
	// ShouldShutDown contains a True if the broker should proceed with its shutdown.
	ShouldShutDown bool
}

func (r *BrokerHeartbeatResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	pe.putBool(r.IsCaughtUp)

	pe.putBool(r.IsFenced)

	pe.putBool(r.ShouldShutDown)

	pe.putUVarint(0)
	return nil
}

func (r *BrokerHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.IsCaughtUp, err = pd.getBool(); err != nil {
		return err
	}

	if r.IsFenced, err = pd.getBool(); err != nil {
		return err
	}

	if r.ShouldShutDown, err = pd.getBool(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *BrokerHeartbeatResponse) GetKey() int16 {
	return 63
}

func (r *BrokerHeartbeatResponse) GetVersion() int16 {
	return r.Version
}

func (r *BrokerHeartbeatResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *BrokerHeartbeatResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *BrokerHeartbeatResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *BrokerHeartbeatResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
