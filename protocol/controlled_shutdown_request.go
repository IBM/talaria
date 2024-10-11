// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ControlledShutdownRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the id of the broker for which controlled shutdown has been requested.
	BrokerID int32
	// BrokerEpoch contains the broker epoch.
	BrokerEpoch int64
}

func (r *ControlledShutdownRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.BrokerID)

	if r.Version >= 2 {
		pe.putInt64(r.BrokerEpoch)
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ControlledShutdownRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if r.BrokerEpoch, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ControlledShutdownRequest) GetKey() int16 {
	return 7
}

func (r *ControlledShutdownRequest) GetVersion() int16 {
	return r.Version
}

func (r *ControlledShutdownRequest) GetHeaderVersion() int16 {
	if r.Version == 0 {
		return 0
	}
	return 1
}

func (r *ControlledShutdownRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *ControlledShutdownRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
