// protocol has been generated from message format json - DO NOT EDIT
package protocol

// FeatureUpdateKey contains the list of updates to finalized features.
type FeatureUpdateKey struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Feature contains the name of the finalized feature to be updated.
	Feature string
	// MaxVersionLevel contains the new maximum version level for the finalized feature. A value >= 1 is valid. A value < 1, is special, and can be used to request the deletion of the finalized feature.
	MaxVersionLevel int16
	// AllowDowngrade contains the downgrade request will fail if the new maximum version level is a value that's not lower than the existing maximum finalized version level.
	AllowDowngrade bool
	// UpgradeType contains a Determine which type of upgrade will be performed: 1 will perform an upgrade only (default), 2 is safe downgrades only (lossless), 3 is unsafe downgrades (lossy).
	UpgradeType int8
}

func (f *FeatureUpdateKey) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	if err := pe.putString(f.Feature); err != nil {
		return err
	}

	pe.putInt16(f.MaxVersionLevel)

	if f.Version == 0 {
		pe.putBool(f.AllowDowngrade)
	}

	if f.Version >= 1 {
		pe.putInt8(f.UpgradeType)
	}

	pe.putUVarint(0)
	return nil
}

func (f *FeatureUpdateKey) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.Feature, err = pd.getString(); err != nil {
		return err
	}

	if f.MaxVersionLevel, err = pd.getInt16(); err != nil {
		return err
	}

	if f.Version == 0 {
		if f.AllowDowngrade, err = pd.getBool(); err != nil {
			return err
		}
	}

	if f.Version >= 1 {
		if f.UpgradeType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type UpdateFeaturesRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// timeoutMs contains a How long to wait in milliseconds before timing out the request.
	timeoutMs int32
	// FeatureUpdates contains the list of updates to finalized features.
	FeatureUpdates []FeatureUpdateKey
	// ValidateOnly contains a True if we should validate the request, but not perform the upgrade or downgrade.
	ValidateOnly bool
}

func (r *UpdateFeaturesRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.timeoutMs)

	if err := pe.putArrayLength(len(r.FeatureUpdates)); err != nil {
		return err
	}
	for _, block := range r.FeatureUpdates {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putBool(r.ValidateOnly)
	}

	pe.putUVarint(0)
	return nil
}

func (r *UpdateFeaturesRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.timeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numFeatureUpdates int
	if numFeatureUpdates, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numFeatureUpdates > 0 {
		r.FeatureUpdates = make([]FeatureUpdateKey, numFeatureUpdates)
		for i := 0; i < numFeatureUpdates; i++ {
			var block FeatureUpdateKey
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.FeatureUpdates[i] = block
		}
	}

	if r.Version >= 1 {
		if r.ValidateOnly, err = pd.getBool(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *UpdateFeaturesRequest) GetKey() int16 {
	return 57
}

func (r *UpdateFeaturesRequest) GetVersion() int16 {
	return r.Version
}

func (r *UpdateFeaturesRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *UpdateFeaturesRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *UpdateFeaturesRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
