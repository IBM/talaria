// protocol has been generated from message format json - DO NOT EDIT
package protocol

// CreatePartitionsAssignment contains the new partition assignments.
type CreatePartitionsAssignment struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerIds contains the assigned broker IDs.
	BrokerIds []int32
}

func (a *CreatePartitionsAssignment) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putInt32Array(a.BrokerIds); err != nil {
		return err
	}

	if a.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *CreatePartitionsAssignment) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.BrokerIds, err = pd.getInt32Array(); err != nil {
		return err
	}

	if a.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// CreatePartitionsTopic contains each topic that we want to create new partitions inside.
type CreatePartitionsTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Count contains the new partition count.
	Count int32
	// Assignments contains the new partition assignments.
	Assignments []CreatePartitionsAssignment
}

func (t *CreatePartitionsTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	pe.putInt32(t.Count)

	if err := pe.putArrayLength(len(t.Assignments)); err != nil {
		return err
	}
	for _, block := range t.Assignments {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *CreatePartitionsTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.Count, err = pd.getInt32(); err != nil {
		return err
	}

	var numAssignments int
	if numAssignments, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Assignments = make([]CreatePartitionsAssignment, numAssignments)
	for i := 0; i < numAssignments; i++ {
		var block CreatePartitionsAssignment
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Assignments[i] = block
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePartitionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains each topic that we want to create new partitions inside.
	Topics []CreatePartitionsTopic
	// TimeoutMs contains the time in ms to wait for the partitions to be created.
	TimeoutMs int32
	// ValidateOnly contains a If true, then validate the request, but don't actually increase the number of partitions.
	ValidateOnly bool
}

func (r *CreatePartitionsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putInt32(r.TimeoutMs)

	pe.putBool(r.ValidateOnly)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreatePartitionsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]CreatePartitionsTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block CreatePartitionsTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ValidateOnly, err = pd.getBool(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreatePartitionsRequest) GetKey() int16 {
	return 37
}

func (r *CreatePartitionsRequest) GetVersion() int16 {
	return r.Version
}

func (r *CreatePartitionsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *CreatePartitionsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *CreatePartitionsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
