// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// EpochEndOffset_FetchResponse contains in case divergence is detected based on the `LastFetchedEpoch` and `FetchOffset` in the request, this field indicates the largest epoch and its end offset such that subsequent records are known to diverge
type EpochEndOffset_FetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Epoch contains a
	Epoch int32
	// EndOffset contains a
	EndOffset int64
}

func (d *EpochEndOffset_FetchResponse) encode(pe packetEncoder, version int16) (err error) {
	d.Version = version
	if d.Version >= 12 {
		pe.putInt32(d.Epoch)
	}

	if d.Version >= 12 {
		pe.putInt64(d.EndOffset)
	}

	return nil
}

func (d *EpochEndOffset_FetchResponse) decode(pd packetDecoder, version int16) (err error) {
	d.Version = version
	if d.Version >= 12 {
		if d.Epoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if d.Version >= 12 {
		if d.EndOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	return nil
}

// LeaderIdAndEpoch_FetchResponse contains a
type LeaderIdAndEpoch_FetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch
	LeaderEpoch int32
}

func (c *LeaderIdAndEpoch_FetchResponse) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if c.Version >= 12 {
		pe.putInt32(c.LeaderID)
	}

	if c.Version >= 12 {
		pe.putInt32(c.LeaderEpoch)
	}

	return nil
}

func (c *LeaderIdAndEpoch_FetchResponse) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Version >= 12 {
		if c.LeaderID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if c.Version >= 12 {
		if c.LeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	return nil
}

// SnapshotId_FetchResponse contains in the case of fetching an offset less than the LogStartOffset, this is the end offset and epoch that should be used in the FetchSnapshot request.
type SnapshotId_FetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EndOffset contains a
	EndOffset int64
	// Epoch contains a
	Epoch int32
}

func (s *SnapshotId_FetchResponse) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	pe.putInt64(s.EndOffset)

	pe.putInt32(s.Epoch)

	return nil
}

func (s *SnapshotId_FetchResponse) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.EndOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if s.Epoch, err = pd.getInt32(); err != nil {
		return err
	}

	return nil
}

// AbortedTransaction contains the aborted transactions.
type AbortedTransaction struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ProducerID contains the producer id associated with the aborted transaction.
	ProducerID int64
	// FirstOffset contains the first offset in the aborted transaction.
	FirstOffset int64
}

func (a *AbortedTransaction) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if a.Version >= 4 {
		pe.putInt64(a.ProducerID)
	}

	if a.Version >= 4 {
		pe.putInt64(a.FirstOffset)
	}

	if a.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *AbortedTransaction) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.Version >= 4 {
		if a.ProducerID, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if a.Version >= 4 {
		if a.FirstOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if a.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// PartitionData_FetchResponse contains the topic partitions.
type PartitionData_FetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code, or 0 if there was no fetch error.
	ErrorCode int16
	// HighWatermark contains the current high water mark.
	HighWatermark int64
	// LastStableOffset contains the last stable offset (or LSO) of the partition. This is the last offset such that the state of all transactional records prior to this offset have been decided (ABORTED or COMMITTED)
	LastStableOffset int64
	// LogStartOffset contains the current log start offset.
	LogStartOffset int64
	// DivergingEpoch contains in case divergence is detected based on the `LastFetchedEpoch` and `FetchOffset` in the request, this field indicates the largest epoch and its end offset such that subsequent records are known to diverge
	DivergingEpoch EpochEndOffset_FetchResponse
	// CurrentLeader contains a
	CurrentLeader LeaderIdAndEpoch_FetchResponse
	// SnapshotID contains in the case of fetching an offset less than the LogStartOffset, this is the end offset and epoch that should be used in the FetchSnapshot request.
	SnapshotID SnapshotId_FetchResponse
	// AbortedTransactions contains the aborted transactions.
	AbortedTransactions []AbortedTransaction
	// PreferredReadReplica contains the preferred read replica for the consumer to use on its next fetch request
	PreferredReadReplica int32
	// Records contains the record data.
	Records RecordBatch
}

func (p *PartitionData_FetchResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	pe.putInt64(p.HighWatermark)

	if p.Version >= 4 {
		pe.putInt64(p.LastStableOffset)
	}

	if p.Version >= 5 {
		pe.putInt64(p.LogStartOffset)
	}

	if p.Version >= 4 {
		if err := pe.putArrayLength(len(p.AbortedTransactions)); err != nil {
			return err
		}
		for _, block := range p.AbortedTransactions {
			if err := block.encode(pe, p.Version); err != nil {
				return err
			}
		}
	}

	if p.Version >= 11 {
		pe.putInt32(p.PreferredReadReplica)
	}

	if err := p.Records.encode(pe, p.Version); err != nil {
		return err
	}

	if p.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionData_FetchResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.HighWatermark, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 4 {
		if p.LastStableOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 5 {
		if p.LogStartOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.Version >= 4 {
		var numAbortedTransactions int
		if numAbortedTransactions, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numAbortedTransactions > 0 {
			p.AbortedTransactions = make([]AbortedTransaction, numAbortedTransactions)
			for i := 0; i < numAbortedTransactions; i++ {
				var block AbortedTransaction
				if err := block.decode(pd, p.Version); err != nil {
					return err
				}
				p.AbortedTransactions[i] = block
			}
		}
	}

	if p.Version >= 11 {
		if p.PreferredReadReplica, err = pd.getInt32(); err != nil {
			return err
		}
	}

	tmprecords := RecordBatch{}
	if err := tmprecords.decode(pd, p.Version); err != nil {
		return err
	}
	p.Records = tmprecords

	if p.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// FetchableTopicResponse contains the response topics.
type FetchableTopicResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the topic name.
	Topic string
	// TopicID contains the unique topic ID
	TopicID uuid.UUID
	// Partitions contains the topic partitions.
	Partitions []PartitionData_FetchResponse
}

func (r *FetchableTopicResponse) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 0 && r.Version <= 12 {
		if err := pe.putString(r.Topic); err != nil {
			return err
		}
	}

	if r.Version >= 13 {
		if err := pe.putUUID(r.TopicID); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Partitions)); err != nil {
		return err
	}
	for _, block := range r.Partitions {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *FetchableTopicResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 0 && r.Version <= 12 {
		if r.Topic, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 13 {
		if r.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		r.Partitions = make([]PartitionData_FetchResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_FetchResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Partitions[i] = block
		}
	}

	if r.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type FetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code.
	ErrorCode int16
	// SessionID contains the fetch session ID, or 0 if this is not part of a fetch session.
	SessionID int32
	// Responses contains the response topics.
	Responses []FetchableTopicResponse
}

func (r *FetchResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 12 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if r.Version >= 7 {
		pe.putInt16(r.ErrorCode)
	}

	if r.Version >= 7 {
		pe.putInt32(r.SessionID)
	}

	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *FetchResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 12 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if r.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if r.SessionID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResponses > 0 {
		r.Responses = make([]FetchableTopicResponse, numResponses)
		for i := 0; i < numResponses; i++ {
			var block FetchableTopicResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Responses[i] = block
		}
	}

	if r.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *FetchResponse) GetKey() int16 {
	return 1
}

func (r *FetchResponse) GetVersion() int16 {
	return r.Version
}

func (r *FetchResponse) GetHeaderVersion() int16 {
	if r.Version >= 12 {
		return 1
	}
	return 0
}

func (r *FetchResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 15
}

func (r *FetchResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *FetchResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
