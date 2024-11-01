// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// UpdatableFeatureResult contains a Results for each feature update.
type UpdatableFeatureResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Feature contains the name of the finalized feature.
	Feature string
	// ErrorCode contains the feature update error code or `0` if the feature update succeeded.
	ErrorCode int16
	// ErrorMessage contains the feature update error, or `null` if the feature update succeeded.
	ErrorMessage *string
}

func (r *UpdatableFeatureResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.Feature); err != nil {
		return err
	}

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *UpdatableFeatureResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Feature, err = pd.getString(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type UpdateFeaturesResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level error code, or `0` if there was no top-level error.
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or `null` if there was no top-level error.
	ErrorMessage *string
	// Results contains a Results for each feature update.
	Results []UpdatableFeatureResult
}

func (r *UpdateFeaturesResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *UpdateFeaturesResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResults > 0 {
		r.Results = make([]UpdatableFeatureResult, numResults)
		for i := 0; i < numResults; i++ {
			var block UpdatableFeatureResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Results[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *UpdateFeaturesResponse) GetKey() int16 {
	return 57
}

func (r *UpdateFeaturesResponse) GetVersion() int16 {
	return r.Version
}

func (r *UpdateFeaturesResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *UpdateFeaturesResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *UpdateFeaturesResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *UpdateFeaturesResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
