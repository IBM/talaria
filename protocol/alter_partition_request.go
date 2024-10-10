// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// BrokerState contains a
type BrokerState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the ID of the broker.
	BrokerID int32
	// BrokerEpoch contains the epoch of the broker. It will be -1 if the epoch check is not supported.
	BrokerEpoch int64
}

func (n *BrokerState) encode(pe packetEncoder, version int16) (err error) {
	n.Version = version
	if n.Version >= 3 {
		pe.putInt32(n.BrokerID)
	}

	if n.Version >= 3 {
		pe.putInt64(n.BrokerEpoch)
	}

	pe.putUVarint(0)
	return nil
}

func (n *BrokerState) decode(pd packetDecoder, version int16) (err error) {
	n.Version = version
	if n.Version >= 3 {
		if n.BrokerID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if n.Version >= 3 {
		if n.BrokerEpoch, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// PartitionData_AlterPartitionRequest contains a
type PartitionData_AlterPartitionRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index
	PartitionIndex int32
	// LeaderEpoch contains the leader epoch of this partition
	LeaderEpoch int32
	// NewIsr contains the ISR for this partition. Deprecated since version 3.
	NewIsr []int32
	// NewIsrWithEpochs contains a
	NewIsrWithEpochs []BrokerState
	// LeaderRecoveryState contains a 1 if the partition is recovering from an unclean leader election; 0 otherwise.
	LeaderRecoveryState int8
	// PartitionEpoch contains the expected epoch of the partition which is being updated. For legacy cluster this is the ZkVersion in the LeaderAndIsr request.
	PartitionEpoch int32
}

func (p *PartitionData_AlterPartitionRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.LeaderEpoch)

	if p.Version >= 0 && p.Version <= 2 {
		if err := pe.putInt32Array(p.NewIsr); err != nil {
			return err
		}
	}

	if p.Version >= 3 {
		if err := pe.putArrayLength(len(p.NewIsrWithEpochs)); err != nil {
			return err
		}
		for _, block := range p.NewIsrWithEpochs {
			if err := block.encode(pe, p.Version); err != nil {
				return err
			}
		}
	}

	if p.Version >= 1 {
		pe.putInt8(p.LeaderRecoveryState)
	}

	pe.putInt32(p.PartitionEpoch)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_AlterPartitionRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 0 && p.Version <= 2 {
		if p.NewIsr, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if p.Version >= 3 {
		var numNewIsrWithEpochs int
		if numNewIsrWithEpochs, err = pd.getArrayLength(); err != nil {
			return err
		}
		p.NewIsrWithEpochs = make([]BrokerState, numNewIsrWithEpochs)
		for i := 0; i < numNewIsrWithEpochs; i++ {
			var block BrokerState
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.NewIsrWithEpochs[i] = block
		}
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

// TopicData_AlterPartitionRequest contains a
type TopicData_AlterPartitionRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the name of the topic to alter ISRs for
	TopicName string
	// TopicID contains the ID of the topic to alter ISRs for
	TopicID uuid.UUID
	// Partitions contains a
	Partitions []PartitionData_AlterPartitionRequest
}

func (t *TopicData_AlterPartitionRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicData_AlterPartitionRequest) decode(pd packetDecoder, version int16) (err error) {
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
	t.Partitions = make([]PartitionData_AlterPartitionRequest, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block PartitionData_AlterPartitionRequest
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

type AlterPartitionRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the ID of the requesting broker
	BrokerID int32
	// BrokerEpoch contains the epoch of the requesting broker
	BrokerEpoch int64
	// Topics contains a
	Topics []TopicData_AlterPartitionRequest
}

func (r *AlterPartitionRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.BrokerID)

	pe.putInt64(r.BrokerEpoch)

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

func (r *AlterPartitionRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.BrokerEpoch, err = pd.getInt64(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]TopicData_AlterPartitionRequest, numTopics)
	for i := 0; i < numTopics; i++ {
		var block TopicData_AlterPartitionRequest
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

func (r *AlterPartitionRequest) GetKey() int16 {
	return 56
}

func (r *AlterPartitionRequest) GetVersion() int16 {
	return r.Version
}

func (r *AlterPartitionRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *AlterPartitionRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *AlterPartitionRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
