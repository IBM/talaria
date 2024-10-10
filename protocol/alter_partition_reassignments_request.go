// protocol has been generated from message format json - DO NOT EDIT
package protocol

// ReassignablePartition contains the partitions to reassign.
type ReassignablePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// Replicas contains the replicas to place the partitions on, or null to cancel a pending reassignment for this partition.
	Replicas []int32
}

func (p *ReassignablePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	if err := pe.putInt32Array(p.Replicas); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *ReassignablePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Replicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// ReassignableTopic contains the topics to reassign.
type ReassignableTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the partitions to reassign.
	Partitions []ReassignablePartition
}

func (t *ReassignableTopic) encode(pe packetEncoder, version int16) (err error) {
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

func (t *ReassignableTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]ReassignablePartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block ReassignablePartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type AlterPartitionReassignmentsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TimeoutMs contains the time in ms to wait for the request to complete.
	TimeoutMs int32
	// Topics contains the topics to reassign.
	Topics []ReassignableTopic
}

func (r *AlterPartitionReassignmentsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.TimeoutMs)

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

func (r *AlterPartitionReassignmentsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]ReassignableTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block ReassignableTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *AlterPartitionReassignmentsRequest) GetKey() int16 {
	return 45
}

func (r *AlterPartitionReassignmentsRequest) GetVersion() int16 {
	return r.Version
}

func (r *AlterPartitionReassignmentsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *AlterPartitionReassignmentsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AlterPartitionReassignmentsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
