// protocol has been generated from message format json - DO NOT EDIT
package protocol

// PartitionData_VoteRequest contains a
type PartitionData_VoteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// CandidateEpoch contains the bumped epoch of the candidate sending the request
	CandidateEpoch int32
	// CandidateID contains the ID of the voter sending the request
	CandidateID int32
	// LastOffsetEpoch contains the epoch of the last record written to the metadata log
	LastOffsetEpoch int32
	// LastOffset contains the offset of the last record written to the metadata log
	LastOffset int64
}

func (p *PartitionData_VoteRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.CandidateEpoch)

	pe.putInt32(p.CandidateID)

	pe.putInt32(p.LastOffsetEpoch)

	pe.putInt64(p.LastOffset)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_VoteRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.CandidateEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.CandidateID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LastOffsetEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LastOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_VoteRequest contains a
type TopicData_VoteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains a
	Partitions []PartitionData_VoteRequest
}

func (t *TopicData_VoteRequest) encode(pe packetEncoder, version int16) (err error) {
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

	pe.putUVarint(0)
	return nil
}

func (t *TopicData_VoteRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_VoteRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_VoteRequest
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

type VoteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains a
	ClusterID *string
	// Topics contains a
	Topics []TopicData_VoteRequest
}

func (r *VoteRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.ClusterID); err != nil {
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

	pe.putUVarint(0)
	return nil
}

func (r *VoteRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_VoteRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_VoteRequest
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

func (r *VoteRequest) GetKey() int16 {
	return 52
}

func (r *VoteRequest) GetVersion() int16 {
	return r.Version
}

func (r *VoteRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *VoteRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *VoteRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
