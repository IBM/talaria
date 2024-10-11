// protocol has been generated from message format json - DO NOT EDIT
package protocol

// ListOffsetsPartition contains each partition in the request.
type ListOffsetsPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// CurrentLeaderEpoch contains the current leader epoch.
	CurrentLeaderEpoch int32
	// Timestamp contains the current timestamp.
	Timestamp int64
	// MaxNumOffsets contains the maximum number of offsets to report.
	MaxNumOffsets int32
}

func (p *ListOffsetsPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	if p.Version >= 4 {
		pe.putInt32(p.CurrentLeaderEpoch)
	}

	pe.putInt64(p.Timestamp)

	if p.Version == 0 {
		pe.putInt32(p.MaxNumOffsets)
	}

	if p.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *ListOffsetsPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 4 {
		if p.CurrentLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Timestamp, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version == 0 {
		if p.MaxNumOffsets, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// ListOffsetsTopic contains each topic in the request.
type ListOffsetsTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains each partition in the request.
	Partitions []ListOffsetsPartition
}

func (t *ListOffsetsTopic) encode(pe packetEncoder, version int16) (err error) {
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

	if t.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *ListOffsetsTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]ListOffsetsPartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block ListOffsetsPartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	if t.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ListOffsetsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ReplicaID contains the broker ID of the requestor, or -1 if this request is being made by a normal consumer.
	ReplicaID int32
	// IsolationLevel contains a This setting controls the visibility of transactional records. Using READ_UNCOMMITTED (isolation_level = 0) makes all records visible. With READ_COMMITTED (isolation_level = 1), non-transactional and COMMITTED transactional records are visible. To be more concrete, READ_COMMITTED returns all data from offsets smaller than the current LSO (last stable offset), and enables the inclusion of the list of aborted transactions in the result, which allows consumers to discard ABORTED transactional records
	IsolationLevel int8
	// Topics contains each topic in the request.
	Topics []ListOffsetsTopic
}

func (r *ListOffsetsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ReplicaID)

	if r.Version >= 2 {
		pe.putInt8(r.IsolationLevel)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ListOffsetsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ReplicaID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if r.IsolationLevel, err = pd.getInt8(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]ListOffsetsTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block ListOffsetsTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ListOffsetsRequest) GetKey() int16 {
	return 2
}

func (r *ListOffsetsRequest) GetVersion() int16 {
	return r.Version
}

func (r *ListOffsetsRequest) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 2
	}
	return 1
}

func (r *ListOffsetsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *ListOffsetsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
