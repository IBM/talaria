// protocol has been generated from message format json - DO NOT EDIT
package protocol

// DeleteAclsFilter contains the filters to use when deleting ACLs.
type DeleteAclsFilter struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceTypeFilter contains the resource type.
	ResourceTypeFilter int8
	// ResourceNameFilter contains the resource name.
	ResourceNameFilter *string
	// PatternTypeFilter contains the pattern type.
	PatternTypeFilter int8
	// PrincipalFilter contains the principal filter, or null to accept all principals.
	PrincipalFilter *string
	// HostFilter contains the host filter, or null to accept all hosts.
	HostFilter *string
	// Operation contains the ACL operation.
	Operation int8
	// PermissionType contains the permission type.
	PermissionType int8
}

func (f *DeleteAclsFilter) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	pe.putInt8(f.ResourceTypeFilter)

	if err := pe.putNullableString(f.ResourceNameFilter); err != nil {
		return err
	}

	if f.Version >= 1 {
		pe.putInt8(f.PatternTypeFilter)
	}

	if err := pe.putNullableString(f.PrincipalFilter); err != nil {
		return err
	}

	if err := pe.putNullableString(f.HostFilter); err != nil {
		return err
	}

	pe.putInt8(f.Operation)

	pe.putInt8(f.PermissionType)

	if f.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (f *DeleteAclsFilter) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.ResourceTypeFilter, err = pd.getInt8(); err != nil {
		return err
	}

	if f.ResourceNameFilter, err = pd.getNullableString(); err != nil {
		return err
	}

	if f.Version >= 1 {
		if f.PatternTypeFilter, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if f.PrincipalFilter, err = pd.getNullableString(); err != nil {
		return err
	}

	if f.HostFilter, err = pd.getNullableString(); err != nil {
		return err
	}

	if f.Operation, err = pd.getInt8(); err != nil {
		return err
	}

	if f.PermissionType, err = pd.getInt8(); err != nil {
		return err
	}

	if f.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAclsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Filters contains the filters to use when deleting ACLs.
	Filters []DeleteAclsFilter
}

func (r *DeleteAclsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Filters)); err != nil {
		return err
	}
	for _, block := range r.Filters {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteAclsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numFilters int
	if numFilters, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numFilters > 0 {
		r.Filters = make([]DeleteAclsFilter, numFilters)
		for i := 0; i < numFilters; i++ {
			var block DeleteAclsFilter
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Filters[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteAclsRequest) GetKey() int16 {
	return 31
}

func (r *DeleteAclsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DeleteAclsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *DeleteAclsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *DeleteAclsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
