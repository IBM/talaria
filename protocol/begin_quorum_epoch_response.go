// protocol has been generated from message format json - DO NOT EDIT
package protocol

// PartitionData_BeginQuorumEpochResponse contains a
type PartitionData_BeginQuorumEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains a
	ErrorCode int16
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch
	LeaderEpoch int32
}

func (p *PartitionData_BeginQuorumEpochResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	return nil
}

func (p *PartitionData_BeginQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	return nil
}

// TopicData_BeginQuorumEpochResponse contains a
type TopicData_BeginQuorumEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains a
	Partitions []PartitionData_BeginQuorumEpochResponse
}

func (t *TopicData_BeginQuorumEpochResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.TopicName); err != nil {
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

func (t *TopicData_BeginQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]PartitionData_BeginQuorumEpochResponse, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block PartitionData_BeginQuorumEpochResponse
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	return nil
}

type BeginQuorumEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top level error code.
	ErrorCode int16
	// Topics contains a
	Topics []TopicData_BeginQuorumEpochResponse
}

func (r *BeginQuorumEpochResponse) encode(pe packetEncoder) (err error) {
	pe.putInt16(r.ErrorCode)

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

func (r *BeginQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]TopicData_BeginQuorumEpochResponse, numTopics)
	for i := 0; i < numTopics; i++ {
		var block TopicData_BeginQuorumEpochResponse
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	return nil
}

func (r *BeginQuorumEpochResponse) GetKey() int16 {
	return 53
}

func (r *BeginQuorumEpochResponse) GetVersion() int16 {
	return r.Version
}

func (r *BeginQuorumEpochResponse) GetHeaderVersion() int16 {
	return 0
}

func (r *BeginQuorumEpochResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *BeginQuorumEpochResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
