// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type CreateDelegationTokenResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top-level error, or zero if there was no error.
	ErrorCode int16
	// PrincipalType contains the principal type of the token owner.
	PrincipalType string
	// PrincipalName contains the name of the token owner.
	PrincipalName string
	// TokenRequesterPrincipalType contains the principal type of the requester of the token.
	TokenRequesterPrincipalType string
	// TokenRequesterPrincipalName contains the principal type of the requester of the token.
	TokenRequesterPrincipalName string
	// IssueTimestampMs contains a When this token was generated.
	IssueTimestampMs int64
	// ExpiryTimestampMs contains a When this token expires.
	ExpiryTimestampMs int64
	// MaxTimestampMs contains the maximum lifetime of this token.
	MaxTimestampMs int64
	// TokenID contains the token UUID.
	TokenID string
	// Hmac contains a HMAC of the delegation token.
	Hmac []byte
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
}

func (r *CreateDelegationTokenResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putString(r.PrincipalType); err != nil {
		return err
	}

	if err := pe.putString(r.PrincipalName); err != nil {
		return err
	}

	if r.Version >= 3 {
		if err := pe.putString(r.TokenRequesterPrincipalType); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if err := pe.putString(r.TokenRequesterPrincipalName); err != nil {
			return err
		}
	}

	pe.putInt64(r.IssueTimestampMs)

	pe.putInt64(r.ExpiryTimestampMs)

	pe.putInt64(r.MaxTimestampMs)

	if err := pe.putString(r.TokenID); err != nil {
		return err
	}

	if err := pe.putBytes(r.Hmac); err != nil {
		return err
	}

	pe.putInt32(r.ThrottleTimeMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreateDelegationTokenResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.PrincipalType, err = pd.getString(); err != nil {
		return err
	}

	if r.PrincipalName, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.TokenRequesterPrincipalType, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.TokenRequesterPrincipalName, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.IssueTimestampMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ExpiryTimestampMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.MaxTimestampMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.TokenID, err = pd.getString(); err != nil {
		return err
	}

	if r.Hmac, err = pd.getBytes(); err != nil {
		return err
	}

	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateDelegationTokenResponse) GetKey() int16 {
	return 38
}

func (r *CreateDelegationTokenResponse) GetVersion() int16 {
	return r.Version
}

func (r *CreateDelegationTokenResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *CreateDelegationTokenResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *CreateDelegationTokenResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *CreateDelegationTokenResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
