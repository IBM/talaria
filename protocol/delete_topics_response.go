// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	"fmt"
	uuid "github.com/google/uuid"
	"time"
)

// DeletableTopicResult contains the results for each topic we tried to delete.
type DeletableTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name
	Name *string
	// TopicID contains a the unique topic ID
	TopicID uuid.UUID
	// ErrorCode contains the deletion error, or 0 if the deletion succeeded.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
}

func (r *DeletableTopicResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		if err := pe.putNullableString(r.Name); err != nil {
			return err
		}
	} else {
		if r.Name == nil {
			return fmt.Errorf("String field, Name, must not be nil in version %d", r.Version)
		}
		if err := pe.putString(*r.Name); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		if err := pe.putUUID(r.TopicID); err != nil {
			return err
		}
	}

	pe.putInt16(r.ErrorCode)

	if r.Version >= 5 {
		if err := pe.putNullableString(r.ErrorMessage); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeletableTopicResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		if r.Name, err = pd.getNullableString(); err != nil {
			return err
		}
	} else {
		def := ""
		r.Name = &def

		if *r.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		if r.TopicID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 5 {
		if r.ErrorMessage, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteTopicsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Responses contains the results for each topic we tried to delete.
	Responses []DeletableTopicResult
}

func (r *DeleteTopicsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteTopicsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResponses > 0 {
		r.Responses = make([]DeletableTopicResult, numResponses)
		for i := 0; i < numResponses; i++ {
			var block DeletableTopicResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Responses[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteTopicsResponse) GetKey() int16 {
	return 20
}

func (r *DeleteTopicsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DeleteTopicsResponse) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 1
	}
	return 0
}

func (r *DeleteTopicsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 6
}

func (r *DeleteTopicsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DeleteTopicsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
