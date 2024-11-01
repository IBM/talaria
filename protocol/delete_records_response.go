// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DeleteRecordsPartitionResult contains each partition that we wanted to delete records from.
type DeleteRecordsPartitionResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// LowWatermark contains the partition low water mark.
	LowWatermark int64
	// ErrorCode contains the deletion error code, or 0 if the deletion succeeded.
	ErrorCode int16
}

func (p *DeleteRecordsPartitionResult) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt64(p.LowWatermark)

	pe.putInt16(p.ErrorCode)

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *DeleteRecordsPartitionResult) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LowWatermark, err = pd.getInt64(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRecordsTopicResult contains each topic that we wanted to delete records from.
type DeleteRecordsTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains each partition that we wanted to delete records from.
	Partitions []DeleteRecordsPartitionResult
}

func (t *DeleteRecordsTopicResult) encode(pe packetEncoder, version int16) (err error) {
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

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *DeleteRecordsTopicResult) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]DeleteRecordsPartitionResult, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block DeleteRecordsPartitionResult
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRecordsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains each topic that we wanted to delete records from.
	Topics []DeleteRecordsTopicResult
}

func (r *DeleteRecordsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteRecordsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]DeleteRecordsTopicResult, numTopics)
		for i := 0; i < numTopics; i++ {
			var block DeleteRecordsTopicResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteRecordsResponse) GetKey() int16 {
	return 21
}

func (r *DeleteRecordsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DeleteRecordsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *DeleteRecordsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *DeleteRecordsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DeleteRecordsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
