// protocol has been generated from message format json - DO NOT EDIT
package protocol

// EntityData_AlterClientQuotasRequest contains the quota entity to alter.
type EntityData_AlterClientQuotasRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EntityType contains the entity type.
	EntityType string
	// EntityName contains the name of the entity, or null if the default.
	EntityName *string
}

func (e *EntityData_AlterClientQuotasRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (e *EntityData_AlterClientQuotasRequest) decode(pd packetDecoder, version int16) (err error) {
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

// OpData contains a An individual quota configuration entry to alter.
type OpData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Key contains the quota configuration key.
	Key string
	// Value contains the value to set, otherwise ignored if the value is to be removed.
	Value float64
	// Remove contains a Whether the quota configuration value should be removed, otherwise set.
	Remove bool
}

func (o *OpData) encode(pe packetEncoder, version int16) (err error) {
	o.Version = version
	if err := pe.putString(o.Key); err != nil {
		return err
	}

	pe.putFloat64(o.Value)

	pe.putBool(o.Remove)

	if o.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (o *OpData) decode(pd packetDecoder, version int16) (err error) {
	o.Version = version
	if o.Key, err = pd.getString(); err != nil {
		return err
	}

	if o.Value, err = pd.getFloat64(); err != nil {
		return err
	}

	if o.Remove, err = pd.getBool(); err != nil {
		return err
	}

	if o.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// EntryData_AlterClientQuotasRequest contains the quota configuration entries to alter.
type EntryData_AlterClientQuotasRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Entity contains the quota entity to alter.
	Entity []EntityData_AlterClientQuotasRequest
	// Ops contains a An individual quota configuration entry to alter.
	Ops []OpData
}

func (e *EntryData_AlterClientQuotasRequest) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if err := pe.putArrayLength(len(e.Entity)); err != nil {
		return err
	}
	for _, block := range e.Entity {
		if err := block.encode(pe, e.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(e.Ops)); err != nil {
		return err
	}
	for _, block := range e.Ops {
		if err := block.encode(pe, e.Version); err != nil {
			return err
		}
	}

	if e.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (e *EntryData_AlterClientQuotasRequest) decode(pd packetDecoder, version int16) (err error) {
	e.Version = version
	var numEntity int
	if numEntity, err = pd.getArrayLength(); err != nil {
		return err
	}
	e.Entity = make([]EntityData_AlterClientQuotasRequest, numEntity)
	for i := 0; i < numEntity; i++ {
		var block EntityData_AlterClientQuotasRequest
		if err := block.decode(pd, e.Version); err != nil {
			return err
		}
		e.Entity[i] = block
	}

	var numOps int
	if numOps, err = pd.getArrayLength(); err != nil {
		return err
	}
	e.Ops = make([]OpData, numOps)
	for i := 0; i < numOps; i++ {
		var block OpData
		if err := block.decode(pd, e.Version); err != nil {
			return err
		}
		e.Ops[i] = block
	}

	if e.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AlterClientQuotasRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Entries contains the quota configuration entries to alter.
	Entries []EntryData_AlterClientQuotasRequest
	// ValidateOnly contains a Whether the alteration should be validated, but not performed.
	ValidateOnly bool
}

func (r *AlterClientQuotasRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Entries)); err != nil {
		return err
	}
	for _, block := range r.Entries {
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

func (r *AlterClientQuotasRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numEntries int
	if numEntries, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Entries = make([]EntryData_AlterClientQuotasRequest, numEntries)
	for i := 0; i < numEntries; i++ {
		var block EntryData_AlterClientQuotasRequest
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Entries[i] = block
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

func (r *AlterClientQuotasRequest) GetKey() int16 {
	return 49
}

func (r *AlterClientQuotasRequest) GetVersion() int16 {
	return r.Version
}

func (r *AlterClientQuotasRequest) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 2
	}
	return 1
}

func (r *AlterClientQuotasRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *AlterClientQuotasRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
