// protocol has been generated from message format json - DO NOT EDIT
package protocol

type AddOffsetsToTxnRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the transactional id corresponding to the transaction.
	TransactionalID string
	// ProducerID contains a Current producer id in use by the transactional id.
	ProducerID int64
	// ProducerEpoch contains a Current epoch associated with the producer id.
	ProducerEpoch int16
	// GroupID contains the unique group identifier.
	GroupID string
}

func (r *AddOffsetsToTxnRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.TransactionalID); err != nil {
		return err
	}

	pe.putInt64(r.ProducerID)

	pe.putInt16(r.ProducerEpoch)

	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AddOffsetsToTxnRequest) decode(pd packetDecoder, version int16) (err error) {
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

	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AddOffsetsToTxnRequest) GetKey() int16 {
	return 25
}

func (r *AddOffsetsToTxnRequest) GetVersion() int16 {
	return r.Version
}

func (r *AddOffsetsToTxnRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *AddOffsetsToTxnRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *AddOffsetsToTxnRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
