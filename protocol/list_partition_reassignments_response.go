// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// OngoingPartitionReassignment contains the ongoing reassignments for each partition.
type OngoingPartitionReassignment struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the index of the partition.
	PartitionIndex int32
	// Replicas contains the current replica set.
	Replicas []int32
	// AddingReplicas contains the set of replicas we are currently adding.
	AddingReplicas []int32
	// RemovingReplicas contains the set of replicas we are currently removing.
	RemovingReplicas []int32
}

func (p *OngoingPartitionReassignment) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	if err := pe.putInt32Array(p.Replicas); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.AddingReplicas); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.RemovingReplicas); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *OngoingPartitionReassignment) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Replicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.AddingReplicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.RemovingReplicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// OngoingTopicReassignment contains the ongoing reassignments for each topic.
type OngoingTopicReassignment struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the ongoing reassignments for each partition.
	Partitions []OngoingPartitionReassignment
}

func (t *OngoingTopicReassignment) encode(pe packetEncoder, version int16) (err error) {
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

func (t *OngoingTopicReassignment) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]OngoingPartitionReassignment, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block OngoingPartitionReassignment
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

type ListPartitionReassignmentsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level error code, or 0 if there was no error
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// Topics contains the ongoing reassignments for each topic.
	Topics []OngoingTopicReassignment
}

func (r *ListPartitionReassignmentsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
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

func (r *ListPartitionReassignmentsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]OngoingTopicReassignment, numTopics)
	for i := 0; i < numTopics; i++ {
		var block OngoingTopicReassignment
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

func (r *ListPartitionReassignmentsResponse) GetKey() int16 {
	return 46
}

func (r *ListPartitionReassignmentsResponse) GetVersion() int16 {
	return r.Version
}

func (r *ListPartitionReassignmentsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ListPartitionReassignmentsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ListPartitionReassignmentsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ListPartitionReassignmentsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
