// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// ReplicaState_FetchRequest contains a
type ReplicaState_FetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ReplicaID contains the replica ID of the follower, or -1 if this request is from a consumer.
	ReplicaID int32
	// ReplicaEpoch contains the epoch of this follower, or -1 if not available.
	ReplicaEpoch int64
}

func (r *ReplicaState_FetchRequest) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 15 {
		pe.putInt32(r.ReplicaID)
	}

	if r.Version >= 15 {
		pe.putInt64(r.ReplicaEpoch)
	}

	return nil
}

func (r *ReplicaState_FetchRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 15 {
		if r.ReplicaID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 15 {
		if r.ReplicaEpoch, err = pd.getInt64(); err != nil {
			return err
		}
	}

	return nil
}

// FetchPartition contains the partitions to fetch.
type FetchPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// CurrentLeaderEpoch contains the current leader epoch of the partition.
	CurrentLeaderEpoch int32
	// FetchOffset contains the message offset.
	FetchOffset int64
	// LastFetchedEpoch contains the epoch of the last fetched record or -1 if there is none
	LastFetchedEpoch int32
	// LogStartOffset contains the earliest available offset of the follower replica.  The field is only used when the request is sent by the follower.
	LogStartOffset int64
	// PartitionMaxBytes contains the maximum bytes to fetch from this partition.  See KIP-74 for cases where this limit may not be honored.
	PartitionMaxBytes int32
}

func (p *FetchPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	if p.Version >= 9 {
		pe.putInt32(p.CurrentLeaderEpoch)
	}

	pe.putInt64(p.FetchOffset)

	if p.Version >= 12 {
		pe.putInt32(p.LastFetchedEpoch)
	}

	if p.Version >= 5 {
		pe.putInt64(p.LogStartOffset)
	}

	pe.putInt32(p.PartitionMaxBytes)

	if p.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *FetchPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 9 {
		if p.CurrentLeaderEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.FetchOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 12 {
		if p.LastFetchedEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 5 {
		if p.LogStartOffset, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if p.PartitionMaxBytes, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// FetchTopic contains the topics to fetch.
type FetchTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the name of the topic to fetch.
	Topic string
	// TopicID contains the unique topic ID
	TopicID uuid.UUID
	// Partitions contains the partitions to fetch.
	Partitions []FetchPartition
}

func (t *FetchTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 12 {
		if err := pe.putString(t.Topic); err != nil {
			return err
		}
	}

	if t.Version >= 13 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(t.Partitions)); err != nil {
		return err
	}
	for _, block := range t.Partitions {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *FetchTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 12 {
		if t.Topic, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 13 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]FetchPartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block FetchPartition
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// ForgottenTopic contains in an incremental fetch request, the partitions to remove.
type ForgottenTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the topic name.
	Topic string
	// TopicID contains the unique topic ID
	TopicID uuid.UUID
	// Partitions contains the partitions indexes to forget.
	Partitions []int32
}

func (f *ForgottenTopic) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	if f.Version >= 7 && f.Version <= 12 {
		if err := pe.putString(f.Topic); err != nil {
			return err
		}
	}

	if f.Version >= 13 {
		if err := pe.putUUID(f.TopicID); err != nil {
			return err
		}
	}

	if f.Version >= 7 {
		if err := pe.putInt32Array(f.Partitions); err != nil {
			return err
		}
	}

	if f.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (f *ForgottenTopic) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.Version >= 7 && f.Version <= 12 {
		if f.Topic, err = pd.getString(); err != nil {
			return err
		}
	}

	if f.Version >= 13 {
		if f.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if f.Version >= 7 {
		if f.Partitions, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if f.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type FetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the clusterId if known. This is used to validate metadata fetches prior to broker registration.
	ClusterID *string
	// ReplicaID contains the broker ID of the follower, of -1 if this request is from a consumer.
	ReplicaID int32
	// ReplicaState contains a
	ReplicaState ReplicaState_FetchRequest
	// MaxWaitMs contains the maximum time in milliseconds to wait for the response.
	MaxWaitMs int32
	// MinBytes contains the minimum bytes to accumulate in the response.
	MinBytes int32
	// MaxBytes contains the maximum bytes to fetch.  See KIP-74 for cases where this limit may not be honored.
	MaxBytes int32
	// IsolationLevel contains a This setting controls the visibility of transactional records. Using READ_UNCOMMITTED (isolation_level = 0) makes all records visible. With READ_COMMITTED (isolation_level = 1), non-transactional and COMMITTED transactional records are visible. To be more concrete, READ_COMMITTED returns all data from offsets smaller than the current LSO (last stable offset), and enables the inclusion of the list of aborted transactions in the result, which allows consumers to discard ABORTED transactional records
	IsolationLevel int8
	// SessionID contains the fetch session ID.
	SessionID int32
	// SessionEpoch contains the fetch session epoch, which is used for ordering requests in a session.
	SessionEpoch int32
	// Topics contains the topics to fetch.
	Topics []FetchTopic
	// ForgottenTopicsData contains in an incremental fetch request, the partitions to remove.
	ForgottenTopicsData []ForgottenTopic
	// RackID contains a Rack ID of the consumer making this request
	RackID string
}

func (r *FetchRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 12 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 0 && r.Version <= 14 {
		pe.putInt32(r.ReplicaID)
	}

	pe.putInt32(r.MaxWaitMs)

	pe.putInt32(r.MinBytes)

	if r.Version >= 3 {
		pe.putInt32(r.MaxBytes)
	}

	if r.Version >= 4 {
		pe.putInt8(r.IsolationLevel)
	}

	if r.Version >= 7 {
		pe.putInt32(r.SessionID)
	}

	if r.Version >= 7 {
		pe.putInt32(r.SessionEpoch)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if err := pe.putArrayLength(len(r.ForgottenTopicsData)); err != nil {
			return err
		}
		for _, block := range r.ForgottenTopicsData {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 11 {
		if err := pe.putString(r.RackID); err != nil {
			return err
		}
	}

	if r.Version >= 12 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *FetchRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 12 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 0 && r.Version <= 14 {
		if r.ReplicaID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.MaxWaitMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.MinBytes, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.MaxBytes, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if r.IsolationLevel, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if r.SessionID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if r.SessionEpoch, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]FetchTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block FetchTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 7 {
		var numForgottenTopicsData int
		if numForgottenTopicsData, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numForgottenTopicsData > 0 {
			r.ForgottenTopicsData = make([]ForgottenTopic, numForgottenTopicsData)
			for i := 0; i < numForgottenTopicsData; i++ {
				var block ForgottenTopic
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.ForgottenTopicsData[i] = block
			}
		}
	}

	if r.Version >= 11 {
		if r.RackID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 12 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *FetchRequest) GetKey() int16 {
	return 1
}

func (r *FetchRequest) GetVersion() int16 {
	return r.Version
}

func (r *FetchRequest) GetHeaderVersion() int16 {
	if r.Version >= 12 {
		return 2
	}
	return 1
}

func (r *FetchRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 15
}

func (r *FetchRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
