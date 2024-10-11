// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ExpireDelegationTokenRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Hmac contains the HMAC of the delegation token to be expired.
	Hmac []byte
	// ExpiryTimePeriodMs contains the expiry time period in milliseconds.
	ExpiryTimePeriodMs int64
}

func (r *ExpireDelegationTokenRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putBytes(r.Hmac); err != nil {
		return err
	}

	pe.putInt64(r.ExpiryTimePeriodMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ExpireDelegationTokenRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Hmac, err = pd.getBytes(); err != nil {
		return err
	}

	if r.ExpiryTimePeriodMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ExpireDelegationTokenRequest) GetKey() int16 {
	return 40
}

func (r *ExpireDelegationTokenRequest) GetVersion() int16 {
	return r.Version
}

func (r *ExpireDelegationTokenRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *ExpireDelegationTokenRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *ExpireDelegationTokenRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
