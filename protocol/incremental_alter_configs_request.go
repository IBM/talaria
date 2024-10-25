// protocol has been generated from message format json - DO NOT EDIT
package protocol

// AlterableConfig_IncrementalAlterConfigsRequest contains the configurations.
type AlterableConfig_IncrementalAlterConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the configuration key name.
	Name string
	// ConfigOperation contains the type (Set, Delete, Append, Subtract) of operation.
	ConfigOperation int8
	// Value contains the value to set for the configuration key.
	Value *string
}

func (c *AlterableConfig_IncrementalAlterConfigsRequest) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.Name); err != nil {
		return err
	}

	pe.putInt8(c.ConfigOperation)

	if err := pe.putNullableString(c.Value); err != nil {
		return err
	}

	if c.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *AlterableConfig_IncrementalAlterConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Name, err = pd.getString(); err != nil {
		return err
	}

	if c.ConfigOperation, err = pd.getInt8(); err != nil {
		return err
	}

	if c.Value, err = pd.getNullableString(); err != nil {
		return err
	}

	if c.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// AlterConfigsResource_IncrementalAlterConfigsRequest contains the incremental updates for each resource.
type AlterConfigsResource_IncrementalAlterConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceType contains the resource type.
	ResourceType int8
	// ResourceName contains the resource name.
	ResourceName string
	// Configs contains the configurations.
	Configs []AlterableConfig_IncrementalAlterConfigsRequest
}

func (r *AlterConfigsResource_IncrementalAlterConfigsRequest) encode(pe packetEncoder, version int16) (err error) {
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

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterConfigsResource_IncrementalAlterConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
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
		r.Configs = make([]AlterableConfig_IncrementalAlterConfigsRequest, numConfigs)
		for i := 0; i < numConfigs; i++ {
			var block AlterableConfig_IncrementalAlterConfigsRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Configs[i] = block
		}
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type IncrementalAlterConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Resources contains the incremental updates for each resource.
	Resources []AlterConfigsResource_IncrementalAlterConfigsRequest
	// ValidateOnly contains a True if we should validate the request, but not change the configurations.
	ValidateOnly bool
}

func (r *IncrementalAlterConfigsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
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

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *IncrementalAlterConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numResources int
	if numResources, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResources > 0 {
		r.Resources = make([]AlterConfigsResource_IncrementalAlterConfigsRequest, numResources)
		for i := 0; i < numResources; i++ {
			var block AlterConfigsResource_IncrementalAlterConfigsRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Resources[i] = block
		}
	}

	if r.ValidateOnly, err = pd.getBool(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *IncrementalAlterConfigsRequest) GetKey() int16 {
	return 44
}

func (r *IncrementalAlterConfigsRequest) GetVersion() int16 {
	return r.Version
}

func (r *IncrementalAlterConfigsRequest) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 2
	}
	return 1
}

func (r *IncrementalAlterConfigsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *IncrementalAlterConfigsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
