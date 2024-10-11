// protocol has been generated from message format json - DO NOT EDIT
package protocol

// OffsetForLeaderPartition contains each partition to get offsets for.
type OffsetForLeaderPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// CurrentLeaderEpoch contains a An epoch used to fence consumers/replicas with old metadata. If the epoch provided by the client is larger than the current epoch known to the broker, then the UNKNOWN_LEADER_EPOCH error code will be returned. If the provided epoch is smaller, then the FENCED_LEADER_EPOCH error code will be returned.
	CurrentLeaderEpoch int32
	// LeaderEpoch contains the epoch to look up an offset for.
	LeaderEpoch int32
}

func (p *OffsetForLeaderPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	if p.Version >= 2 {
		pe.putInt32(p.CurrentLeaderEpoch)
	}

	pe.putInt32(p.LeaderEpoch)

	if p.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *OffsetForLeaderPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if p.CurrentLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetForLeaderTopic contains each topic to get offsets for.
type OffsetForLeaderTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the topic name.
	Topic string
	// Partitions contains each partition to get offsets for.
	Partitions []OffsetForLeaderPartition
}

func (t *OffsetForLeaderTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Topic); err != nil {
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

	if t.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *OffsetForLeaderTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Topic, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]OffsetForLeaderPartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block OffsetForLeaderPartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	if t.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type OffsetForLeaderEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ReplicaID contains the broker ID of the follower, of -1 if this request is from a consumer.
	ReplicaID int32
	// Topics contains each topic to get offsets for.
	Topics []OffsetForLeaderTopic
}

func (r *OffsetForLeaderEpochRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		pe.putInt32(r.ReplicaID)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *OffsetForLeaderEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.ReplicaID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]OffsetForLeaderTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block OffsetForLeaderTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *OffsetForLeaderEpochRequest) GetKey() int16 {
	return 23
}

func (r *OffsetForLeaderEpochRequest) GetVersion() int16 {
	return r.Version
}

func (r *OffsetForLeaderEpochRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *OffsetForLeaderEpochRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *OffsetForLeaderEpochRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
