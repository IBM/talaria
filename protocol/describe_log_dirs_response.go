// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DescribeLogDirsPartition contains a
type DescribeLogDirsPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// PartitionSize contains the size of the log segments in this partition in bytes.
	PartitionSize int64
	// OffsetLag contains the lag of the log's LEO w.r.t. partition's HW (if it is the current log for the partition) or current replica's LEO (if it is the future log for the partition)
	OffsetLag int64
	// IsFutureKey contains a True if this log is created by AlterReplicaLogDirsRequest and will replace the current log of the replica in the future.
	IsFutureKey bool
}

func (p *DescribeLogDirsPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt64(p.PartitionSize)

	pe.putInt64(p.OffsetLag)

	pe.putBool(p.IsFutureKey)

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *DescribeLogDirsPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.PartitionSize, err = pd.getInt64(); err != nil {
		return err
	}

	if p.OffsetLag, err = pd.getInt64(); err != nil {
		return err
	}

	if p.IsFutureKey, err = pd.getBool(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribeLogDirsTopic contains each topic.
type DescribeLogDirsTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains a
	Partitions []DescribeLogDirsPartition
}

func (t *DescribeLogDirsTopic) encode(pe packetEncoder, version int16) (err error) {
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

func (t *DescribeLogDirsTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]DescribeLogDirsPartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block DescribeLogDirsPartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribeLogDirsResult contains the log directories.
type DescribeLogDirsResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// LogDir contains the absolute log directory path.
	LogDir string
	// Topics contains each topic.
	Topics []DescribeLogDirsTopic
	// TotalBytes contains the total size in bytes of the volume the log directory is in.
	TotalBytes int64
	// UsableBytes contains the usable size in bytes of the volume the log directory is in.
	UsableBytes int64
}

func (r *DescribeLogDirsResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt16(r.ErrorCode)

	if err := pe.putString(r.LogDir); err != nil {
		return err
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
		pe.putInt64(r.TotalBytes)
	}

	if r.Version >= 4 {
		pe.putInt64(r.UsableBytes)
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeLogDirsResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.LogDir, err = pd.getString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]DescribeLogDirsTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block DescribeLogDirsTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.Version >= 4 {
		if r.TotalBytes, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if r.UsableBytes, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogDirsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// Results contains the log directories.
	Results []DescribeLogDirsResult
}

func (r *DescribeLogDirsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if r.Version >= 3 {
		pe.putInt16(r.ErrorCode)
	}

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

func (r *DescribeLogDirsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Results = make([]DescribeLogDirsResult, numResults)
	for i := 0; i < numResults; i++ {
		var block DescribeLogDirsResult
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

func (r *DescribeLogDirsResponse) GetKey() int16 {
	return 35
}

func (r *DescribeLogDirsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeLogDirsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *DescribeLogDirsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *DescribeLogDirsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeLogDirsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
