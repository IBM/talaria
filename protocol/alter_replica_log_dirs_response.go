// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// AlterReplicaLogDirPartitionResult contains the results for each partition.
type AlterReplicaLogDirPartitionResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (p *AlterReplicaLogDirPartitionResult) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *AlterReplicaLogDirPartitionResult) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
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

// AlterReplicaLogDirTopicResult contains the results for each topic.
type AlterReplicaLogDirTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the name of the topic.
	TopicName string
	// Partitions contains the results for each partition.
	Partitions []AlterReplicaLogDirPartitionResult
}

func (r *AlterReplicaLogDirTopicResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.TopicName); err != nil {
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

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterReplicaLogDirTopicResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Partitions = make([]AlterReplicaLogDirPartitionResult, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block AlterReplicaLogDirPartitionResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Partitions[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AlterReplicaLogDirsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Results contains the results for each topic.
	Results []AlterReplicaLogDirTopicResult
}

func (r *AlterReplicaLogDirsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterReplicaLogDirsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Results = make([]AlterReplicaLogDirTopicResult, numResults)
	for i := 0; i < numResults; i++ {
		var block AlterReplicaLogDirTopicResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Results[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AlterReplicaLogDirsResponse) GetKey() int16 {
	return 34
}

func (r *AlterReplicaLogDirsResponse) GetVersion() int16 {
	return r.Version
}

func (r *AlterReplicaLogDirsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *AlterReplicaLogDirsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *AlterReplicaLogDirsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AlterReplicaLogDirsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
