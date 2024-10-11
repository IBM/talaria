// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DescribedDelegationTokenRenewer contains a Those who are able to renew this token before it expires.
type DescribedDelegationTokenRenewer struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PrincipalType contains the renewer principal type
	PrincipalType string
	// PrincipalName contains the renewer principal name
	PrincipalName string
}

func (r *DescribedDelegationTokenRenewer) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.PrincipalType); err != nil {
		return err
	}

	if err := pe.putString(r.PrincipalName); err != nil {
		return err
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribedDelegationTokenRenewer) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.PrincipalType, err = pd.getString(); err != nil {
		return err
	}

	if r.PrincipalName, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribedDelegationToken contains the tokens.
type DescribedDelegationToken struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PrincipalType contains the token principal type.
	PrincipalType string
	// PrincipalName contains the token principal name.
	PrincipalName string
	// TokenRequesterPrincipalType contains the principal type of the requester of the token.
	TokenRequesterPrincipalType string
	// TokenRequesterPrincipalName contains the principal type of the requester of the token.
	TokenRequesterPrincipalName string
	// IssueTimestamp contains the token issue timestamp in milliseconds.
	IssueTimestamp int64
	// ExpiryTimestamp contains the token expiry timestamp in milliseconds.
	ExpiryTimestamp int64
	// MaxTimestamp contains the token maximum timestamp length in milliseconds.
	MaxTimestamp int64
	// TokenID contains the token ID.
	TokenID string
	// Hmac contains the token HMAC.
	Hmac []byte
	// Renewers contains a Those who are able to renew this token before it expires.
	Renewers []DescribedDelegationTokenRenewer
}

func (t *DescribedDelegationToken) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.PrincipalType); err != nil {
		return err
	}

	if err := pe.putString(t.PrincipalName); err != nil {
		return err
	}

	if t.Version >= 3 {
		if err := pe.putString(t.TokenRequesterPrincipalType); err != nil {
			return err
		}
	}

	if t.Version >= 3 {
		if err := pe.putString(t.TokenRequesterPrincipalName); err != nil {
			return err
		}
	}

	pe.putInt64(t.IssueTimestamp)

	pe.putInt64(t.ExpiryTimestamp)

	pe.putInt64(t.MaxTimestamp)

	if err := pe.putString(t.TokenID); err != nil {
		return err
	}

	if err := pe.putBytes(t.Hmac); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(t.Renewers)); err != nil {
		return err
	}
	for _, block := range t.Renewers {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *DescribedDelegationToken) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.PrincipalType, err = pd.getString(); err != nil {
		return err
	}

	if t.PrincipalName, err = pd.getString(); err != nil {
		return err
	}

	if t.Version >= 3 {
		if t.TokenRequesterPrincipalType, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 3 {
		if t.TokenRequesterPrincipalName, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.IssueTimestamp, err = pd.getInt64(); err != nil {
		return err
	}

	if t.ExpiryTimestamp, err = pd.getInt64(); err != nil {
		return err
	}

	if t.MaxTimestamp, err = pd.getInt64(); err != nil {
		return err
	}

	if t.TokenID, err = pd.getString(); err != nil {
		return err
	}

	if t.Hmac, err = pd.getBytes(); err != nil {
		return err
	}

	var numRenewers int
	if numRenewers, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Renewers = make([]DescribedDelegationTokenRenewer, numRenewers)
	for i := 0; i < numRenewers; i++ {
		var block DescribedDelegationTokenRenewer
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Renewers[i] = block
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDelegationTokenResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// Tokens contains the tokens.
	Tokens []DescribedDelegationToken
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
}

func (r *DescribeDelegationTokenResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.Tokens)); err != nil {
		return err
	}
	for _, block := range r.Tokens {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putInt32(r.ThrottleTimeMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeDelegationTokenResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numTokens int
	if numTokens, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Tokens = make([]DescribedDelegationToken, numTokens)
	for i := 0; i < numTokens; i++ {
		var block DescribedDelegationToken
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Tokens[i] = block
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

func (r *DescribeDelegationTokenResponse) GetKey() int16 {
	return 41
}

func (r *DescribeDelegationTokenResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeDelegationTokenResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *DescribeDelegationTokenResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *DescribeDelegationTokenResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeDelegationTokenResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
