// protocol has been generated from message format json - DO NOT EDIT
package protocol

// SnapshotId_FetchSnapshotRequest contains the snapshot endOffset and epoch to fetch
type SnapshotId_FetchSnapshotRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EndOffset contains a
	EndOffset int64
	// Epoch contains a
	Epoch int32
}

func (s *SnapshotId_FetchSnapshotRequest) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	pe.putInt64(s.EndOffset)

	pe.putInt32(s.Epoch)

	pe.putUVarint(0)
	return nil
}

func (s *SnapshotId_FetchSnapshotRequest) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.EndOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if s.Epoch, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// PartitionSnapshot_FetchSnapshotRequest contains the partitions to fetch
type PartitionSnapshot_FetchSnapshotRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index
	Partition int32
	// CurrentLeaderEpoch contains the current leader epoch of the partition, -1 for unknown leader epoch
	CurrentLeaderEpoch int32
	// SnapshotID contains the snapshot endOffset and epoch to fetch
	SnapshotID SnapshotId_FetchSnapshotRequest
	// Position contains the byte position within the snapshot to start fetching from
	Position int64
}

func (p *PartitionSnapshot_FetchSnapshotRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	pe.putInt32(p.CurrentLeaderEpoch)

	if err := p.SnapshotID.encode(pe, p.Version); err != nil {
		return err
	}

	pe.putInt64(p.Position)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionSnapshot_FetchSnapshotRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.CurrentLeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	tmpSnapshotId_FetchSnapshotRequest := SnapshotId_FetchSnapshotRequest{}
	if err := tmpSnapshotId_FetchSnapshotRequest.decode(pd, p.Version); err != nil {
		return err
	}
	p.SnapshotID = tmpSnapshotId_FetchSnapshotRequest

	if p.Position, err = pd.getInt64(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicSnapshot_FetchSnapshotRequest contains the topics to fetch
type TopicSnapshot_FetchSnapshotRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the topic to fetch
	Name string
	// Partitions contains the partitions to fetch
	Partitions []PartitionSnapshot_FetchSnapshotRequest
}

func (t *TopicSnapshot_FetchSnapshotRequest) encode(pe packetEncoder, version int16) (err error) {
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

	pe.putUVarint(0)
	return nil
}

func (t *TopicSnapshot_FetchSnapshotRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionSnapshot_FetchSnapshotRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionSnapshot_FetchSnapshotRequest
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type FetchSnapshotRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the clusterId if known, this is used to validate metadata fetches prior to broker registration
	ClusterID *string
	// ReplicaID contains the broker ID of the follower
	ReplicaID int32
	// MaxBytes contains the maximum bytes to fetch from all of the snapshots
	MaxBytes int32
	// Topics contains the topics to fetch
	Topics []TopicSnapshot_FetchSnapshotRequest
}

func (r *FetchSnapshotRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ReplicaID)

	pe.putInt32(r.MaxBytes)

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *FetchSnapshotRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ReplicaID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.MaxBytes, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicSnapshot_FetchSnapshotRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicSnapshot_FetchSnapshotRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *FetchSnapshotRequest) GetKey() int16 {
	return 59
}

func (r *FetchSnapshotRequest) GetVersion() int16 {
	return r.Version
}

func (r *FetchSnapshotRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *FetchSnapshotRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *FetchSnapshotRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
