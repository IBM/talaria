// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// EntityData_AlterClientQuotasResponse contains the quota entity to alter.
type EntityData_AlterClientQuotasResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EntityType contains the entity type.
	EntityType string
	// EntityName contains the name of the entity, or null if the default.
	EntityName *string
}

func (e *EntityData_AlterClientQuotasResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (e *EntityData_AlterClientQuotasResponse) decode(pd packetDecoder, version int16) (err error) {
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

// EntryData_AlterClientQuotasResponse contains the quota configuration entries to alter.
type EntryData_AlterClientQuotasResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or `0` if the quota alteration succeeded.
	ErrorCode int16
	// ErrorMessage contains the error message, or `null` if the quota alteration succeeded.
	ErrorMessage *string
	// Entity contains the quota entity to alter.
	Entity []EntityData_AlterClientQuotasResponse
}

func (e *EntryData_AlterClientQuotasResponse) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	pe.putInt16(e.ErrorCode)

	if err := pe.putNullableString(e.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(e.Entity)); err != nil {
		return err
	}
	for _, block := range e.Entity {
		if err := block.encode(pe, e.Version); err != nil {
			return err
		}
	}

	if e.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (e *EntryData_AlterClientQuotasResponse) decode(pd packetDecoder, version int16) (err error) {
	e.Version = version
	if e.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if e.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numEntity int
	if numEntity, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numEntity > 0 {
		e.Entity = make([]EntityData_AlterClientQuotasResponse, numEntity)
		for i := 0; i < numEntity; i++ {
			var block EntityData_AlterClientQuotasResponse
			if err := block.decode(pd, e.Version); err != nil {
				return err
			}
			e.Entity[i] = block
		}
	}

	if e.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AlterClientQuotasResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Entries contains the quota configuration entries to alter.
	Entries []EntryData_AlterClientQuotasResponse
}

func (r *AlterClientQuotasResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

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

func (r *AlterClientQuotasResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numEntries int
	if numEntries, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numEntries > 0 {
		r.Entries = make([]EntryData_AlterClientQuotasResponse, numEntries)
		for i := 0; i < numEntries; i++ {
			var block EntryData_AlterClientQuotasResponse
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

func (r *AlterClientQuotasResponse) GetKey() int16 {
	return 49
}

func (r *AlterClientQuotasResponse) GetVersion() int16 {
	return r.Version
}

func (r *AlterClientQuotasResponse) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 1
	}
	return 0
}

func (r *AlterClientQuotasResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *AlterClientQuotasResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AlterClientQuotasResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
