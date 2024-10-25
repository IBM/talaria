// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// LeaderAndIsrPartitionState_LeaderAndIsrRequest contains a
type LeaderAndIsrPartitionState_LeaderAndIsrRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.  This is only present in v0 or v1.
	TopicName string
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ControllerEpoch contains the controller epoch.
	ControllerEpoch int32
	// Leader contains the broker ID of the leader.
	Leader int32
	// LeaderEpoch contains the leader epoch.
	LeaderEpoch int32
	// Isr contains the in-sync replica IDs.
	Isr []int32
	// PartitionEpoch contains the current epoch for the partition. The epoch is a monotonically increasing value which is incremented after every partition change. (Since the LeaderAndIsr request is only used by the legacy controller, this corresponds to the zkVersion)
	PartitionEpoch int32
	// Replicas contains the replica IDs.
	Replicas []int32
	// AddingReplicas contains the replica IDs that we are adding this partition to, or null if no replicas are being added.
	AddingReplicas []int32
	// RemovingReplicas contains the replica IDs that we are removing this partition from, or null if no replicas are being removed.
	RemovingReplicas []int32
	// IsNew contains a Whether the replica should have existed on the broker or not.
	IsNew bool
	// LeaderRecoveryState contains a 1 if the partition is recovering from an unclean leader election; 0 otherwise.
	LeaderRecoveryState int8
}

func (l *LeaderAndIsrPartitionState_LeaderAndIsrRequest) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 0 && l.Version <= 1 {
		if err := pe.putString(l.TopicName); err != nil {
			return err
		}
	}

	pe.putInt32(l.PartitionIndex)

	pe.putInt32(l.ControllerEpoch)

	pe.putInt32(l.Leader)

	pe.putInt32(l.LeaderEpoch)

	if err := pe.putInt32Array(l.Isr); err != nil {
		return err
	}

	pe.putInt32(l.PartitionEpoch)

	if err := pe.putInt32Array(l.Replicas); err != nil {
		return err
	}

	if l.Version >= 3 {
		if err := pe.putInt32Array(l.AddingReplicas); err != nil {
			return err
		}
	}

	if l.Version >= 3 {
		if err := pe.putInt32Array(l.RemovingReplicas); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		pe.putBool(l.IsNew)
	}

	if l.Version >= 6 {
		pe.putInt8(l.LeaderRecoveryState)
	}

	if l.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (l *LeaderAndIsrPartitionState_LeaderAndIsrRequest) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 0 && l.Version <= 1 {
		if l.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if l.ControllerEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if l.Leader, err = pd.getInt32(); err != nil {
		return err
	}

	if l.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if l.Isr, err = pd.getInt32Array(); err != nil {
		return err
	}

	if l.PartitionEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if l.Replicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if l.Version >= 3 {
		if l.AddingReplicas, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if l.Version >= 3 {
		if l.RemovingReplicas, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		if l.IsNew, err = pd.getBool(); err != nil {
			return err
		}
	}

	if l.Version >= 6 {
		if l.LeaderRecoveryState, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if l.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// LeaderAndIsrTopicState contains each topic.
type LeaderAndIsrTopicState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// PartitionStates contains the state of each partition
	PartitionStates []LeaderAndIsrPartitionState_LeaderAndIsrRequest
}

func (t *LeaderAndIsrTopicState) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 2 {
		if err := pe.putString(t.TopicName); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		if err := pe.putArrayLength(len(t.PartitionStates)); err != nil {
			return err
		}
		for _, block := range t.PartitionStates {
			if err := block.encode(pe, t.Version); err != nil {
				return err
			}
		}
	}

	if t.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *LeaderAndIsrTopicState) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 2 {
		if t.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		var numPartitionStates int
		if numPartitionStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numPartitionStates > 0 {
			t.PartitionStates = make([]LeaderAndIsrPartitionState_LeaderAndIsrRequest, numPartitionStates)
			for i := 0; i < numPartitionStates; i++ {
				var block LeaderAndIsrPartitionState_LeaderAndIsrRequest
				if err := block.decode(pd, t.Version); err != nil {
					return err
				}
				t.PartitionStates[i] = block
			}
		}
	}

	if t.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// LeaderAndIsrLiveLeader contains the current live leaders.
type LeaderAndIsrLiveLeader struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the leader's broker ID.
	BrokerID int32
	// HostName contains the leader's hostname.
	HostName string
	// Port contains the leader's port.
	Port int32
}

func (l *LeaderAndIsrLiveLeader) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	pe.putInt32(l.BrokerID)

	if err := pe.putString(l.HostName); err != nil {
		return err
	}

	pe.putInt32(l.Port)

	if l.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (l *LeaderAndIsrLiveLeader) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if l.HostName, err = pd.getString(); err != nil {
		return err
	}

	if l.Port, err = pd.getInt32(); err != nil {
		return err
	}

	if l.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type LeaderAndIsrRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ControllerID contains the current controller ID.
	ControllerID int32
	// isKRaftController contains a If KRaft controller id is used during migration. See KIP-866
	isKRaftController bool
	// ControllerEpoch contains the current controller epoch.
	ControllerEpoch int32
	// BrokerEpoch contains the current broker epoch.
	BrokerEpoch int64
	// Type contains the type that indicates whether all topics are included in the request
	Type int8
	// UngroupedPartitionStates contains the state of each partition, in a v0 or v1 message.
	UngroupedPartitionStates []LeaderAndIsrPartitionState_LeaderAndIsrRequest
	// TopicStates contains each topic.
	TopicStates []LeaderAndIsrTopicState
	// LiveLeaders contains the current live leaders.
	LiveLeaders []LeaderAndIsrLiveLeader
}

func (r *LeaderAndIsrRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ControllerID)

	if r.Version >= 7 {
		pe.putBool(r.isKRaftController)
	}

	pe.putInt32(r.ControllerEpoch)

	if r.Version >= 2 {
		pe.putInt64(r.BrokerEpoch)
	}

	if r.Version >= 5 {
		pe.putInt8(r.Type)
	}

	if r.Version >= 0 && r.Version <= 1 {
		if err := pe.putArrayLength(len(r.UngroupedPartitionStates)); err != nil {
			return err
		}
		for _, block := range r.UngroupedPartitionStates {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 2 {
		if err := pe.putArrayLength(len(r.TopicStates)); err != nil {
			return err
		}
		for _, block := range r.TopicStates {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if err := pe.putArrayLength(len(r.LiveLeaders)); err != nil {
		return err
	}
	for _, block := range r.LiveLeaders {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *LeaderAndIsrRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ControllerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 7 {
		if r.isKRaftController, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.ControllerEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if r.BrokerEpoch, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		if r.Type, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 1 {
		var numUngroupedPartitionStates int
		if numUngroupedPartitionStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numUngroupedPartitionStates > 0 {
			r.UngroupedPartitionStates = make([]LeaderAndIsrPartitionState_LeaderAndIsrRequest, numUngroupedPartitionStates)
			for i := 0; i < numUngroupedPartitionStates; i++ {
				var block LeaderAndIsrPartitionState_LeaderAndIsrRequest
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.UngroupedPartitionStates[i] = block
			}
		}
	}

	if r.Version >= 2 {
		var numTopicStates int
		if numTopicStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopicStates > 0 {
			r.TopicStates = make([]LeaderAndIsrTopicState, numTopicStates)
			for i := 0; i < numTopicStates; i++ {
				var block LeaderAndIsrTopicState
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.TopicStates[i] = block
			}
		}
	}

	var numLiveLeaders int
	if numLiveLeaders, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numLiveLeaders > 0 {
		r.LiveLeaders = make([]LeaderAndIsrLiveLeader, numLiveLeaders)
		for i := 0; i < numLiveLeaders; i++ {
			var block LeaderAndIsrLiveLeader
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.LiveLeaders[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *LeaderAndIsrRequest) GetKey() int16 {
	return 4
}

func (r *LeaderAndIsrRequest) GetVersion() int16 {
	return r.Version
}

func (r *LeaderAndIsrRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *LeaderAndIsrRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 7
}

func (r *LeaderAndIsrRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
