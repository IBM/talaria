// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// BatchIndexAndErrorMessage contains the batch indices of records that caused the batch to be dropped
type BatchIndexAndErrorMessage struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BatchIndex contains the batch index of the record that cause the batch to be dropped
	BatchIndex int32
	// BatchIndexErrorMessage contains the error message of the record that caused the batch to be dropped
	BatchIndexErrorMessage *string
}

func (r *BatchIndexAndErrorMessage) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 8 {
		pe.putInt32(r.BatchIndex)
	}

	if r.Version >= 8 {
		if err := pe.putNullableString(r.BatchIndexErrorMessage); err != nil {
			return err
		}
	}

	if r.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *BatchIndexAndErrorMessage) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 8 {
		if r.BatchIndex, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 8 {
		if r.BatchIndexErrorMessage, err = pd.getNullableString(); err != nil {
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

// PartitionProduceResponse contains each partition that we produced to within the topic.
type PartitionProduceResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Index contains the partition index.
	Index int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// BaseOffset contains the base offset.
	BaseOffset int64
	// LogAppendTimeMs contains the timestamp returned by broker after appending the messages. If CreateTime is used for the topic, the timestamp will be -1.  If LogAppendTime is used for the topic, the timestamp will be the broker local time when the messages are appended.
	LogAppendTimeMs int64
	// LogStartOffset contains the log start offset.
	LogStartOffset int64
	// RecordErrors contains the batch indices of records that caused the batch to be dropped
	RecordErrors []BatchIndexAndErrorMessage
	// ErrorMessage contains the global error message summarizing the common root cause of the records that caused the batch to be dropped
	ErrorMessage *string
}

func (p *PartitionProduceResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Index)

	pe.putInt16(p.ErrorCode)

	pe.putInt64(p.BaseOffset)

	if p.Version >= 2 {
		pe.putInt64(p.LogAppendTimeMs)
	}

	if p.Version >= 5 {
		pe.putInt64(p.LogStartOffset)
	}

	if p.Version >= 8 {
		if err := pe.putArrayLength(len(p.RecordErrors)); err != nil {
			return err
		}
		for _, block := range p.RecordErrors {
			if err := block.encode(pe, p.Version); err != nil {
				return err
			}
		}
	}

	if p.Version >= 8 {
		if err := pe.putNullableString(p.ErrorMessage); err != nil {
			return err
		}
	}

	if p.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionProduceResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Index, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.BaseOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if p.LogAppendTimeMs, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 5 {
		if p.LogStartOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 8 {
		var numRecordErrors int
		if numRecordErrors, err = pd.getArrayLength(); err != nil {
			return err
		}
		p.RecordErrors = make([]BatchIndexAndErrorMessage, numRecordErrors)
		for i := 0; i < numRecordErrors; i++ {
			var block BatchIndexAndErrorMessage
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.RecordErrors[i] = block
		}
	}

	if p.Version >= 8 {
		if p.ErrorMessage, err = pd.getNullableString(); err != nil {
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

// TopicProduceResponse contains each produce response
type TopicProduceResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name string
	// PartitionResponses contains each partition that we produced to within the topic.
	PartitionResponses []PartitionProduceResponse
}

func (r *TopicProduceResponse) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.Name); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.PartitionResponses)); err != nil {
		return err
	}
	for _, block := range r.PartitionResponses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *TopicProduceResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitionResponses int
	if numPartitionResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.PartitionResponses = make([]PartitionProduceResponse, numPartitionResponses)
	for i := 0; i < numPartitionResponses; i++ {
		var block PartitionProduceResponse
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.PartitionResponses[i] = block
	}

	if r.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ProduceResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Responses contains each produce response
	Responses []TopicProduceResponse
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
}

func (r *ProduceResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 9 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if r.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ProduceResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 9 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Responses = make([]TopicProduceResponse, numResponses)
	for i := 0; i < numResponses; i++ {
		var block TopicProduceResponse
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Responses[i] = block
	}

	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
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

func (r *ProduceResponse) GetKey() int16 {
	return 0
}

func (r *ProduceResponse) GetVersion() int16 {
	return r.Version
}

func (r *ProduceResponse) GetHeaderVersion() int16 {
	if r.Version >= 9 {
		return 1
	}
	return 0
}

func (r *ProduceResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 9
}

func (r *ProduceResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ProduceResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
