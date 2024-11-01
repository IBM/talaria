// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// OffsetDeleteResponsePartition contains the responses for each partition in the topic.
type OffsetDeleteResponsePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (p *OffsetDeleteResponsePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	return nil
}

func (p *OffsetDeleteResponsePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	return nil
}

// OffsetDeleteResponseTopic contains the responses for each topic.
type OffsetDeleteResponseTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the responses for each partition in the topic.
	Partitions []OffsetDeleteResponsePartition
}

func (t *OffsetDeleteResponseTopic) encode(pe packetEncoder, version int16) (err error) {
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

	return nil
}

func (t *OffsetDeleteResponseTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]OffsetDeleteResponsePartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block OffsetDeleteResponsePartition
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	return nil
}

type OffsetDeleteResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top-level error code, or 0 if there was no error.
	ErrorCode int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains the responses for each topic.
	Topics []OffsetDeleteResponseTopic
}

func (r *OffsetDeleteResponse) encode(pe packetEncoder) (err error) {
	pe.putInt16(r.ErrorCode)

	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	return nil
}

func (r *OffsetDeleteResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]OffsetDeleteResponseTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block OffsetDeleteResponseTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	return nil
}

func (r *OffsetDeleteResponse) GetKey() int16 {
	return 47
}

func (r *OffsetDeleteResponse) GetVersion() int16 {
	return r.Version
}

func (r *OffsetDeleteResponse) GetHeaderVersion() int16 {
	return 0
}

func (r *OffsetDeleteResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *OffsetDeleteResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *OffsetDeleteResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
