// protocol has been generated from message format json - DO NOT EDIT
package protocol

type RenewDelegationTokenRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Hmac contains the HMAC of the delegation token to be renewed.
	Hmac []byte
	// RenewPeriodMs contains the renewal time period in milliseconds.
	RenewPeriodMs int64
}

func (r *RenewDelegationTokenRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putBytes(r.Hmac); err != nil {
		return err
	}

	pe.putInt64(r.RenewPeriodMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *RenewDelegationTokenRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Hmac, err = pd.getBytes(); err != nil {
		return err
	}

	if r.RenewPeriodMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *RenewDelegationTokenRequest) GetKey() int16 {
	return 39
}

func (r *RenewDelegationTokenRequest) GetVersion() int16 {
	return r.Version
}

func (r *RenewDelegationTokenRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *RenewDelegationTokenRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *RenewDelegationTokenRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
