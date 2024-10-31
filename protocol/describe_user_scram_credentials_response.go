// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// CredentialInfo contains the mechanism and related information associated with the user's SCRAM credentials.
type CredentialInfo struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Mechanism contains the SCRAM mechanism.
	Mechanism int8
	// Iterations contains the number of iterations used in the SCRAM credential.
	Iterations int32
}

func (c *CredentialInfo) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	pe.putInt8(c.Mechanism)

	pe.putInt32(c.Iterations)

	pe.putUVarint(0)
	return nil
}

func (c *CredentialInfo) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Mechanism, err = pd.getInt8(); err != nil {
		return err
	}

	if c.Iterations, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// DescribeUserScramCredentialsResult contains the results for descriptions, one per user.
type DescribeUserScramCredentialsResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// User contains the user name.
	User string
	// ErrorCode contains the user-level error code.
	ErrorCode int16
	// ErrorMessage contains the user-level error message, if any.
	ErrorMessage *string
	// CredentialInfos contains the mechanism and related information associated with the user's SCRAM credentials.
	CredentialInfos []CredentialInfo
}

func (r *DescribeUserScramCredentialsResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.User); err != nil {
		return err
	}

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.CredentialInfos)); err != nil {
		return err
	}
	for _, block := range r.CredentialInfos {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeUserScramCredentialsResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.User, err = pd.getString(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numCredentialInfos int
	if numCredentialInfos, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numCredentialInfos > 0 {
		r.CredentialInfos = make([]CredentialInfo, numCredentialInfos)
		for i := 0; i < numCredentialInfos; i++ {
			var block CredentialInfo
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.CredentialInfos[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeUserScramCredentialsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the message-level error code, 0 except for user authorization or infrastructure issues.
	ErrorCode int16
	// ErrorMessage contains the message-level error message, if any.
	ErrorMessage *string
	// Results contains the results for descriptions, one per user.
	Results []DescribeUserScramCredentialsResult
}

func (r *DescribeUserScramCredentialsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeUserScramCredentialsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResults > 0 {
		r.Results = make([]DescribeUserScramCredentialsResult, numResults)
		for i := 0; i < numResults; i++ {
			var block DescribeUserScramCredentialsResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Results[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeUserScramCredentialsResponse) GetKey() int16 {
	return 50
}

func (r *DescribeUserScramCredentialsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeUserScramCredentialsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeUserScramCredentialsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeUserScramCredentialsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeUserScramCredentialsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
