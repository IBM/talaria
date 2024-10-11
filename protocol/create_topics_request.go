// protocol has been generated from message format json - DO NOT EDIT
package protocol

// CreatableReplicaAssignment contains the manual partition assignment, or the empty array if we are using automatic assignment.
type CreatableReplicaAssignment struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// BrokerIds contains the brokers to place the partition on.
	BrokerIds []int32
}

func (a *CreatableReplicaAssignment) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt32(a.PartitionIndex)

	if err := pe.putInt32Array(a.BrokerIds); err != nil {
		return err
	}

	if a.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *CreatableReplicaAssignment) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if a.BrokerIds, err = pd.getInt32Array(); err != nil {
		return err
	}

	if a.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// CreateableTopicConfig contains the custom topic configurations to set.
type CreateableTopicConfig struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the configuration name.
	Name string
	// Value contains the configuration value.
	Value *string
}

func (c *CreateableTopicConfig) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.Name); err != nil {
		return err
	}

	if err := pe.putNullableString(c.Value); err != nil {
		return err
	}

	if c.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *CreateableTopicConfig) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Name, err = pd.getString(); err != nil {
		return err
	}

	if c.Value, err = pd.getNullableString(); err != nil {
		return err
	}

	if c.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// CreatableTopic contains the topics to create.
type CreatableTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// NumPartitions contains the number of partitions to create in the topic, or -1 if we are either specifying a manual partition assignment or using the default partitions.
	NumPartitions int32
	// ReplicationFactor contains the number of replicas to create for each partition in the topic, or -1 if we are either specifying a manual partition assignment or using the default replication factor.
	ReplicationFactor int16
	// Assignments contains the manual partition assignment, or the empty array if we are using automatic assignment.
	Assignments []CreatableReplicaAssignment
	// Configs contains the custom topic configurations to set.
	Configs []CreateableTopicConfig
}

func (t *CreatableTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	pe.putInt32(t.NumPartitions)

	pe.putInt16(t.ReplicationFactor)

	if err := pe.putArrayLength(len(t.Assignments)); err != nil {
		return err
	}
	for _, block := range t.Assignments {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(t.Configs)); err != nil {
		return err
	}
	for _, block := range t.Configs {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *CreatableTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.NumPartitions, err = pd.getInt32(); err != nil {
		return err
	}

	if t.ReplicationFactor, err = pd.getInt16(); err != nil {
		return err
	}

	var numAssignments int
	if numAssignments, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Assignments = make([]CreatableReplicaAssignment, numAssignments)
	for i := 0; i < numAssignments; i++ {
		var block CreatableReplicaAssignment
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Assignments[i] = block
	}

	var numConfigs int
	if numConfigs, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Configs = make([]CreateableTopicConfig, numConfigs)
	for i := 0; i < numConfigs; i++ {
		var block CreateableTopicConfig
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Configs[i] = block
	}

	if t.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTopicsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains the topics to create.
	Topics []CreatableTopic
	// timeoutMs contains a How long to wait in milliseconds before timing out the request.
	timeoutMs int32
	// validateOnly contains a If true, check that the topics can be created as specified, but don't create anything.
	validateOnly bool
}

func (r *CreateTopicsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 5 {
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

	pe.putInt32(r.timeoutMs)

	if r.Version >= 1 {
		pe.putBool(r.validateOnly)
	}

	if r.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreateTopicsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 5 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]CreatableTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block CreatableTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.timeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.validateOnly, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateTopicsRequest) GetKey() int16 {
	return 19
}

func (r *CreateTopicsRequest) GetVersion() int16 {
	return r.Version
}

func (r *CreateTopicsRequest) GetHeaderVersion() int16 {
	if r.Version >= 5 {
		return 2
	}
	return 1
}

func (r *CreateTopicsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 7
}

func (r *CreateTopicsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
