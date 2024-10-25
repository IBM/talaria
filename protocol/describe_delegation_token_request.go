// protocol has been generated from message format json - DO NOT EDIT
package protocol

// DescribeDelegationTokenOwner contains each owner that we want to describe delegation tokens for, or null to describe all tokens.
type DescribeDelegationTokenOwner struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PrincipalType contains the owner principal type.
	PrincipalType string
	// PrincipalName contains the owner principal name.
	PrincipalName string
}

func (o *DescribeDelegationTokenOwner) encode(pe packetEncoder, version int16) (err error) {
	o.Version = version
	if err := pe.putString(o.PrincipalType); err != nil {
		return err
	}

	if err := pe.putString(o.PrincipalName); err != nil {
		return err
	}

	if o.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (o *DescribeDelegationTokenOwner) decode(pd packetDecoder, version int16) (err error) {
	o.Version = version
	if o.PrincipalType, err = pd.getString(); err != nil {
		return err
	}

	if o.PrincipalName, err = pd.getString(); err != nil {
		return err
	}

	if o.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDelegationTokenRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Owners contains each owner that we want to describe delegation tokens for, or null to describe all tokens.
	Owners []DescribeDelegationTokenOwner
}

func (r *DescribeDelegationTokenRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Owners)); err != nil {
		return err
	}
	for _, block := range r.Owners {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeDelegationTokenRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numOwners int
	if numOwners, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numOwners > 0 {
		r.Owners = make([]DescribeDelegationTokenOwner, numOwners)
		for i := 0; i < numOwners; i++ {
			var block DescribeDelegationTokenOwner
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Owners[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeDelegationTokenRequest) GetKey() int16 {
	return 41
}

func (r *DescribeDelegationTokenRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeDelegationTokenRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *DescribeDelegationTokenRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *DescribeDelegationTokenRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
