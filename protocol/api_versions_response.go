// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// ApiVersion contains the APIs supported by the broker.
type ApiVersion struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ApiKey contains the API index.
	ApiKey int16
	// MinVersion contains the minimum supported version, inclusive.
	MinVersion int16
	// MaxVersion contains the maximum supported version, inclusive.
	MaxVersion int16
}

func (a *ApiVersion) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt16(a.ApiKey)

	pe.putInt16(a.MinVersion)

	pe.putInt16(a.MaxVersion)

	if a.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *ApiVersion) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.ApiKey, err = pd.getInt16(); err != nil {
		return err
	}

	if a.MinVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if a.MaxVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if a.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// SupportedFeatureKey contains a Features supported by the broker.
type SupportedFeatureKey struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the feature.
	Name string
	// MinVersion contains the minimum supported version for the feature.
	MinVersion int16
	// MaxVersion contains the maximum supported version for the feature.
	MaxVersion int16
}

func (s *SupportedFeatureKey) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	if s.Version >= 3 {
		if err := pe.putString(s.Name); err != nil {
			return err
		}
	}

	if s.Version >= 3 {
		pe.putInt16(s.MinVersion)
	}

	if s.Version >= 3 {
		pe.putInt16(s.MaxVersion)
	}

	return nil
}

func (s *SupportedFeatureKey) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.Version >= 3 {
		if s.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if s.Version >= 3 {
		if s.MinVersion, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if s.Version >= 3 {
		if s.MaxVersion, err = pd.getInt16(); err != nil {
			return err
		}
	}

	return nil
}

// FinalizedFeatureKey contains the information is valid only if FinalizedFeaturesEpoch >= 0.
type FinalizedFeatureKey struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the feature.
	Name string
	// MaxVersionLevel contains the cluster-wide finalized max version level for the feature.
	MaxVersionLevel int16
	// MinVersionLevel contains the cluster-wide finalized min version level for the feature.
	MinVersionLevel int16
}

func (f *FinalizedFeatureKey) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	if f.Version >= 3 {
		if err := pe.putString(f.Name); err != nil {
			return err
		}
	}

	if f.Version >= 3 {
		pe.putInt16(f.MaxVersionLevel)
	}

	if f.Version >= 3 {
		pe.putInt16(f.MinVersionLevel)
	}

	return nil
}

func (f *FinalizedFeatureKey) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.Version >= 3 {
		if f.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if f.Version >= 3 {
		if f.MaxVersionLevel, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if f.Version >= 3 {
		if f.MinVersionLevel, err = pd.getInt16(); err != nil {
			return err
		}
	}

	return nil
}

type ApiVersionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top-level error code.
	ErrorCode int16
	// ApiKeys contains the APIs supported by the broker.
	ApiKeys []ApiVersion
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// SupportedFeatures contains a Features supported by the broker.
	SupportedFeatures []SupportedFeatureKey
	// FinalizedFeaturesEpoch contains the monotonically increasing epoch for the finalized features information. Valid values are >= 0. A value of -1 is special and represents unknown epoch.
	FinalizedFeaturesEpoch int64
	// FinalizedFeatures contains the information is valid only if FinalizedFeaturesEpoch >= 0.
	FinalizedFeatures []FinalizedFeatureKey
	// ZkMigrationReady contains a Set by a KRaft controller if the required configurations for ZK migration are present
	ZkMigrationReady bool
}

func (r *ApiVersionsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.ApiKeys)); err != nil {
		return err
	}
	for _, block := range r.ApiKeys {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ApiVersionsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numApiKeys int
	if numApiKeys, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numApiKeys > 0 {
		r.ApiKeys = make([]ApiVersion, numApiKeys)
		for i := 0; i < numApiKeys; i++ {
			var block ApiVersion
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.ApiKeys[i] = block
		}
	}

	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ApiVersionsResponse) GetKey() int16 {
	return 18
}

func (r *ApiVersionsResponse) GetVersion() int16 {
	return r.Version
}

func (r *ApiVersionsResponse) GetHeaderVersion() int16 {
	return 0
}

func (r *ApiVersionsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *ApiVersionsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ApiVersionsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
