// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	"fmt"
	uuid "github.com/google/uuid"
	"time"
)

// MetadataResponseBroker contains each broker in the response.
type MetadataResponseBroker struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// NodeID contains the broker ID.
	NodeID int32
	// Host contains the broker hostname.
	Host string
	// Port contains the broker port.
	Port int32
	// Rack contains the rack of the broker, or null if it has not been assigned to a rack.
	Rack *string
}

func (b *MetadataResponseBroker) encode(pe packetEncoder, version int16) (err error) {
	b.Version = version
	pe.putInt32(b.NodeID)

	if err := pe.putString(b.Host); err != nil {
		return err
	}

	pe.putInt32(b.Port)

	if b.Version >= 1 {
		if err := pe.putNullableString(b.Rack); err != nil {
			return err
		}
	}

	if b.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (b *MetadataResponseBroker) decode(pd packetDecoder, version int16) (err error) {
	b.Version = version
	if b.NodeID, err = pd.getInt32(); err != nil {
		return err
	}

	if b.Host, err = pd.getString(); err != nil {
		return err
	}

	if b.Port, err = pd.getInt32(); err != nil {
		return err
	}

	if b.Version >= 1 {
		if b.Rack, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if b.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// MetadataResponsePartition contains each partition in the topic.
type MetadataResponsePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the partition error, or 0 if there was no error.
	ErrorCode int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// LeaderID contains the ID of the leader broker.
	LeaderID int32
	// LeaderEpoch contains the leader epoch of this partition.
	LeaderEpoch int32
	// ReplicaNodes contains the set of all nodes that host this partition.
	ReplicaNodes []int32
	// IsrNodes contains the set of nodes that are in sync with the leader for this partition.
	IsrNodes []int32
	// OfflineReplicas contains the set of offline replicas of this partition.
	OfflineReplicas []int32
}

func (p *MetadataResponsePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.LeaderID)

	if p.Version >= 7 {
		pe.putInt32(p.LeaderEpoch)
	}

	if err := pe.putInt32Array(p.ReplicaNodes); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.IsrNodes); err != nil {
		return err
	}

	if p.Version >= 5 {
		if err := pe.putInt32Array(p.OfflineReplicas); err != nil {
			return err
		}
	}

	if p.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *MetadataResponsePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 7 {
		if p.LeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.ReplicaNodes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.IsrNodes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.Version >= 5 {
		if p.OfflineReplicas, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if p.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// MetadataResponseTopic contains each topic in the response.
type MetadataResponseTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the topic error, or 0 if there was no error.
	ErrorCode int16
	// Name contains the topic name.
	Name *string
	// TopicID contains the topic id.
	TopicID uuid.UUID
	// IsInternal contains a True if the topic is internal.
	IsInternal bool
	// Partitions contains each partition in the topic.
	Partitions []MetadataResponsePartition
	// TopicAuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this topic.
	TopicAuthorizedOperations int32
}

func (t *MetadataResponseTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	pe.putInt16(t.ErrorCode)

	if t.Version >= 12 {
		if err := pe.putNullableString(t.Name); err != nil {
			return err
		}
	} else {
		if t.Name == nil {
			return fmt.Errorf("String field, Name, must not be nil in version %d", t.Version)
		}
		if err := pe.putString(*t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 10 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if t.Version >= 1 {
		pe.putBool(t.IsInternal)
	}

	if err := pe.putArrayLength(len(t.Partitions)); err != nil {
		return err
	}
	for _, block := range t.Partitions {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 8 {
		pe.putInt32(t.TopicAuthorizedOperations)
	}

	if t.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *MetadataResponseTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if t.Version >= 12 {
		if t.Name, err = pd.getNullableString(); err != nil {
			return err
		}
	} else {
		if *t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 10 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.Version >= 1 {
		if t.IsInternal, err = pd.getBool(); err != nil {
			return err
		}
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]MetadataResponsePartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block MetadataResponsePartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	if t.Version >= 8 {
		if t.TopicAuthorizedOperations, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if t.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type MetadataResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Brokers contains each broker in the response.
	Brokers []MetadataResponseBroker
	// ClusterID contains the cluster ID that responding broker belongs to.
	ClusterID *string
	// ControllerID contains the ID of the controller broker.
	ControllerID int32
	// Topics contains each topic in the response.
	Topics []MetadataResponseTopic
	// ClusterAuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this cluster.
	ClusterAuthorizedOperations int32
}

func (r *MetadataResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 9 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if err := pe.putArrayLength(len(r.Brokers)); err != nil {
		return err
	}
	for _, block := range r.Brokers {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		if err := pe.putNullableString(r.ClusterID); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putInt32(r.ControllerID)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 8 && r.Version <= 1 {
		pe.putInt32(r.ClusterAuthorizedOperations)
	}

	if r.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *MetadataResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 9 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numBrokers int
	if numBrokers, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Brokers = make([]MetadataResponseBroker, numBrokers)
	for i := 0; i < numBrokers; i++ {
		var block MetadataResponseBroker
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Brokers[i] = block
	}

	if r.Version >= 2 {
		if r.ClusterID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		if r.ControllerID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]MetadataResponseTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block MetadataResponseTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.Version >= 8 && r.Version <= 1 {
		if r.ClusterAuthorizedOperations, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *MetadataResponse) GetKey() int16 {
	return 3
}

func (r *MetadataResponse) GetVersion() int16 {
	return r.Version
}

func (r *MetadataResponse) GetHeaderVersion() int16 {
	if r.Version >= 9 {
		return 1
	}
	return 0
}

func (r *MetadataResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 12
}

func (r *MetadataResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *MetadataResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
