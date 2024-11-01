// protocol has been generated from message format json - DO NOT EDIT
package protocol

// ListPartitionReassignmentsTopics contains the topics to list partition reassignments for, or null to list everything.
type ListPartitionReassignmentsTopics struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name string
	// PartitionIndexes contains the partitions to list partition reassignments for.
	PartitionIndexes []int32
}

func (t *ListPartitionReassignmentsTopics) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.PartitionIndexes); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *ListPartitionReassignmentsTopics) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.PartitionIndexes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ListPartitionReassignmentsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TimeoutMs contains the time in ms to wait for the request to complete.
	TimeoutMs int32
	// Topics contains the topics to list partition reassignments for, or null to list everything.
	Topics []ListPartitionReassignmentsTopics
}

func (r *ListPartitionReassignmentsRequest) encode(pe packetEncoder) (err error) {
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

func (r *ListPartitionReassignmentsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]ListPartitionReassignmentsTopics, numTopics)
		for i := 0; i < numTopics; i++ {
			var block ListPartitionReassignmentsTopics
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

func (r *ListPartitionReassignmentsRequest) GetKey() int16 {
	return 46
}

func (r *ListPartitionReassignmentsRequest) GetVersion() int16 {
	return r.Version
}

func (r *ListPartitionReassignmentsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ListPartitionReassignmentsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ListPartitionReassignmentsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
