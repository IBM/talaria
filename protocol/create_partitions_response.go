// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// CreatePartitionsTopicResult contains the partition creation results for each topic.
type CreatePartitionsTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// ErrorCode contains the result error, or zero if there was no error.
	ErrorCode int16
	// ErrorMessage contains the result message, or null if there was no error.
	ErrorMessage *string
}

func (r *CreatePartitionsTopicResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.Name); err != nil {
		return err
	}

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreatePartitionsTopicResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Name, err = pd.getString(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePartitionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Results contains the partition creation results for each topic.
	Results []CreatePartitionsTopicResult
}

func (r *CreatePartitionsResponse) encode(pe packetEncoder) (err error) {
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

func (r *CreatePartitionsResponse) decode(pd packetDecoder, version int16) (err error) {
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
	r.Results = make([]CreatePartitionsTopicResult, numResults)
	for i := 0; i < numResults; i++ {
		var block CreatePartitionsTopicResult
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

func (r *CreatePartitionsResponse) GetKey() int16 {
	return 37
}

func (r *CreatePartitionsResponse) GetVersion() int16 {
	return r.Version
}

func (r *CreatePartitionsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *CreatePartitionsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *CreatePartitionsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *CreatePartitionsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
