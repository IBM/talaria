// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// ProducerState contains a
type ProducerState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ProducerID contains a
	ProducerID int64
	// ProducerEpoch contains a
	ProducerEpoch int32
	// LastSequence contains a
	LastSequence int32
	// LastTimestamp contains a
	LastTimestamp int64
	// CoordinatorEpoch contains a
	CoordinatorEpoch int32
	// CurrentTxnStartOffset contains a
	CurrentTxnStartOffset int64
}

func (a *ProducerState) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt64(a.ProducerID)

	pe.putInt32(a.ProducerEpoch)

	pe.putInt32(a.LastSequence)

	pe.putInt64(a.LastTimestamp)

	pe.putInt32(a.CoordinatorEpoch)

	pe.putInt64(a.CurrentTxnStartOffset)

	pe.putUVarint(0)
	return nil
}

func (a *ProducerState) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if a.ProducerEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if a.LastSequence, err = pd.getInt32(); err != nil {
		return err
	}

	if a.LastTimestamp, err = pd.getInt64(); err != nil {
		return err
	}

	if a.CoordinatorEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if a.CurrentTxnStartOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// PartitionResponse contains each partition in the response.
type PartitionResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the partition error message, which may be null if no additional details are available
	ErrorMessage *string
	// ActiveProducers contains a
	ActiveProducers []ProducerState
}

func (p *PartitionResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(p.ActiveProducers)); err != nil {
		return err
	}
	for _, block := range p.ActiveProducers {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numActiveProducers int
	if numActiveProducers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numActiveProducers > 0 {
		p.ActiveProducers = make([]ProducerState, numActiveProducers)
		for i := 0; i < numActiveProducers; i++ {
			var block ProducerState
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.ActiveProducers[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicResponse contains each topic in the response.
type TopicResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name string
	// Partitions contains each partition in the response.
	Partitions []PartitionResponse
}

func (t *TopicResponse) encode(pe packetEncoder, version int16) (err error) {
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

	pe.putUVarint(0)
	return nil
}

func (t *TopicResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeProducersResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains each topic in the response.
	Topics []TopicResponse
}

func (r *DescribeProducersResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

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

func (r *DescribeProducersResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeProducersResponse) GetKey() int16 {
	return 61
}

func (r *DescribeProducersResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeProducersResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeProducersResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeProducersResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeProducersResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
