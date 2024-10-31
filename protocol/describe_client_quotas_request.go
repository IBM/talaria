// protocol has been generated from message format json - DO NOT EDIT
package protocol

// ComponentData contains a Filter components to apply to quota entities.
type ComponentData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EntityType contains the entity type that the filter component applies to.
	EntityType string
	// MatchType contains a How to match the entity {0 = exact name, 1 = default name, 2 = any specified name}.
	MatchType int8
	// Match contains the string to match against, or null if unused for the match type.
	Match *string
}

func (c *ComponentData) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.EntityType); err != nil {
		return err
	}

	pe.putInt8(c.MatchType)

	if err := pe.putNullableString(c.Match); err != nil {
		return err
	}

	if c.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *ComponentData) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.EntityType, err = pd.getString(); err != nil {
		return err
	}

	if c.MatchType, err = pd.getInt8(); err != nil {
		return err
	}

	if c.Match, err = pd.getNullableString(); err != nil {
		return err
	}

	if c.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClientQuotasRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Components contains a Filter components to apply to quota entities.
	Components []ComponentData
	// Strict contains a Whether the match is strict, i.e. should exclude entities with unspecified entity types.
	Strict bool
}

func (r *DescribeClientQuotasRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Components)); err != nil {
		return err
	}
	for _, block := range r.Components {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putBool(r.Strict)

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeClientQuotasRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numComponents int
	if numComponents, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numComponents > 0 {
		r.Components = make([]ComponentData, numComponents)
		for i := 0; i < numComponents; i++ {
			var block ComponentData
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Components[i] = block
		}
	}

	if r.Strict, err = pd.getBool(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeClientQuotasRequest) GetKey() int16 {
	return 48
}

func (r *DescribeClientQuotasRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeClientQuotasRequest) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 2
	}
	return 1
}

func (r *DescribeClientQuotasRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *DescribeClientQuotasRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
