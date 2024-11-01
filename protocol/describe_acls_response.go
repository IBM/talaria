// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// AclDescription contains the ACLs.
type AclDescription struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Principal contains the ACL principal.
	Principal string
	// Host contains the ACL host.
	Host string
	// Operation contains the ACL operation.
	Operation int8
	// PermissionType contains the ACL permission type.
	PermissionType int8
}

func (a *AclDescription) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putString(a.Principal); err != nil {
		return err
	}

	if err := pe.putString(a.Host); err != nil {
		return err
	}

	pe.putInt8(a.Operation)

	pe.putInt8(a.PermissionType)

	if a.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *AclDescription) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.Principal, err = pd.getString(); err != nil {
		return err
	}

	if a.Host, err = pd.getString(); err != nil {
		return err
	}

	if a.Operation, err = pd.getInt8(); err != nil {
		return err
	}

	if a.PermissionType, err = pd.getInt8(); err != nil {
		return err
	}

	if a.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribeAclsResource contains each Resource that is referenced in an ACL.
type DescribeAclsResource struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceType contains the resource type.
	ResourceType int8
	// ResourceName contains the resource name.
	ResourceName string
	// PatternType contains the resource pattern type.
	PatternType int8
	// Acls contains the ACLs.
	Acls []AclDescription
}

func (r *DescribeAclsResource) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt8(r.ResourceType)

	if err := pe.putString(r.ResourceName); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putInt8(r.PatternType)
	}

	if err := pe.putArrayLength(len(r.Acls)); err != nil {
		return err
	}
	for _, block := range r.Acls {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeAclsResource) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ResourceType, err = pd.getInt8(); err != nil {
		return err
	}

	if r.ResourceName, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.PatternType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	var numAcls int
	if numAcls, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numAcls > 0 {
		r.Acls = make([]AclDescription, numAcls)
		for i := 0; i < numAcls; i++ {
			var block AclDescription
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Acls[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAclsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// Resources contains each Resource that is referenced in an ACL.
	Resources []DescribeAclsResource
}

func (r *DescribeAclsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Resources)); err != nil {
		return err
	}
	for _, block := range r.Resources {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeAclsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
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

	var numResources int
	if numResources, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResources > 0 {
		r.Resources = make([]DescribeAclsResource, numResources)
		for i := 0; i < numResources; i++ {
			var block DescribeAclsResource
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Resources[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeAclsResponse) GetKey() int16 {
	return 29
}

func (r *DescribeAclsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeAclsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *DescribeAclsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *DescribeAclsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeAclsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
