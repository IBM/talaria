// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DescribeConfigsSynonym contains the synonyms for this configuration key.
type DescribeConfigsSynonym struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the synonym name.
	Name string
	// Value contains the synonym value.
	Value *string
	// Source contains the synonym source.
	Source int8
}

func (s *DescribeConfigsSynonym) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	if s.Version >= 1 {
		if err := pe.putString(s.Name); err != nil {
			return err
		}
	}

	if s.Version >= 1 {
		if err := pe.putNullableString(s.Value); err != nil {
			return err
		}
	}

	if s.Version >= 1 {
		pe.putInt8(s.Source)
	}

	if s.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (s *DescribeConfigsSynonym) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.Version >= 1 {
		if s.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if s.Version >= 1 {
		if s.Value, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if s.Version >= 1 {
		if s.Source, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if s.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribeConfigsResourceResult contains each listed configuration.
type DescribeConfigsResourceResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the configuration name.
	Name string
	// Value contains the configuration value.
	Value *string
	// ReadOnly contains a True if the configuration is read-only.
	ReadOnly bool
	// IsDefault contains a True if the configuration is not set.
	IsDefault bool
	// ConfigSource contains the configuration source.
	ConfigSource int8
	// IsSensitive contains a True if this configuration is sensitive.
	IsSensitive bool
	// Synonyms contains the synonyms for this configuration key.
	Synonyms []DescribeConfigsSynonym
	// ConfigType contains the configuration data type. Type can be one of the following values - BOOLEAN, STRING, INT, SHORT, LONG, DOUBLE, LIST, CLASS, PASSWORD
	ConfigType int8
	// Documentation contains the configuration documentation.
	Documentation *string
}

func (c *DescribeConfigsResourceResult) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.Name); err != nil {
		return err
	}

	if err := pe.putNullableString(c.Value); err != nil {
		return err
	}

	pe.putBool(c.ReadOnly)

	if c.Version == 0 {
		pe.putBool(c.IsDefault)
	}

	if c.Version >= 1 {
		pe.putInt8(c.ConfigSource)
	}

	pe.putBool(c.IsSensitive)

	if c.Version >= 1 {
		if err := pe.putArrayLength(len(c.Synonyms)); err != nil {
			return err
		}
		for _, block := range c.Synonyms {
			if err := block.encode(pe, c.Version); err != nil {
				return err
			}
		}
	}

	if c.Version >= 3 {
		pe.putInt8(c.ConfigType)
	}

	if c.Version >= 3 {
		if err := pe.putNullableString(c.Documentation); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *DescribeConfigsResourceResult) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Name, err = pd.getString(); err != nil {
		return err
	}

	if c.Value, err = pd.getNullableString(); err != nil {
		return err
	}

	if c.ReadOnly, err = pd.getBool(); err != nil {
		return err
	}

	if c.Version == 0 {
		if c.IsDefault, err = pd.getBool(); err != nil {
			return err
		}
	}

	if c.Version >= 1 {
		if c.ConfigSource, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if c.IsSensitive, err = pd.getBool(); err != nil {
		return err
	}

	if c.Version >= 1 {
		var numSynonyms int
		if numSynonyms, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numSynonyms > 0 {
			c.Synonyms = make([]DescribeConfigsSynonym, numSynonyms)
			for i := 0; i < numSynonyms; i++ {
				var block DescribeConfigsSynonym
				if err := block.decode(pd, c.Version); err != nil {
					return err
				}
				c.Synonyms[i] = block
			}
		}
	}

	if c.Version >= 3 {
		if c.ConfigType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if c.Version >= 3 {
		if c.Documentation, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribeConfigsResult contains the results for each resource.
type DescribeConfigsResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if we were able to successfully describe the configurations.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if we were able to successfully describe the configurations.
	ErrorMessage *string
	// ResourceType contains the resource type.
	ResourceType int8
	// ResourceName contains the resource name.
	ResourceName string
	// Configs contains each listed configuration.
	Configs []DescribeConfigsResourceResult
}

func (r *DescribeConfigsResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

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

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeConfigsResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

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
		r.Configs = make([]DescribeConfigsResourceResult, numConfigs)
		for i := 0; i < numConfigs; i++ {
			var block DescribeConfigsResourceResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Configs[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeConfigsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Results contains the results for each resource.
	Results []DescribeConfigsResult
}

func (r *DescribeConfigsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeConfigsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResults > 0 {
		r.Results = make([]DescribeConfigsResult, numResults)
		for i := 0; i < numResults; i++ {
			var block DescribeConfigsResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Results[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeConfigsResponse) GetKey() int16 {
	return 32
}

func (r *DescribeConfigsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeConfigsResponse) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 1
	}
	return 0
}

func (r *DescribeConfigsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *DescribeConfigsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeConfigsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
