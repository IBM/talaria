// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// CreatableTopicConfigs contains a Configuration of the topic.
type CreatableTopicConfigs struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the configuration name.
	Name string
	// Value contains the configuration value.
	Value *string
	// ReadOnly contains a True if the configuration is read-only.
	ReadOnly bool
	// ConfigSource contains the configuration source.
	ConfigSource int8
	// IsSensitive contains a True if this configuration is sensitive.
	IsSensitive bool
}

func (c *CreatableTopicConfigs) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if c.Version >= 5 {
		if err := pe.putString(c.Name); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		if err := pe.putNullableString(c.Value); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		pe.putBool(c.ReadOnly)
	}

	if c.Version >= 5 {
		pe.putInt8(c.ConfigSource)
	}

	if c.Version >= 5 {
		pe.putBool(c.IsSensitive)
	}

	if c.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *CreatableTopicConfigs) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Version >= 5 {
		if c.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		if c.Value, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		if c.ReadOnly, err = pd.getBool(); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		if c.ConfigSource, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		if c.IsSensitive, err = pd.getBool(); err != nil {
			return err
		}
	}

	if c.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// CreatableTopicResult contains a Results for each topic we tried to create.
type CreatableTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// TopicID contains the unique topic ID
	TopicID uuid.UUID
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// TopicConfigErrorCode contains a Optional topic config error returned if configs are not returned in the response.
	TopicConfigErrorCode int16
	// NumPartitions contains a Number of partitions of the topic.
	NumPartitions int32
	// ReplicationFactor contains a Replication factor of the topic.
	ReplicationFactor int16
	// Configs contains a Configuration of the topic.
	Configs []CreatableTopicConfigs
}

func (t *CreatableTopicResult) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if t.Version >= 7 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	pe.putInt16(t.ErrorCode)

	if t.Version >= 1 {
		if err := pe.putNullableString(t.ErrorMessage); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		pe.putInt32(t.NumPartitions)
	}

	if t.Version >= 5 {
		pe.putInt16(t.ReplicationFactor)
	}

	if t.Version >= 5 {
		if err := pe.putArrayLength(len(t.Configs)); err != nil {
			return err
		}
		for _, block := range t.Configs {
			if err := block.encode(pe, t.Version); err != nil {
				return err
			}
		}
	}

	if t.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *CreatableTopicResult) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.Version >= 7 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if t.Version >= 1 {
		if t.ErrorMessage, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		if t.NumPartitions, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		if t.ReplicationFactor, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if t.Version >= 5 {
		var numConfigs int
		if numConfigs, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numConfigs > 0 {
			t.Configs = make([]CreatableTopicConfigs, numConfigs)
			for i := 0; i < numConfigs; i++ {
				var block CreatableTopicConfigs
				if err := block.decode(pd, t.Version); err != nil {
					return err
				}
				t.Configs[i] = block
			}
		}
	}

	if t.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTopicsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains a Results for each topic we tried to create.
	Topics []CreatableTopicResult
}

func (r *CreateTopicsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 5 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 2 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreateTopicsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 5 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 2 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]CreatableTopicResult, numTopics)
		for i := 0; i < numTopics; i++ {
			var block CreatableTopicResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateTopicsResponse) GetKey() int16 {
	return 19
}

func (r *CreateTopicsResponse) GetVersion() int16 {
	return r.Version
}

func (r *CreateTopicsResponse) GetHeaderVersion() int16 {
	if r.Version >= 5 {
		return 1
	}
	return 0
}

func (r *CreateTopicsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 7
}

func (r *CreateTopicsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *CreateTopicsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
