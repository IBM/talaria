// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// OffsetFetchResponsePartition contains the responses per partition
type OffsetFetchResponsePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// CommittedOffset contains the committed message offset.
	CommittedOffset int64
	// CommittedLeaderEpoch contains the leader epoch.
	CommittedLeaderEpoch int32
	// Metadata contains the partition metadata.
	Metadata *string
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (p *OffsetFetchResponsePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 0 && p.Version <= 7 {
		pe.putInt32(p.PartitionIndex)
	}

	if p.Version >= 0 && p.Version <= 7 {
		pe.putInt64(p.CommittedOffset)
	}

	if p.Version >= 5 && p.Version <= 7 {
		pe.putInt32(p.CommittedLeaderEpoch)
	}

	if p.Version >= 0 && p.Version <= 7 {
		if err := pe.putNullableString(p.Metadata); err != nil {
			return err
		}
	}

	if p.Version >= 0 && p.Version <= 7 {
		pe.putInt16(p.ErrorCode)
	}

	if p.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *OffsetFetchResponsePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 0 && p.Version <= 7 {
		if p.PartitionIndex, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 0 && p.Version <= 7 {
		if p.CommittedOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 5 && p.Version <= 7 {
		if p.CommittedLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 0 && p.Version <= 7 {
		if p.Metadata, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if p.Version >= 0 && p.Version <= 7 {
		if p.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if p.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetFetchResponseTopic contains the responses per topic.
type OffsetFetchResponseTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the responses per partition
	Partitions []OffsetFetchResponsePartition
}

func (t *OffsetFetchResponseTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 7 {
		if err := pe.putString(t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 0 && t.Version <= 7 {
		if err := pe.putArrayLength(len(t.Partitions)); err != nil {
			return err
		}
		for _, block := range t.Partitions {
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

func (t *OffsetFetchResponseTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 7 {
		if t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 0 && t.Version <= 7 {
		var numPartitions int
		if numPartitions, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numPartitions > 0 {
			t.Partitions = make([]OffsetFetchResponsePartition, numPartitions)
			for i := 0; i < numPartitions; i++ {
				var block OffsetFetchResponsePartition
				if err := block.decode(pd, t.Version); err != nil {
					return err
				}
				t.Partitions[i] = block
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

// OffsetFetchResponsePartitions contains the responses per partition
type OffsetFetchResponsePartitions struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// CommittedOffset contains the committed message offset.
	CommittedOffset int64
	// CommittedLeaderEpoch contains the leader epoch.
	CommittedLeaderEpoch int32
	// Metadata contains the partition metadata.
	Metadata *string
	// ErrorCode contains the partition-level error code, or 0 if there was no error.
	ErrorCode int16
}

func (p *OffsetFetchResponsePartitions) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 8 {
		pe.putInt32(p.PartitionIndex)
	}

	if p.Version >= 8 {
		pe.putInt64(p.CommittedOffset)
	}

	if p.Version >= 8 {
		pe.putInt32(p.CommittedLeaderEpoch)
	}

	if p.Version >= 8 {
		if err := pe.putNullableString(p.Metadata); err != nil {
			return err
		}
	}

	if p.Version >= 8 {
		pe.putInt16(p.ErrorCode)
	}

	if p.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *OffsetFetchResponsePartitions) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 8 {
		if p.PartitionIndex, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 8 {
		if p.CommittedOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 8 {
		if p.CommittedLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 8 {
		if p.Metadata, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if p.Version >= 8 {
		if p.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if p.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetFetchResponseTopics contains the responses per topic.
type OffsetFetchResponseTopics struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the responses per partition
	Partitions []OffsetFetchResponsePartitions
}

func (t *OffsetFetchResponseTopics) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 8 {
		if err := pe.putString(t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 8 {
		if err := pe.putArrayLength(len(t.Partitions)); err != nil {
			return err
		}
		for _, block := range t.Partitions {
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

func (t *OffsetFetchResponseTopics) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 8 {
		if t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 8 {
		var numPartitions int
		if numPartitions, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numPartitions > 0 {
			t.Partitions = make([]OffsetFetchResponsePartitions, numPartitions)
			for i := 0; i < numPartitions; i++ {
				var block OffsetFetchResponsePartitions
				if err := block.decode(pd, t.Version); err != nil {
					return err
				}
				t.Partitions[i] = block
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

// OffsetFetchResponseGroup contains the responses per group id.
type OffsetFetchResponseGroup struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// groupID contains the group ID.
	groupID string
	// Topics contains the responses per topic.
	Topics []OffsetFetchResponseTopics
	// ErrorCode contains the group-level error code, or 0 if there was no error.
	ErrorCode int16
}

func (g *OffsetFetchResponseGroup) encode(pe packetEncoder, version int16) (err error) {
	g.Version = version
	if g.Version >= 8 {
		if err := pe.putString(g.groupID); err != nil {
			return err
		}
	}

	if g.Version >= 8 {
		if err := pe.putArrayLength(len(g.Topics)); err != nil {
			return err
		}
		for _, block := range g.Topics {
			if err := block.encode(pe, g.Version); err != nil {
				return err
			}
		}
	}

	if g.Version >= 8 {
		pe.putInt16(g.ErrorCode)
	}

	if g.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (g *OffsetFetchResponseGroup) decode(pd packetDecoder, version int16) (err error) {
	g.Version = version
	if g.Version >= 8 {
		if g.groupID, err = pd.getString(); err != nil {
			return err
		}
	}

	if g.Version >= 8 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopics > 0 {
			g.Topics = make([]OffsetFetchResponseTopics, numTopics)
			for i := 0; i < numTopics; i++ {
				var block OffsetFetchResponseTopics
				if err := block.decode(pd, g.Version); err != nil {
					return err
				}
				g.Topics[i] = block
			}
		}
	}

	if g.Version >= 8 {
		if g.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if g.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type OffsetFetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains the responses per topic.
	Topics []OffsetFetchResponseTopic
	// ErrorCode contains the top-level error code, or 0 if there was no error.
	ErrorCode int16
	// Groups contains the responses per group id.
	Groups []OffsetFetchResponseGroup
}

func (r *OffsetFetchResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if r.Version >= 0 && r.Version <= 7 {
		if err := pe.putArrayLength(len(r.Topics)); err != nil {
			return err
		}
		for _, block := range r.Topics {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 2 && r.Version <= 7 {
		pe.putInt16(r.ErrorCode)
	}

	if r.Version >= 8 {
		if err := pe.putArrayLength(len(r.Groups)); err != nil {
			return err
		}
		for _, block := range r.Groups {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *OffsetFetchResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 7 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopics > 0 {
			r.Topics = make([]OffsetFetchResponseTopic, numTopics)
			for i := 0; i < numTopics; i++ {
				var block OffsetFetchResponseTopic
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.Topics[i] = block
			}
		}
	}

	if r.Version >= 2 && r.Version <= 7 {
		if r.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if r.Version >= 8 {
		var numGroups int
		if numGroups, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numGroups > 0 {
			r.Groups = make([]OffsetFetchResponseGroup, numGroups)
			for i := 0; i < numGroups; i++ {
				var block OffsetFetchResponseGroup
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.Groups[i] = block
			}
		}
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *OffsetFetchResponse) GetKey() int16 {
	return 9
}

func (r *OffsetFetchResponse) GetVersion() int16 {
	return r.Version
}

func (r *OffsetFetchResponse) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 1
	}
	return 0
}

func (r *OffsetFetchResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *OffsetFetchResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *OffsetFetchResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
