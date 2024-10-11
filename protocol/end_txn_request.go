// protocol has been generated from message format json - DO NOT EDIT
package protocol

type EndTxnRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the ID of the transaction to end.
	TransactionalID string
	// ProducerID contains the producer ID.
	ProducerID int64
	// ProducerEpoch contains the current epoch associated with the producer.
	ProducerEpoch int16
	// Committed contains a True if the transaction was committed, false if it was aborted.
	Committed bool
}

func (r *EndTxnRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.TransactionalID); err != nil {
		return err
	}

	pe.putInt64(r.ProducerID)

	pe.putInt16(r.ProducerEpoch)

	pe.putBool(r.Committed)

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *EndTxnRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.TransactionalID, err = pd.getString(); err != nil {
		return err
	}

	if r.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ProducerEpoch, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Committed, err = pd.getBool(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *EndTxnRequest) GetKey() int16 {
	return 26
}

func (r *EndTxnRequest) GetVersion() int16 {
	return r.Version
}

func (r *EndTxnRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *EndTxnRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *EndTxnRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
