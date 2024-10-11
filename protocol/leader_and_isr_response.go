// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// LeaderAndIsrPartitionError_LeaderAndIsrResponse contains a
type LeaderAndIsrPartitionError_LeaderAndIsrResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition error code, or 0 if there was no error.
	ErrorCode int16
}

func (l *LeaderAndIsrPartitionError_LeaderAndIsrResponse) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 0 && l.Version <= 4 {
		if err := pe.putString(l.TopicName); err != nil {
			return err
		}
	}

	pe.putInt32(l.PartitionIndex)

	pe.putInt16(l.ErrorCode)

	if l.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (l *LeaderAndIsrPartitionError_LeaderAndIsrResponse) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 0 && l.Version <= 4 {
		if l.TopicName, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if l.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if l.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// LeaderAndIsrTopicError contains each topic
type LeaderAndIsrTopicError struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the unique topic ID
	TopicID uuid.UUID
	// PartitionErrors contains each partition.
	PartitionErrors []LeaderAndIsrPartitionError_LeaderAndIsrResponse
}

func (t *LeaderAndIsrTopicError) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 5 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		if err := pe.putArrayLength(len(t.PartitionErrors)); err != nil {
			return err
		}
		for _, block := range t.PartitionErrors {
			if err := block.encode(pe, t.Version); err != nil {
				return err
			}
		}
	}

	if t.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *LeaderAndIsrTopicError) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 5 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		var numPartitionErrors int
		if numPartitionErrors, err = pd.getArrayLength(); err != nil {
			return err
		}
		t.PartitionErrors = make([]LeaderAndIsrPartitionError_LeaderAndIsrResponse, numPartitionErrors)
		for i := 0; i < numPartitionErrors; i++ {
			var block LeaderAndIsrPartitionError_LeaderAndIsrResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.PartitionErrors[i] = block
		}
	}

	if t.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type LeaderAndIsrResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// PartitionErrors contains each partition in v0 to v4 message.
	PartitionErrors []LeaderAndIsrPartitionError_LeaderAndIsrResponse
	// Topics contains each topic
	Topics []LeaderAndIsrTopicError
}

func (r *LeaderAndIsrResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if r.Version >= 0 && r.Version <= 4 {
		if err := pe.putArrayLength(len(r.PartitionErrors)); err != nil {
			return err
		}
		for _, block := range r.PartitionErrors {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 5 {
		if err := pe.putArrayLength(len(r.Topics)); err != nil {
			return err
		}
		for _, block := range r.Topics {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *LeaderAndIsrResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 0 && r.Version <= 4 {
		var numPartitionErrors int
		if numPartitionErrors, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.PartitionErrors = make([]LeaderAndIsrPartitionError_LeaderAndIsrResponse, numPartitionErrors)
		for i := 0; i < numPartitionErrors; i++ {
			var block LeaderAndIsrPartitionError_LeaderAndIsrResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.PartitionErrors[i] = block
		}
	}

	if r.Version >= 5 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Topics = make([]LeaderAndIsrTopicError, numTopics)
		for i := 0; i < numTopics; i++ {
			var block LeaderAndIsrTopicError
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *LeaderAndIsrResponse) GetKey() int16 {
	return 4
}

func (r *LeaderAndIsrResponse) GetVersion() int16 {
	return r.Version
}

func (r *LeaderAndIsrResponse) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 1
	}
	return 0
}

func (r *LeaderAndIsrResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 7
}

func (r *LeaderAndIsrResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
