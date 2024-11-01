// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// UpdateMetadataPartitionState_UpdateMetadataRequest contains a
type UpdateMetadataPartitionState_UpdateMetadataRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains in older versions of this RPC, the topic name.
	TopicName string
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ControllerEpoch contains the controller epoch.
	ControllerEpoch int32
	// Leader contains the ID of the broker which is the current partition leader.
	Leader int32
	// LeaderEpoch contains the leader epoch of this partition.
	LeaderEpoch int32
	// Isr contains the brokers which are in the ISR for this partition.
	Isr []int32
	// ZkVersion contains the Zookeeper version.
	ZkVersion int32
	// Replicas contains a All the replicas of this partition.
	Replicas []int32
	// OfflineReplicas contains the replicas of this partition which are offline.
	OfflineReplicas []int32
}

func (u *UpdateMetadataPartitionState_UpdateMetadataRequest) encode(pe packetEncoder, version int16) (err error) {
	u.Version = version
	if u.Version >= 0 && u.Version <= 4 {
		if err := pe.putString(u.TopicName); err != nil {
			return err
		}
	}

	pe.putInt32(u.PartitionIndex)

	pe.putInt32(u.ControllerEpoch)

	pe.putInt32(u.Leader)

	pe.putInt32(u.LeaderEpoch)

	if err := pe.putInt32Array(u.Isr); err != nil {
		return err
	}

	pe.putInt32(u.ZkVersion)

	if err := pe.putInt32Array(u.Replicas); err != nil {
		return err
	}

	if u.Version >= 4 {
		if err := pe.putInt32Array(u.OfflineReplicas); err != nil {
			return err
		}
	}

	if u.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (u *UpdateMetadataPartitionState_UpdateMetadataRequest) decode(pd packetDecoder, version int16) (err error) {
	u.Version = version
	if u.Version >= 0 && u.Version <= 4 {
		if u.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if u.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if u.ControllerEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if u.Leader, err = pd.getInt32(); err != nil {
		return err
	}

	if u.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if u.Isr, err = pd.getInt32Array(); err != nil {
		return err
	}

	if u.ZkVersion, err = pd.getInt32(); err != nil {
		return err
	}

	if u.Replicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if u.Version >= 4 {
		if u.OfflineReplicas, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if u.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// UpdateMetadataTopicState contains in newer versions of this RPC, each topic that we would like to update.
type UpdateMetadataTopicState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// TopicID contains the topic id.
	TopicID uuid.UUID
	// PartitionStates contains the partition that we would like to update.
	PartitionStates []UpdateMetadataPartitionState_UpdateMetadataRequest
}

func (t *UpdateMetadataTopicState) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 5 {
		if err := pe.putString(t.TopicName); err != nil {
			return err
		}
	}

	if t.Version >= 7 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		if err := pe.putArrayLength(len(t.PartitionStates)); err != nil {
			return err
		}
		for _, block := range t.PartitionStates {
			if err := block.encode(pe, t.Version); err != nil {
				return err
			}
		}
	}

	if t.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *UpdateMetadataTopicState) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 5 {
		if t.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 7 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		var numPartitionStates int
		if numPartitionStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numPartitionStates > 0 {
			t.PartitionStates = make([]UpdateMetadataPartitionState_UpdateMetadataRequest, numPartitionStates)
			for i := 0; i < numPartitionStates; i++ {
				var block UpdateMetadataPartitionState_UpdateMetadataRequest
				if err := block.decode(pd, t.Version); err != nil {
					return err
				}
				t.PartitionStates[i] = block
			}
		}
	}

	if t.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// UpdateMetadataEndpoint contains the broker endpoints.
type UpdateMetadataEndpoint struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Port contains the port of this endpoint
	Port int32
	// Host contains the hostname of this endpoint
	Host string
	// Listener contains the listener name.
	Listener string
	// SecurityProtocol contains the security protocol type.
	SecurityProtocol int16
}

func (e *UpdateMetadataEndpoint) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if e.Version >= 1 {
		pe.putInt32(e.Port)
	}

	if e.Version >= 1 {
		if err := pe.putString(e.Host); err != nil {
			return err
		}
	}

	if e.Version >= 3 {
		if err := pe.putString(e.Listener); err != nil {
			return err
		}
	}

	if e.Version >= 1 {
		pe.putInt16(e.SecurityProtocol)
	}

	if e.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (e *UpdateMetadataEndpoint) decode(pd packetDecoder, version int16) (err error) {
	e.Version = version
	if e.Version >= 1 {
		if e.Port, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if e.Version >= 1 {
		if e.Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if e.Version >= 3 {
		if e.Listener, err = pd.getString(); err != nil {
			return err
		}
	}

	if e.Version >= 1 {
		if e.SecurityProtocol, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if e.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// UpdateMetadataBroker contains a
type UpdateMetadataBroker struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ID contains the broker id.
	ID int32
	// V0Host contains the broker hostname.
	V0Host string
	// V0Port contains the broker port.
	V0Port int32
	// Endpoints contains the broker endpoints.
	Endpoints []UpdateMetadataEndpoint
	// Rack contains the rack which this broker belongs to.
	Rack *string
}

func (l *UpdateMetadataBroker) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	pe.putInt32(l.ID)

	if l.Version == 0 {
		if err := pe.putString(l.V0Host); err != nil {
			return err
		}
	}

	if l.Version == 0 {
		pe.putInt32(l.V0Port)
	}

	if l.Version >= 1 {
		if err := pe.putArrayLength(len(l.Endpoints)); err != nil {
			return err
		}
		for _, block := range l.Endpoints {
			if err := block.encode(pe, l.Version); err != nil {
				return err
			}
		}
	}

	if l.Version >= 2 {
		if err := pe.putNullableString(l.Rack); err != nil {
			return err
		}
	}

	if l.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (l *UpdateMetadataBroker) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.ID, err = pd.getInt32(); err != nil {
		return err
	}

	if l.Version == 0 {
		if l.V0Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.Version == 0 {
		if l.V0Port, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		var numEndpoints int
		if numEndpoints, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numEndpoints > 0 {
			l.Endpoints = make([]UpdateMetadataEndpoint, numEndpoints)
			for i := 0; i < numEndpoints; i++ {
				var block UpdateMetadataEndpoint
				if err := block.decode(pd, l.Version); err != nil {
					return err
				}
				l.Endpoints[i] = block
			}
		}
	}

	if l.Version >= 2 {
		if l.Rack, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if l.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMetadataRequest struct {
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
	// UngroupedPartitionStates contains in older versions of this RPC, each partition that we would like to update.
	UngroupedPartitionStates []UpdateMetadataPartitionState_UpdateMetadataRequest
	// TopicStates contains in newer versions of this RPC, each topic that we would like to update.
	TopicStates []UpdateMetadataTopicState
	// LiveBrokers contains a
	LiveBrokers []UpdateMetadataBroker
}

func (r *UpdateMetadataRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ControllerID)

	if r.Version >= 8 {
		pe.putBool(r.isKRaftController)
	}

	pe.putInt32(r.ControllerEpoch)

	if r.Version >= 5 {
		pe.putInt64(r.BrokerEpoch)
	}

	if r.Version >= 0 && r.Version <= 4 {
		if err := pe.putArrayLength(len(r.UngroupedPartitionStates)); err != nil {
			return err
		}
		for _, block := range r.UngroupedPartitionStates {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 5 {
		if err := pe.putArrayLength(len(r.TopicStates)); err != nil {
			return err
		}
		for _, block := range r.TopicStates {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if err := pe.putArrayLength(len(r.LiveBrokers)); err != nil {
		return err
	}
	for _, block := range r.LiveBrokers {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *UpdateMetadataRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ControllerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 8 {
		if r.isKRaftController, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.ControllerEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 5 {
		if r.BrokerEpoch, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 4 {
		var numUngroupedPartitionStates int
		if numUngroupedPartitionStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numUngroupedPartitionStates > 0 {
			r.UngroupedPartitionStates = make([]UpdateMetadataPartitionState_UpdateMetadataRequest, numUngroupedPartitionStates)
			for i := 0; i < numUngroupedPartitionStates; i++ {
				var block UpdateMetadataPartitionState_UpdateMetadataRequest
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.UngroupedPartitionStates[i] = block
			}
		}
	}

	if r.Version >= 5 {
		var numTopicStates int
		if numTopicStates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopicStates > 0 {
			r.TopicStates = make([]UpdateMetadataTopicState, numTopicStates)
			for i := 0; i < numTopicStates; i++ {
				var block UpdateMetadataTopicState
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.TopicStates[i] = block
			}
		}
	}

	var numLiveBrokers int
	if numLiveBrokers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numLiveBrokers > 0 {
		r.LiveBrokers = make([]UpdateMetadataBroker, numLiveBrokers)
		for i := 0; i < numLiveBrokers; i++ {
			var block UpdateMetadataBroker
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.LiveBrokers[i] = block
		}
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *UpdateMetadataRequest) GetKey() int16 {
	return 6
}

func (r *UpdateMetadataRequest) GetVersion() int16 {
	return r.Version
}

func (r *UpdateMetadataRequest) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 2
	}
	return 1
}

func (r *UpdateMetadataRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *UpdateMetadataRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
