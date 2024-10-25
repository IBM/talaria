// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// EpochEndOffset_OffsetForLeaderEpochResponse contains each partition in the topic we fetched offsets for.
type EpochEndOffset_OffsetForLeaderEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code 0, or if there was no error.
	ErrorCode int16
	// Partition contains the partition index.
	Partition int32
	// LeaderEpoch contains the leader epoch of the partition.
	LeaderEpoch int32
	// EndOffset contains the end offset of the epoch.
	EndOffset int64
}

func (p *EpochEndOffset_OffsetForLeaderEpochResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.Partition)

	if p.Version >= 1 {
		pe.putInt32(p.LeaderEpoch)
	}

	pe.putInt64(p.EndOffset)

	if p.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *EpochEndOffset_OffsetForLeaderEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 1 {
		if p.LeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.EndOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetForLeaderTopicResult contains each topic we fetched offsets for.
type OffsetForLeaderTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the topic name.
	Topic string
	// Partitions contains each partition in the topic we fetched offsets for.
	Partitions []EpochEndOffset_OffsetForLeaderEpochResponse
}

func (t *OffsetForLeaderTopicResult) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Topic); err != nil {
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

	if t.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *OffsetForLeaderTopicResult) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Topic, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]EpochEndOffset_OffsetForLeaderEpochResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block EpochEndOffset_OffsetForLeaderEpochResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type OffsetForLeaderEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains each topic we fetched offsets for.
	Topics []OffsetForLeaderTopicResult
}

func (r *OffsetForLeaderEpochResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 2 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *OffsetForLeaderEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 2 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]OffsetForLeaderTopicResult, numTopics)
		for i := 0; i < numTopics; i++ {
			var block OffsetForLeaderTopicResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *OffsetForLeaderEpochResponse) GetKey() int16 {
	return 23
}

func (r *OffsetForLeaderEpochResponse) GetVersion() int16 {
	return r.Version
}

func (r *OffsetForLeaderEpochResponse) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 1
	}
	return 0
}

func (r *OffsetForLeaderEpochResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *OffsetForLeaderEpochResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *OffsetForLeaderEpochResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
