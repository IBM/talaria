// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DeleteAclsMatchingACL contains the ACLs which matched this filter.
type DeleteAclsMatchingACL struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the deletion error code, or 0 if the deletion succeeded.
	ErrorCode int16
	// ErrorMessage contains the deletion error message, or null if the deletion succeeded.
	ErrorMessage *string
	// ResourceType contains the ACL resource type.
	ResourceType int8
	// ResourceName contains the ACL resource name.
	ResourceName string
	// PatternType contains the ACL resource pattern type.
	PatternType int8
	// Principal contains the ACL principal.
	Principal string
	// Host contains the ACL host.
	Host string
	// Operation contains the ACL operation.
	Operation int8
	// PermissionType contains the ACL permission type.
	PermissionType int8
}

func (m *DeleteAclsMatchingACL) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	pe.putInt16(m.ErrorCode)

	if err := pe.putNullableString(m.ErrorMessage); err != nil {
		return err
	}

	pe.putInt8(m.ResourceType)

	if err := pe.putString(m.ResourceName); err != nil {
		return err
	}

	if m.Version >= 1 {
		pe.putInt8(m.PatternType)
	}

	if err := pe.putString(m.Principal); err != nil {
		return err
	}

	if err := pe.putString(m.Host); err != nil {
		return err
	}

	pe.putInt8(m.Operation)

	pe.putInt8(m.PermissionType)

	if m.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *DeleteAclsMatchingACL) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if m.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if m.ResourceType, err = pd.getInt8(); err != nil {
		return err
	}

	if m.ResourceName, err = pd.getString(); err != nil {
		return err
	}

	if m.Version >= 1 {
		if m.PatternType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if m.Principal, err = pd.getString(); err != nil {
		return err
	}

	if m.Host, err = pd.getString(); err != nil {
		return err
	}

	if m.Operation, err = pd.getInt8(); err != nil {
		return err
	}

	if m.PermissionType, err = pd.getInt8(); err != nil {
		return err
	}

	if m.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DeleteAclsFilterResult contains the results for each filter.
type DeleteAclsFilterResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if the filter succeeded.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if the filter succeeded.
	ErrorMessage *string
	// MatchingAcls contains the ACLs which matched this filter.
	MatchingAcls []DeleteAclsMatchingACL
}

func (f *DeleteAclsFilterResult) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	pe.putInt16(f.ErrorCode)

	if err := pe.putNullableString(f.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(f.MatchingAcls)); err != nil {
		return err
	}
	for _, block := range f.MatchingAcls {
		if err := block.encode(pe, f.Version); err != nil {
			return err
		}
	}

	if f.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (f *DeleteAclsFilterResult) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if f.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numMatchingAcls int
	if numMatchingAcls, err = pd.getArrayLength(); err != nil {
		return err
	}
	f.MatchingAcls = make([]DeleteAclsMatchingACL, numMatchingAcls)
	for i := 0; i < numMatchingAcls; i++ {
		var block DeleteAclsMatchingACL
		if err := block.decode(pd, f.Version); err != nil {
			return err
		}
		f.MatchingAcls[i] = block
	}

	if f.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAclsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// FilterResults contains the results for each filter.
	FilterResults []DeleteAclsFilterResult
}

func (r *DeleteAclsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.FilterResults)); err != nil {
		return err
	}
	for _, block := range r.FilterResults {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteAclsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numFilterResults int
	if numFilterResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.FilterResults = make([]DeleteAclsFilterResult, numFilterResults)
	for i := 0; i < numFilterResults; i++ {
		var block DeleteAclsFilterResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.FilterResults[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteAclsResponse) GetKey() int16 {
	return 31
}

func (r *DeleteAclsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DeleteAclsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *DeleteAclsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *DeleteAclsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DeleteAclsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
