// protocol has been generated from message format json - DO NOT EDIT
package protocol

type InitProducerIdRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the transactional id, or null if the producer is not transactional.
	TransactionalID *string
	// TransactionTimeoutMs contains the time in ms to wait before aborting idle transactions sent by this producer. This is only relevant if a TransactionalId has been defined.
	TransactionTimeoutMs int32
	// ProducerID contains the producer id. This is used to disambiguate requests if a transactional id is reused following its expiration.
	ProducerID int64
	// ProducerEpoch contains the producer's current epoch. This will be checked against the producer epoch on the broker, and the request will return an error if they do not match.
	ProducerEpoch int16
}

func (r *InitProducerIdRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putNullableString(r.TransactionalID); err != nil {
		return err
	}

	pe.putInt32(r.TransactionTimeoutMs)

	if r.Version >= 3 {
		pe.putInt64(r.ProducerID)
	}

	if r.Version >= 3 {
		pe.putInt16(r.ProducerEpoch)
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *InitProducerIdRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.TransactionalID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.TransactionTimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.ProducerID, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.ProducerEpoch, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *InitProducerIdRequest) GetKey() int16 {
	return 22
}

func (r *InitProducerIdRequest) GetVersion() int16 {
	return r.Version
}

func (r *InitProducerIdRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *InitProducerIdRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *InitProducerIdRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
