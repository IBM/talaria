// protocol has been generated from message format json - DO NOT EDIT
package protocol

type DescribeAclsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceTypeFilter contains the resource type.
	ResourceTypeFilter int8
	// ResourceNameFilter contains the resource name, or null to match any resource name.
	ResourceNameFilter *string
	// PatternTypeFilter contains the resource pattern to match.
	PatternTypeFilter int8
	// PrincipalFilter contains the principal to match, or null to match any principal.
	PrincipalFilter *string
	// HostFilter contains the host to match, or null to match any host.
	HostFilter *string
	// Operation contains the operation to match.
	Operation int8
	// PermissionType contains the permission type to match.
	PermissionType int8
}

func (r *DescribeAclsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt8(r.ResourceTypeFilter)

	if err := pe.putNullableString(r.ResourceNameFilter); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putInt8(r.PatternTypeFilter)
	}

	if err := pe.putNullableString(r.PrincipalFilter); err != nil {
		return err
	}

	if err := pe.putNullableString(r.HostFilter); err != nil {
		return err
	}

	pe.putInt8(r.Operation)

	pe.putInt8(r.PermissionType)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeAclsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ResourceTypeFilter, err = pd.getInt8(); err != nil {
		return err
	}

	if r.ResourceNameFilter, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.PatternTypeFilter, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if r.PrincipalFilter, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.HostFilter, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.Operation, err = pd.getInt8(); err != nil {
		return err
	}

	if r.PermissionType, err = pd.getInt8(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeAclsRequest) GetKey() int16 {
	return 29
}

func (r *DescribeAclsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeAclsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *DescribeAclsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *DescribeAclsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
