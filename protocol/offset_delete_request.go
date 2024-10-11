// protocol has been generated from message format json - DO NOT EDIT
package protocol

// OffsetDeleteRequestPartition contains each partition to delete offsets for.
type OffsetDeleteRequestPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
}

func (p *OffsetDeleteRequestPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	return nil
}

func (p *OffsetDeleteRequestPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	return nil
}

// OffsetDeleteRequestTopic contains the topics to delete offsets for
type OffsetDeleteRequestTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains each partition to delete offsets for.
	Partitions []OffsetDeleteRequestPartition
}

func (t *OffsetDeleteRequestTopic) encode(pe packetEncoder, version int16) (err error) {
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

	return nil
}

func (t *OffsetDeleteRequestTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]OffsetDeleteRequestPartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block OffsetDeleteRequestPartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	return nil
}

type OffsetDeleteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the unique group identifier.
	GroupID string
	// Topics contains the topics to delete offsets for
	Topics []OffsetDeleteRequestTopic
}

func (r *OffsetDeleteRequest) encode(pe packetEncoder) (err error) {
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	return nil
}

func (r *OffsetDeleteRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]OffsetDeleteRequestTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block OffsetDeleteRequestTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	return nil
}

func (r *OffsetDeleteRequest) GetKey() int16 {
	return 47
}

func (r *OffsetDeleteRequest) GetVersion() int16 {
	return r.Version
}

func (r *OffsetDeleteRequest) GetHeaderVersion() int16 {
	return 1
}

func (r *OffsetDeleteRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *OffsetDeleteRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
