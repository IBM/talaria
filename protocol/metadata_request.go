// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	"fmt"
	uuid "github.com/google/uuid"
)

// MetadataRequestTopic contains the topics to fetch metadata for.
type MetadataRequestTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic id.
	TopicID uuid.UUID
	// Name contains the topic name.
	Name *string
}

func (t *MetadataRequestTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 10 {
		if err := pe.putUUID(t.TopicID); err != nil {
			return err
		}
	}

	if t.Version >= 10 {
		if err := pe.putNullableString(t.Name); err != nil {
			return err
		}
	} else {
		if t.Name == nil {
			return fmt.Errorf("String field, Name, must not be nil in version %d", t.Version)
		}
		if err := pe.putString(*t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *MetadataRequestTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 10 {
		if t.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if t.Version >= 10 {
		if t.Name, err = pd.getNullableString(); err != nil {
			return err
		}
	} else {
		if *t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type MetadataRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains the topics to fetch metadata for.
	Topics []MetadataRequestTopic
	// AllowAutoTopicCreation contains a If this is true, the broker may auto-create topics that we requested which do not already exist, if it is configured to do so.
	AllowAutoTopicCreation bool
	// IncludeClusterAuthorizedOperations contains a Whether to include cluster authorized operations.
	IncludeClusterAuthorizedOperations bool
	// IncludeTopicAuthorizedOperations contains a Whether to include topic authorized operations.
	IncludeTopicAuthorizedOperations bool
}

func (r *MetadataRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 9 {
		pe = FlexibleEncoderFrom(pe)
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
		pe.putBool(r.AllowAutoTopicCreation)
	}

	if r.Version >= 8 && r.Version <= 10 {
		pe.putBool(r.IncludeClusterAuthorizedOperations)
	}

	if r.Version >= 8 {
		pe.putBool(r.IncludeTopicAuthorizedOperations)
	}

	if r.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *MetadataRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 9 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]MetadataRequestTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block MetadataRequestTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 4 {
		if r.AllowAutoTopicCreation, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version >= 8 && r.Version <= 10 {
		if r.IncludeClusterAuthorizedOperations, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version >= 8 {
		if r.IncludeTopicAuthorizedOperations, err = pd.getBool(); err != nil {
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

func (r *MetadataRequest) GetKey() int16 {
	return 3
}

func (r *MetadataRequest) GetVersion() int16 {
	return r.Version
}

func (r *MetadataRequest) GetHeaderVersion() int16 {
	if r.Version >= 9 {
		return 2
	}
	return 1
}

func (r *MetadataRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 12
}

func (r *MetadataRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
