// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// ReassignablePartitionResponse contains the responses to partitions to reassign
type ReassignablePartitionResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code for this partition, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message for this partition, or null if there was no error.
	ErrorMessage *string
}

func (p *ReassignablePartitionResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *ReassignablePartitionResponse) decode(pd packetDecoder, version int16) (err error) {
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

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// ReassignableTopicResponse contains the responses to topics to reassign.
type ReassignableTopicResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name string
	// Partitions contains the responses to partitions to reassign
	Partitions []ReassignablePartitionResponse
}

func (r *ReassignableTopicResponse) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.Name); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Partitions)); err != nil {
		return err
	}
	for _, block := range r.Partitions {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ReassignableTopicResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Partitions = make([]ReassignablePartitionResponse, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block ReassignablePartitionResponse
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Partitions[i] = block
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type AlterPartitionReassignmentsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// Responses contains the responses to topics to reassign.
	Responses []ReassignableTopicResponse
}

func (r *AlterPartitionReassignmentsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *AlterPartitionReassignmentsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Responses = make([]ReassignableTopicResponse, numResponses)
	for i := 0; i < numResponses; i++ {
		var block ReassignableTopicResponse
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Responses[i] = block
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *AlterPartitionReassignmentsResponse) GetKey() int16 {
	return 45
}

func (r *AlterPartitionReassignmentsResponse) GetVersion() int16 {
	return r.Version
}

func (r *AlterPartitionReassignmentsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *AlterPartitionReassignmentsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AlterPartitionReassignmentsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AlterPartitionReassignmentsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
