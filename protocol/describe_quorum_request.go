// protocol has been generated from message format json - DO NOT EDIT
package protocol

// PartitionData_DescribeQuorumRequest contains a
type PartitionData_DescribeQuorumRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
}

func (p *PartitionData_DescribeQuorumRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_DescribeQuorumRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_DescribeQuorumRequest contains a
type TopicData_DescribeQuorumRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains a
	Partitions []PartitionData_DescribeQuorumRequest
}

func (t *TopicData_DescribeQuorumRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicData_DescribeQuorumRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_DescribeQuorumRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_DescribeQuorumRequest
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

type DescribeQuorumRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains a
	Topics []TopicData_DescribeQuorumRequest
}

func (r *DescribeQuorumRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
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

func (r *DescribeQuorumRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_DescribeQuorumRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_DescribeQuorumRequest
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

func (r *DescribeQuorumRequest) GetKey() int16 {
	return 55
}

func (r *DescribeQuorumRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeQuorumRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeQuorumRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *DescribeQuorumRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
