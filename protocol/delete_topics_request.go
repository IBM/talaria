// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// DeleteTopicState contains the name or topic ID of the topic
type DeleteTopicState struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name *string
	// TopicID contains the unique topic ID
	TopicID uuid.UUID
}

func (t *DeleteTopicState) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 6 {
		if err := pe.putNullableString(t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 6 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *DeleteTopicState) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 6 {
		if t.Name, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if t.Version >= 6 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteTopicsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains the name or topic ID of the topic
	Topics []DeleteTopicState
	// TopicNames contains the names of the topics to delete
	TopicNames []string
	// TimeoutMs contains the length of time in milliseconds to wait for the deletions to complete.
	TimeoutMs int32
}

func (r *DeleteTopicsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 6 {
		if err := pe.putArrayLength(len(r.Topics)); err != nil {
			return err
		}
		for _, block := range r.Topics {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 0 && r.Version <= 5 {
		if err := pe.putStringArray(r.TopicNames); err != nil {
			return err
		}
	}

	pe.putInt32(r.TimeoutMs)

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteTopicsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 6 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopics > 0 {
			r.Topics = make([]DeleteTopicState, numTopics)
			for i := 0; i < numTopics; i++ {
				var block DeleteTopicState
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.Topics[i] = block
			}
		}
	}

	if r.Version >= 0 && r.Version <= 5 {
		if r.TopicNames, err = pd.getStringArray(); err != nil {
			return err
		}
	}

	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteTopicsRequest) GetKey() int16 {
	return 20
}

func (r *DeleteTopicsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DeleteTopicsRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *DeleteTopicsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 6
}

func (r *DeleteTopicsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
