// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// AlterUserScramCredentialsResult contains the results for deletions and alterations, one per affected user.
type AlterUserScramCredentialsResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// User contains the user name.
	User string
	// ErrorCode contains the error code.
	ErrorCode int16
	// ErrorMessage contains the error message, if any.
	ErrorMessage *string
}

func (r *AlterUserScramCredentialsResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.User); err != nil {
		return err
	}

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *AlterUserScramCredentialsResult) decode(pd packetDecoder, version int16) (err error) {
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

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type AlterUserScramCredentialsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Results contains the results for deletions and alterations, one per affected user.
	Results []AlterUserScramCredentialsResult
}

func (r *AlterUserScramCredentialsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

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

func (r *AlterUserScramCredentialsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Results = make([]AlterUserScramCredentialsResult, numResults)
	for i := 0; i < numResults; i++ {
		var block AlterUserScramCredentialsResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Results[i] = block
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *AlterUserScramCredentialsResponse) GetKey() int16 {
	return 51
}

func (r *AlterUserScramCredentialsResponse) GetVersion() int16 {
	return r.Version
}

func (r *AlterUserScramCredentialsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *AlterUserScramCredentialsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AlterUserScramCredentialsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AlterUserScramCredentialsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
