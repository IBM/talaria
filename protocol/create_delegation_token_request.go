// protocol has been generated from message format json - DO NOT EDIT
package protocol

// CreatableRenewers contains a A list of those who are allowed to renew this token before it expires.
type CreatableRenewers struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PrincipalType contains the type of the Kafka principal.
	PrincipalType string
	// PrincipalName contains the name of the Kafka principal.
	PrincipalName string
}

func (r *CreatableRenewers) encode(pe packetEncoder, version int16) (err error) {
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

func (r *CreatableRenewers) decode(pd packetDecoder, version int16) (err error) {
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

type CreateDelegationTokenRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// OwnerPrincipalType contains the principal type of the owner of the token. If it's null it defaults to the token request principal.
	OwnerPrincipalType *string
	// OwnerPrincipalName contains the principal name of the owner of the token. If it's null it defaults to the token request principal.
	OwnerPrincipalName *string
	// Renewers contains a A list of those who are allowed to renew this token before it expires.
	Renewers []CreatableRenewers
	// MaxLifetimeMs contains the maximum lifetime of the token in milliseconds, or -1 to use the server side default.
	MaxLifetimeMs int64
}

func (r *CreateDelegationTokenRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		if err := pe.putNullableString(r.OwnerPrincipalType); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if err := pe.putNullableString(r.OwnerPrincipalName); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Renewers)); err != nil {
		return err
	}
	for _, block := range r.Renewers {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putInt64(r.MaxLifetimeMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreateDelegationTokenRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.OwnerPrincipalType, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.OwnerPrincipalName, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	var numRenewers int
	if numRenewers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numRenewers > 0 {
		r.Renewers = make([]CreatableRenewers, numRenewers)
		for i := 0; i < numRenewers; i++ {
			var block CreatableRenewers
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Renewers[i] = block
		}
	}

	if r.MaxLifetimeMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateDelegationTokenRequest) GetKey() int16 {
	return 38
}

func (r *CreateDelegationTokenRequest) GetVersion() int16 {
	return r.Version
}

func (r *CreateDelegationTokenRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *CreateDelegationTokenRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *CreateDelegationTokenRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
