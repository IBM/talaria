// protocol has been generated from message format json - DO NOT EDIT
package protocol

// TxnOffsetCommitRequestPartition contains the partitions inside the topic that we want to committ offsets for.
type TxnOffsetCommitRequestPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the index of the partition within the topic.
	PartitionIndex int32
	// CommittedOffset contains the message offset to be committed.
	CommittedOffset int64
	// CommittedLeaderEpoch contains the leader epoch of the last consumed record.
	CommittedLeaderEpoch int32
	// CommittedMetadata contains a Any associated metadata the client wants to keep.
	CommittedMetadata *string
}

func (p *TxnOffsetCommitRequestPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt64(p.CommittedOffset)

	if p.Version >= 2 {
		pe.putInt32(p.CommittedLeaderEpoch)
	}

	if err := pe.putNullableString(p.CommittedMetadata); err != nil {
		return err
	}

	if p.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *TxnOffsetCommitRequestPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.CommittedOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if p.CommittedLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.CommittedMetadata, err = pd.getNullableString(); err != nil {
		return err
	}

	if p.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// TxnOffsetCommitRequestTopic contains each topic that we want to commit offsets for.
type TxnOffsetCommitRequestTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the partitions inside the topic that we want to committ offsets for.
	Partitions []TxnOffsetCommitRequestPartition
}

func (t *TxnOffsetCommitRequestTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(t.Partitions)); err != nil {
		return err
	}
	for _, block := range t.Partitions {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *TxnOffsetCommitRequestTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]TxnOffsetCommitRequestPartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block TxnOffsetCommitRequestPartition
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type TxnOffsetCommitRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the ID of the transaction.
	TransactionalID string
	// GroupID contains the ID of the group.
	GroupID string
	// ProducerID contains the current producer ID in use by the transactional ID.
	ProducerID int64
	// ProducerEpoch contains the current epoch associated with the producer ID.
	ProducerEpoch int16
	// GenerationID contains the generation of the consumer.
	GenerationID int32
	// MemberID contains the member ID assigned by the group coordinator.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
	// Topics contains each topic that we want to commit offsets for.
	Topics []TxnOffsetCommitRequestTopic
}

func (r *TxnOffsetCommitRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.TransactionalID); err != nil {
		return err
	}

	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	pe.putInt64(r.ProducerID)

	pe.putInt16(r.ProducerEpoch)

	if r.Version >= 3 {
		pe.putInt32(r.GenerationID)
	}

	if r.Version >= 3 {
		if err := pe.putString(r.MemberID); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if err := pe.putNullableString(r.GroupInstanceID); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *TxnOffsetCommitRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.TransactionalID, err = pd.getString(); err != nil {
		return err
	}

	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ProducerEpoch, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.GenerationID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.MemberID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TxnOffsetCommitRequestTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TxnOffsetCommitRequestTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *TxnOffsetCommitRequest) GetKey() int16 {
	return 28
}

func (r *TxnOffsetCommitRequest) GetVersion() int16 {
	return r.Version
}

func (r *TxnOffsetCommitRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *TxnOffsetCommitRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *TxnOffsetCommitRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
