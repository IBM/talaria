// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// EntityData_DescribeClientQuotasResponse contains the quota entity description.
type EntityData_DescribeClientQuotasResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EntityType contains the entity type.
	EntityType string
	// EntityName contains the entity name, or null if the default.
	EntityName *string
}

func (e *EntityData_DescribeClientQuotasResponse) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if err := pe.putString(e.EntityType); err != nil {
		return err
	}

	if err := pe.putNullableString(e.EntityName); err != nil {
		return err
	}

	if e.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (e *EntityData_DescribeClientQuotasResponse) decode(pd packetDecoder, version int16) (err error) {
	e.Version = version
	if e.EntityType, err = pd.getString(); err != nil {
		return err
	}

	if e.EntityName, err = pd.getNullableString(); err != nil {
		return err
	}

	if e.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// ValueData contains the quota values for the entity.
type ValueData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Key contains the quota configuration key.
	Key string
	// Value contains the quota configuration value.
	Value float64
}

func (v *ValueData) encode(pe packetEncoder, version int16) (err error) {
	v.Version = version
	if err := pe.putString(v.Key); err != nil {
		return err
	}

	pe.putFloat64(v.Value)

	if v.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (v *ValueData) decode(pd packetDecoder, version int16) (err error) {
	v.Version = version
	if v.Key, err = pd.getString(); err != nil {
		return err
	}

	if v.Value, err = pd.getFloat64(); err != nil {
		return err
	}

	if v.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// EntryData_DescribeClientQuotasResponse contains a A result entry.
type EntryData_DescribeClientQuotasResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Entity contains the quota entity description.
	Entity []EntityData_DescribeClientQuotasResponse
	// Values contains the quota values for the entity.
	Values []ValueData
}

func (e *EntryData_DescribeClientQuotasResponse) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if err := pe.putArrayLength(len(e.Entity)); err != nil {
		return err
	}
	for _, block := range e.Entity {
		if err := block.encode(pe, e.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(e.Values)); err != nil {
		return err
	}
	for _, block := range e.Values {
		if err := block.encode(pe, e.Version); err != nil {
			return err
		}
	}

	if e.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (e *EntryData_DescribeClientQuotasResponse) decode(pd packetDecoder, version int16) (err error) {
	e.Version = version
	var numEntity int
	if numEntity, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numEntity > 0 {
		e.Entity = make([]EntityData_DescribeClientQuotasResponse, numEntity)
		for i := 0; i < numEntity; i++ {
			var block EntityData_DescribeClientQuotasResponse
			if err := block.decode(pd, e.Version); err != nil {
				return err
			}
			e.Entity[i] = block
		}
	}

	var numValues int
	if numValues, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numValues > 0 {
		e.Values = make([]ValueData, numValues)
		for i := 0; i < numValues; i++ {
			var block ValueData
			if err := block.decode(pd, e.Version); err != nil {
				return err
			}
			e.Values[i] = block
		}
	}

	if e.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClientQuotasResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or `0` if the quota description succeeded.
	ErrorCode int16
	// ErrorMessage contains the error message, or `null` if the quota description succeeded.
	ErrorMessage *string
	// Entries contains a A result entry.
	Entries []EntryData_DescribeClientQuotasResponse
}

func (r *DescribeClientQuotasResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Entries)); err != nil {
		return err
	}
	for _, block := range r.Entries {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeClientQuotasResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numEntries int
	if numEntries, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numEntries > 0 {
		r.Entries = make([]EntryData_DescribeClientQuotasResponse, numEntries)
		for i := 0; i < numEntries; i++ {
			var block EntryData_DescribeClientQuotasResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Entries[i] = block
		}
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeClientQuotasResponse) GetKey() int16 {
	return 48
}

func (r *DescribeClientQuotasResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeClientQuotasResponse) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 1
	}
	return 0
}

func (r *DescribeClientQuotasResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *DescribeClientQuotasResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeClientQuotasResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
