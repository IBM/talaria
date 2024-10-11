// protocol has been generated from message format json - DO NOT EDIT
package protocol

type UnregisterBrokerRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the broker ID to unregister.
	BrokerID int32
}

func (r *UnregisterBrokerRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.BrokerID)

	pe.putUVarint(0)
	return nil
}

func (r *UnregisterBrokerRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *UnregisterBrokerRequest) GetKey() int16 {
	return 64
}

func (r *UnregisterBrokerRequest) GetVersion() int16 {
	return r.Version
}

func (r *UnregisterBrokerRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *UnregisterBrokerRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *UnregisterBrokerRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
