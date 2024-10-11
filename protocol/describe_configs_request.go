// protocol has been generated from message format json - DO NOT EDIT
package protocol

// DescribeConfigsResource contains the resources whose configurations we want to describe.
type DescribeConfigsResource struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceType contains the resource type.
	ResourceType int8
	// ResourceName contains the resource name.
	ResourceName string
	// ConfigurationKeys contains the configuration keys to list, or null to list all configuration keys.
	ConfigurationKeys []string
}

func (r *DescribeConfigsResource) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt8(r.ResourceType)

	if err := pe.putString(r.ResourceName); err != nil {
		return err
	}

	if err := pe.putStringArray(r.ConfigurationKeys); err != nil {
		return err
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeConfigsResource) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ResourceType, err = pd.getInt8(); err != nil {
		return err
	}

	if r.ResourceName, err = pd.getString(); err != nil {
		return err
	}

	if r.ConfigurationKeys, err = pd.getStringArray(); err != nil {
		return err
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeConfigsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Resources contains the resources whose configurations we want to describe.
	Resources []DescribeConfigsResource
	// IncludeSynonyms contains a True if we should include all synonyms.
	IncludeSynonyms bool
	// IncludeDocumentation contains a True if we should include configuration documentation.
	IncludeDocumentation bool
}

func (r *DescribeConfigsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
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

	if r.Version >= 1 {
		pe.putBool(r.IncludeSynonyms)
	}

	if r.Version >= 3 {
		pe.putBool(r.IncludeDocumentation)
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numResources int
	if numResources, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Resources = make([]DescribeConfigsResource, numResources)
	for i := 0; i < numResources; i++ {
		var block DescribeConfigsResource
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Resources[i] = block
	}

	if r.Version >= 1 {
		if r.IncludeSynonyms, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.IncludeDocumentation, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeConfigsRequest) GetKey() int16 {
	return 32
}

func (r *DescribeConfigsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeConfigsRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *DescribeConfigsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *DescribeConfigsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
