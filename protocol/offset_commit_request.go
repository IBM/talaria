// protocol has been generated from message format json - DO NOT EDIT
package protocol

// OffsetCommitRequestPartition contains each partition to commit offsets for.
type OffsetCommitRequestPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// CommittedOffset contains the message offset to be committed.
	CommittedOffset int64
	// CommittedLeaderEpoch contains the leader epoch of this partition.
	CommittedLeaderEpoch int32
	// CommitTimestamp contains the timestamp of the commit.
	CommitTimestamp int64
	// CommittedMetadata contains a Any associated metadata the client wants to keep.
	CommittedMetadata *string
}

func (p *OffsetCommitRequestPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt64(p.CommittedOffset)

	if p.Version >= 6 {
		pe.putInt32(p.CommittedLeaderEpoch)
	}

	if p.Version == 1 {
		pe.putInt64(p.CommitTimestamp)
	}

	if err := pe.putNullableString(p.CommittedMetadata); err != nil {
		return err
	}

	if p.Version >= 8 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *OffsetCommitRequestPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.CommittedOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 6 {
		if p.CommittedLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version == 1 {
		if p.CommitTimestamp, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.CommittedMetadata, err = pd.getNullableString(); err != nil {
		return err
	}

	if p.Version >= 8 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetCommitRequestTopic contains the topics to commit offsets for.
type OffsetCommitRequestTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains each partition to commit offsets for.
	Partitions []OffsetCommitRequestPartition
}

func (t *OffsetCommitRequestTopic) encode(pe packetEncoder, version int16) (err error) {
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

	if t.Version >= 8 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *OffsetCommitRequestTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]OffsetCommitRequestPartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block OffsetCommitRequestPartition
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 8 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type OffsetCommitRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the unique group identifier.
	GroupID string
	// GenerationID contains the generation of the group.
	GenerationID int32
	// MemberID contains the member ID assigned by the group coordinator.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
	// RetentionTimeMs contains the time period in ms to retain the offset.
	RetentionTimeMs int64
	// Topics contains the topics to commit offsets for.
	Topics []OffsetCommitRequestTopic
}

func (r *OffsetCommitRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 8 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putInt32(r.GenerationID)
	}

	if r.Version >= 1 {
		if err := pe.putString(r.MemberID); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if err := pe.putNullableString(r.GroupInstanceID); err != nil {
			return err
		}
	}

	if r.Version >= 2 && r.Version <= 4 {
		pe.putInt64(r.RetentionTimeMs)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 8 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *OffsetCommitRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 8 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.GenerationID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		if r.MemberID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if r.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 2 && r.Version <= 4 {
		if r.RetentionTimeMs, err = pd.getInt64(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]OffsetCommitRequestTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block OffsetCommitRequestTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 8 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *OffsetCommitRequest) GetKey() int16 {
	return 8
}

func (r *OffsetCommitRequest) GetVersion() int16 {
	return r.Version
}

func (r *OffsetCommitRequest) GetHeaderVersion() int16 {
	if r.Version >= 8 {
		return 2
	}
	return 1
}

func (r *OffsetCommitRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *OffsetCommitRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
