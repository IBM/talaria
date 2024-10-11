// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// PartitionResult contains the results for each partition
type PartitionResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionID contains the partition id
	PartitionID int32
	// ErrorCode contains the result error, or zero if there was no error.
	ErrorCode int16
	// ErrorMessage contains the result message, or null if there was no error.
	ErrorMessage *string
}

func (p *PartitionResult) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionID)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionResult) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// ReplicaElectionResult contains the election results, or an empty array if the requester did not have permission and the request asks for all partitions.
type ReplicaElectionResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the topic name
	Topic string
	// PartitionResult contains the results for each partition
	PartitionResult []PartitionResult
}

func (r *ReplicaElectionResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.Topic); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.PartitionResult)); err != nil {
		return err
	}
	for _, block := range r.PartitionResult {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ReplicaElectionResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Topic, err = pd.getString(); err != nil {
		return err
	}

	var numPartitionResult int
	if numPartitionResult, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.PartitionResult = make([]PartitionResult, numPartitionResult)
	for i := 0; i < numPartitionResult; i++ {
		var block PartitionResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.PartitionResult[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ElectLeadersResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code.
	ErrorCode int16
	// ReplicaElectionResults contains the election results, or an empty array if the requester did not have permission and the request asks for all partitions.
	ReplicaElectionResults []ReplicaElectionResult
}

func (r *ElectLeadersResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if r.Version >= 1 {
		pe.putInt16(r.ErrorCode)
	}

	if err := pe.putArrayLength(len(r.ReplicaElectionResults)); err != nil {
		return err
	}
	for _, block := range r.ReplicaElectionResults {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ElectLeadersResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	var numReplicaElectionResults int
	if numReplicaElectionResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.ReplicaElectionResults = make([]ReplicaElectionResult, numReplicaElectionResults)
	for i := 0; i < numReplicaElectionResults; i++ {
		var block ReplicaElectionResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.ReplicaElectionResults[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ElectLeadersResponse) GetKey() int16 {
	return 43
}

func (r *ElectLeadersResponse) GetVersion() int16 {
	return r.Version
}

func (r *ElectLeadersResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *ElectLeadersResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *ElectLeadersResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ElectLeadersResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
