// protocol has been generated from message format json - DO NOT EDIT
package protocol

type BrokerHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the broker ID.
	BrokerID int32
	// BrokerEpoch contains the broker epoch.
	BrokerEpoch int64
	// CurrentMetadataOffset contains the highest metadata offset which the broker has reached.
	CurrentMetadataOffset int64
	// WantFence contains a True if the broker wants to be fenced, false otherwise.
	WantFence bool
	// WantShutDown contains a True if the broker wants to be shut down, false otherwise.
	WantShutDown bool
}

func (r *BrokerHeartbeatRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.BrokerID)

	pe.putInt64(r.BrokerEpoch)

	pe.putInt64(r.CurrentMetadataOffset)

	pe.putBool(r.WantFence)

	pe.putBool(r.WantShutDown)

	pe.putUVarint(0)
	return nil
}

func (r *BrokerHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.BrokerEpoch, err = pd.getInt64(); err != nil {
		return err
	}

	if r.CurrentMetadataOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if r.WantFence, err = pd.getBool(); err != nil {
		return err
	}

	if r.WantShutDown, err = pd.getBool(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *BrokerHeartbeatRequest) GetKey() int16 {
	return 63
}

func (r *BrokerHeartbeatRequest) GetVersion() int16 {
	return r.Version
}

func (r *BrokerHeartbeatRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *BrokerHeartbeatRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *BrokerHeartbeatRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
