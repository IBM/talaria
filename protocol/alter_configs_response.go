// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// AlterConfigsResourceResponse_AlterConfigsResponse contains the responses for each resource.
type AlterConfigsResourceResponse_AlterConfigsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the resource error code.
	ErrorCode int16
	// ErrorMessage contains the resource error message, or null if there was no error.
	ErrorMessage *string
	// ResourceType contains the resource type.
	ResourceType int8
	// ResourceName contains the resource name.
	ResourceName string
}

func (r *AlterConfigsResourceResponse_AlterConfigsResponse) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	pe.putInt8(r.ResourceType)

	if err := pe.putString(r.ResourceName); err != nil {
		return err
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterConfigsResourceResponse_AlterConfigsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.ResourceType, err = pd.getInt8(); err != nil {
		return err
	}

	if r.ResourceName, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AlterConfigsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Responses contains the responses for each resource.
	Responses []AlterConfigsResourceResponse_AlterConfigsResponse
}

func (r *AlterConfigsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterConfigsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResponses > 0 {
		r.Responses = make([]AlterConfigsResourceResponse_AlterConfigsResponse, numResponses)
		for i := 0; i < numResponses; i++ {
			var block AlterConfigsResourceResponse_AlterConfigsResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Responses[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AlterConfigsResponse) GetKey() int16 {
	return 33
}

func (r *AlterConfigsResponse) GetVersion() int16 {
	return r.Version
}

func (r *AlterConfigsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *AlterConfigsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *AlterConfigsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AlterConfigsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
