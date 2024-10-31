// protocol has been generated from message format json - DO NOT EDIT
package protocol

// AlterableConfig_AlterConfigsRequest contains the configurations.
type AlterableConfig_AlterConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the configuration key name.
	Name string
	// Value contains the value to set for the configuration key.
	Value *string
}

func (c *AlterableConfig_AlterConfigsRequest) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.Name); err != nil {
		return err
	}

	if err := pe.putNullableString(c.Value); err != nil {
		return err
	}

	if c.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *AlterableConfig_AlterConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Name, err = pd.getString(); err != nil {
		return err
	}

	if c.Value, err = pd.getNullableString(); err != nil {
		return err
	}

	if c.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// AlterConfigsResource_AlterConfigsRequest contains the updates for each resource.
type AlterConfigsResource_AlterConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceType contains the resource type.
	ResourceType int8
	// ResourceName contains the resource name.
	ResourceName string
	// Configs contains the configurations.
	Configs []AlterableConfig_AlterConfigsRequest
}

func (r *AlterConfigsResource_AlterConfigsRequest) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt8(r.ResourceType)

	if err := pe.putString(r.ResourceName); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Configs)); err != nil {
		return err
	}
	for _, block := range r.Configs {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterConfigsResource_AlterConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ResourceType, err = pd.getInt8(); err != nil {
		return err
	}

	if r.ResourceName, err = pd.getString(); err != nil {
		return err
	}

	var numConfigs int
	if numConfigs, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numConfigs > 0 {
		r.Configs = make([]AlterableConfig_AlterConfigsRequest, numConfigs)
		for i := 0; i < numConfigs; i++ {
			var block AlterableConfig_AlterConfigsRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Configs[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AlterConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Resources contains the updates for each resource.
	Resources []AlterConfigsResource_AlterConfigsRequest
	// ValidateOnly contains a True if we should validate the request, but not change the configurations.
	ValidateOnly bool
}

func (r *AlterConfigsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Resources)); err != nil {
		return err
	}
	for _, block := range r.Resources {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putBool(r.ValidateOnly)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numResources int
	if numResources, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResources > 0 {
		r.Resources = make([]AlterConfigsResource_AlterConfigsRequest, numResources)
		for i := 0; i < numResources; i++ {
			var block AlterConfigsResource_AlterConfigsRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Resources[i] = block
		}
	}

	if r.ValidateOnly, err = pd.getBool(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AlterConfigsRequest) GetKey() int16 {
	return 33
}

func (r *AlterConfigsRequest) GetVersion() int16 {
	return r.Version
}

func (r *AlterConfigsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *AlterConfigsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *AlterConfigsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
