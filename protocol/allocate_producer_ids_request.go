// protocol has been generated from message format json - DO NOT EDIT
package protocol

type AllocateProducerIdsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the ID of the requesting broker
	BrokerID int32
	// BrokerEpoch contains the epoch of the requesting broker
	BrokerEpoch int64
}

func (r *AllocateProducerIdsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.BrokerID)

	pe.putInt64(r.BrokerEpoch)

	pe.putUVarint(0)
	return nil
}

func (r *AllocateProducerIdsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.BrokerID, err = pd.getInt32(); err != nil {
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

func (r *AllocateProducerIdsRequest) GetKey() int16 {
	return 67
}

func (r *AllocateProducerIdsRequest) GetVersion() int16 {
	return r.Version
}

func (r *AllocateProducerIdsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *AllocateProducerIdsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AllocateProducerIdsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
