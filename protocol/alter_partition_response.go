// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// PartitionData_AlterPartitionResponse contains a
type PartitionData_AlterPartitionResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index
	PartitionIndex int32
	// ErrorCode contains the partition level error code
	ErrorCode int16
	// LeaderID contains the broker ID of the leader.
	LeaderID int32
	// LeaderEpoch contains the leader epoch.
	LeaderEpoch int32
	// Isr contains the in-sync replica IDs.
	Isr []int32
	// LeaderRecoveryState contains a 1 if the partition is recovering from an unclean leader election; 0 otherwise.
	LeaderRecoveryState int8
	// PartitionEpoch contains the current epoch for the partition for KRaft controllers. The current ZK version for the legacy controllers.
	PartitionEpoch int32
}

func (p *PartitionData_AlterPartitionResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	if err := pe.putInt32Array(p.Isr); err != nil {
		return err
	}

	if p.Version >= 1 {
		pe.putInt8(p.LeaderRecoveryState)
	}

	pe.putInt32(p.PartitionEpoch)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_AlterPartitionResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Isr, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.Version >= 1 {
		if p.LeaderRecoveryState, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if p.PartitionEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_AlterPartitionResponse contains a
type TopicData_AlterPartitionResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the name of the topic
	TopicName string
	// TopicID contains the ID of the topic
	TopicID uuid.UUID
	// Partitions contains a
	Partitions []PartitionData_AlterPartitionResponse
}

func (t *TopicData_AlterPartitionResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 1 {
		if err := pe.putString(t.TopicName); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
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

func (t *TopicData_AlterPartitionResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 1 {
		if t.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]PartitionData_AlterPartitionResponse, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block PartitionData_AlterPartitionResponse
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

type AlterPartitionResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code
	ErrorCode int16
	// Topics contains a
	Topics []TopicData_AlterPartitionResponse
}

func (r *AlterPartitionResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

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

func (r *AlterPartitionResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]TopicData_AlterPartitionResponse, numTopics)
	for i := 0; i < numTopics; i++ {
		var block TopicData_AlterPartitionResponse
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

func (r *AlterPartitionResponse) GetKey() int16 {
	return 56
}

func (r *AlterPartitionResponse) GetVersion() int16 {
	return r.Version
}

func (r *AlterPartitionResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *AlterPartitionResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *AlterPartitionResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AlterPartitionResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
