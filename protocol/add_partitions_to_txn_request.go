// protocol has been generated from message format json - DO NOT EDIT
package protocol

// AddPartitionsToTxnTopic_AddPartitionsToTxnRequest contains a
type AddPartitionsToTxnTopic_AddPartitionsToTxnRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the topic.
	Name string
	// Partitions contains the partition indexes to add to the transaction
	Partitions []int32
}

func (a *AddPartitionsToTxnTopic_AddPartitionsToTxnRequest) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putString(a.Name); err != nil {
		return err
	}

	if err := pe.putInt32Array(a.Partitions); err != nil {
		return err
	}

	if a.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *AddPartitionsToTxnTopic_AddPartitionsToTxnRequest) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.Name, err = pd.getString(); err != nil {
		return err
	}

	if a.Partitions, err = pd.getInt32Array(); err != nil {
		return err
	}

	if a.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// AddPartitionsToTxnTransaction contains a List of transactions to add partitions to.
type AddPartitionsToTxnTransaction struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the transactional id corresponding to the transaction.
	TransactionalID string
	// ProducerID contains a Current producer id in use by the transactional id.
	ProducerID int64
	// ProducerEpoch contains a Current epoch associated with the producer id.
	ProducerEpoch int16
	// VerifyOnly contains a Boolean to signify if we want to check if the partition is in the transaction rather than add it.
	VerifyOnly bool
	// Topics contains the partitions to add to the transaction.
	Topics []AddPartitionsToTxnTopic_AddPartitionsToTxnRequest
}

func (t *AddPartitionsToTxnTransaction) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 4 {
		if err := pe.putString(t.TransactionalID); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		pe.putInt64(t.ProducerID)
	}

	if t.Version >= 4 {
		pe.putInt16(t.ProducerEpoch)
	}

	if t.Version >= 4 {
		pe.putBool(t.VerifyOnly)
	}

	if t.Version >= 4 {
		if err := pe.putArrayLength(len(t.Topics)); err != nil {
			return err
		}
		for _, block := range t.Topics {
			if err := block.encode(pe, t.Version); err != nil {
				return err
			}
		}
	}

	if t.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *AddPartitionsToTxnTransaction) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 4 {
		if t.TransactionalID, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		if t.ProducerID, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		if t.ProducerEpoch, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		if t.VerifyOnly, err = pd.getBool(); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopics > 0 {
			t.Topics = make([]AddPartitionsToTxnTopic_AddPartitionsToTxnRequest, numTopics)
			for i := 0; i < numTopics; i++ {
				var block AddPartitionsToTxnTopic_AddPartitionsToTxnRequest
				if err := block.decode(pd, t.Version); err != nil {
					return err
				}
				t.Topics[i] = block
			}
		}
	}

	if t.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AddPartitionsToTxnRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Transactions contains a List of transactions to add partitions to.
	Transactions []AddPartitionsToTxnTransaction
	// V3AndBelowTransactionalID contains the transactional id corresponding to the transaction.
	V3AndBelowTransactionalID string
	// V3AndBelowProducerID contains a Current producer id in use by the transactional id.
	V3AndBelowProducerID int64
	// V3AndBelowProducerEpoch contains a Current epoch associated with the producer id.
	V3AndBelowProducerEpoch int16
	// V3AndBelowTopics contains the partitions to add to the transaction.
	V3AndBelowTopics []AddPartitionsToTxnTopic_AddPartitionsToTxnRequest
}

func (r *AddPartitionsToTxnRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 4 {
		if err := pe.putArrayLength(len(r.Transactions)); err != nil {
			return err
		}
		for _, block := range r.Transactions {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if err := pe.putString(r.V3AndBelowTransactionalID); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		pe.putInt64(r.V3AndBelowProducerID)
	}

	if r.Version >= 0 && r.Version <= 3 {
		pe.putInt16(r.V3AndBelowProducerEpoch)
	}

	if r.Version >= 0 && r.Version <= 3 {
		if err := pe.putArrayLength(len(r.V3AndBelowTopics)); err != nil {
			return err
		}
		for _, block := range r.V3AndBelowTopics {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AddPartitionsToTxnRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 4 {
		var numTransactions int
		if numTransactions, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTransactions > 0 {
			r.Transactions = make([]AddPartitionsToTxnTransaction, numTransactions)
			for i := 0; i < numTransactions; i++ {
				var block AddPartitionsToTxnTransaction
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.Transactions[i] = block
			}
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.V3AndBelowTransactionalID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.V3AndBelowProducerID, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.V3AndBelowProducerEpoch, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		var numV3AndBelowTopics int
		if numV3AndBelowTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numV3AndBelowTopics > 0 {
			r.V3AndBelowTopics = make([]AddPartitionsToTxnTopic_AddPartitionsToTxnRequest, numV3AndBelowTopics)
			for i := 0; i < numV3AndBelowTopics; i++ {
				var block AddPartitionsToTxnTopic_AddPartitionsToTxnRequest
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.V3AndBelowTopics[i] = block
			}
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AddPartitionsToTxnRequest) GetKey() int16 {
	return 24
}

func (r *AddPartitionsToTxnRequest) GetVersion() int16 {
	return r.Version
}

func (r *AddPartitionsToTxnRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *AddPartitionsToTxnRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *AddPartitionsToTxnRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
