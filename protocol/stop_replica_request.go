// protocol has been generated from message format json - DO NOT EDIT
package protocol

// StopReplicaPartitionV0 contains the partitions to stop.
type StopReplicaPartitionV0 struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// PartitionIndex contains the partition index.
	PartitionIndex int32
}

func (u *StopReplicaPartitionV0) encode(pe packetEncoder, version int16) (err error) {
	u.Version = version
	if u.Version == 0 {
		if err := pe.putString(u.TopicName); err != nil {
			return err
		}
	}

	if u.Version == 0 {
		pe.putInt32(u.PartitionIndex)
	}

	if u.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (u *StopReplicaPartitionV0) decode(pd packetDecoder, version int16) (err error) {
	u.Version = version
	if u.Version == 0 {
		if u.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if u.Version == 0 {
		if u.PartitionIndex, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if u.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// StopReplicaTopicV1 contains the topics to stop.
type StopReplicaTopicV1 struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// PartitionIndexes contains the partition indexes.
	PartitionIndexes []int32
}

func (t *StopReplicaTopicV1) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 1 && t.Version <= 2 {
		if err := pe.putString(t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 1 && t.Version <= 2 {
		if err := pe.putInt32Array(t.PartitionIndexes); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *StopReplicaTopicV1) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 1 && t.Version <= 2 {
		if t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 1 && t.Version <= 2 {
		if t.PartitionIndexes, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// StopReplicaPartitionState contains the state of each partition
type StopReplicaPartitionState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// LeaderEpoch contains the leader epoch.
	LeaderEpoch int32
	// DeletePartition contains a Whether this partition should be deleted.
	DeletePartition bool
}

func (p *StopReplicaPartitionState) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 3 {
		pe.putInt32(p.PartitionIndex)
	}

	if p.Version >= 3 {
		pe.putInt32(p.LeaderEpoch)
	}

	if p.Version >= 3 {
		pe.putBool(p.DeletePartition)
	}

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *StopReplicaPartitionState) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 3 {
		if p.PartitionIndex, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 3 {
		if p.LeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 3 {
		if p.DeletePartition, err = pd.getBool(); err != nil {
			return err
		}
	}

	if p.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// StopReplicaTopicState contains each topic.
type StopReplicaTopicState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// PartitionStates contains the state of each partition
	PartitionStates []StopReplicaPartitionState
}

func (t *StopReplicaTopicState) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 3 {
		if err := pe.putString(t.TopicName); err != nil {
			return err
		}
	}

	if t.Version >= 3 {
		if err := pe.putArrayLength(len(t.PartitionStates)); err != nil {
			return err
		}
		for _, block := range t.PartitionStates {
			if err := block.encode(pe, t.Version); err != nil {
				return err
			}
		}
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *StopReplicaTopicState) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 3 {
		if t.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 3 {
		var numPartitionStates int
		if numPartitionStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		t.PartitionStates = make([]StopReplicaPartitionState, numPartitionStates)
		for i := 0; i < numPartitionStates; i++ {
			var block StopReplicaPartitionState
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.PartitionStates[i] = block
		}
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type StopReplicaRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ControllerID contains the controller id.
	ControllerID int32
	// isKRaftController contains a If KRaft controller id is used during migration. See KIP-866
	isKRaftController bool
	// ControllerEpoch contains the controller epoch.
	ControllerEpoch int32
	// BrokerEpoch contains the broker epoch.
	BrokerEpoch int64
	// DeletePartitions contains a Whether these partitions should be deleted.
	DeletePartitions bool
	// UngroupedPartitions contains the partitions to stop.
	UngroupedPartitions []StopReplicaPartitionV0
	// Topics contains the topics to stop.
	Topics []StopReplicaTopicV1
	// TopicStates contains each topic.
	TopicStates []StopReplicaTopicState
}

func (r *StopReplicaRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ControllerID)

	if r.Version >= 4 {
		pe.putBool(r.isKRaftController)
	}

	pe.putInt32(r.ControllerEpoch)

	if r.Version >= 1 {
		pe.putInt64(r.BrokerEpoch)
	}

	if r.Version >= 0 && r.Version <= 2 {
		pe.putBool(r.DeletePartitions)
	}

	if r.Version == 0 {
		if err := pe.putArrayLength(len(r.UngroupedPartitions)); err != nil {
			return err
		}
		for _, block := range r.UngroupedPartitions {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 1 && r.Version <= 2 {
		if err := pe.putArrayLength(len(r.Topics)); err != nil {
			return err
		}
		for _, block := range r.Topics {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 3 {
		if err := pe.putArrayLength(len(r.TopicStates)); err != nil {
			return err
		}
		for _, block := range r.TopicStates {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *StopReplicaRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ControllerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 4 {
		if r.isKRaftController, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.ControllerEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.BrokerEpoch, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 2 {
		if r.DeletePartitions, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version == 0 {
		var numUngroupedPartitions int
		if numUngroupedPartitions, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.UngroupedPartitions = make([]StopReplicaPartitionV0, numUngroupedPartitions)
		for i := 0; i < numUngroupedPartitions; i++ {
			var block StopReplicaPartitionV0
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.UngroupedPartitions[i] = block
		}
	}

	if r.Version >= 1 && r.Version <= 2 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Topics = make([]StopReplicaTopicV1, numTopics)
		for i := 0; i < numTopics; i++ {
			var block StopReplicaTopicV1
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 3 {
		var numTopicStates int
		if numTopicStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.TopicStates = make([]StopReplicaTopicState, numTopicStates)
		for i := 0; i < numTopicStates; i++ {
			var block StopReplicaTopicState
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.TopicStates[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *StopReplicaRequest) GetKey() int16 {
	return 5
}

func (r *StopReplicaRequest) GetVersion() int16 {
	return r.Version
}

func (r *StopReplicaRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *StopReplicaRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *StopReplicaRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
