// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// ListOffsetsPartitionResponse contains each partition in the response.
type ListOffsetsPartitionResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition error code, or 0 if there was no error.
	ErrorCode int16
	// OldStyleOffsets contains the result offsets.
	OldStyleOffsets []int64
	// Timestamp contains the timestamp associated with the returned offset.
	Timestamp int64
	// Offset contains the returned offset.
	Offset int64
	// LeaderEpoch contains a
	LeaderEpoch int32
}

func (p *ListOffsetsPartitionResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if p.Version == 0 {
		if err := pe.putInt64Array(p.OldStyleOffsets); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		pe.putInt64(p.Timestamp)
	}

	if p.Version >= 1 {
		pe.putInt64(p.Offset)
	}

	if p.Version >= 4 {
		pe.putInt32(p.LeaderEpoch)
	}

	if p.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *ListOffsetsPartitionResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Version == 0 {
		if p.OldStyleOffsets, err = pd.getInt64Array(); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if p.Timestamp, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if p.Offset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 4 {
		if p.LeaderEpoch, err = pd.getInt32(); err != nil {
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

// ListOffsetsTopicResponse contains each topic in the response.
type ListOffsetsTopicResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name string
	// Partitions contains each partition in the response.
	Partitions []ListOffsetsPartitionResponse
}

func (t *ListOffsetsTopicResponse) encode(pe packetEncoder, version int16) (err error) {
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

	if t.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *ListOffsetsTopicResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]ListOffsetsPartitionResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block ListOffsetsPartitionResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ListOffsetsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains each topic in the response.
	Topics []ListOffsetsTopicResponse
}

func (r *ListOffsetsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
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

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ListOffsetsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
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
		r.Topics = make([]ListOffsetsTopicResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block ListOffsetsTopicResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ListOffsetsResponse) GetKey() int16 {
	return 2
}

func (r *ListOffsetsResponse) GetVersion() int16 {
	return r.Version
}

func (r *ListOffsetsResponse) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 1
	}
	return 0
}

func (r *ListOffsetsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *ListOffsetsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ListOffsetsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
