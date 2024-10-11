// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DescribedGroupMember contains the group members.
type DescribedGroupMember struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the member ID assigned by the group coordinator.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
	// ClientID contains the client ID used in the member's latest join group request.
	ClientID string
	// ClientHost contains the client host.
	ClientHost string
	// MemberMetadata contains the metadata corresponding to the current group protocol in use.
	MemberMetadata []byte
	// MemberAssignment contains the current assignment provided by the group leader.
	MemberAssignment []byte
}

func (m *DescribedGroupMember) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	if err := pe.putString(m.MemberID); err != nil {
		return err
	}

	if m.Version >= 4 {
		if err := pe.putNullableString(m.GroupInstanceID); err != nil {
			return err
		}
	}

	if err := pe.putString(m.ClientID); err != nil {
		return err
	}

	if err := pe.putString(m.ClientHost); err != nil {
		return err
	}

	if err := pe.putBytes(m.MemberMetadata); err != nil {
		return err
	}

	if err := pe.putBytes(m.MemberAssignment); err != nil {
		return err
	}

	if m.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *DescribedGroupMember) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if m.Version >= 4 {
		if m.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if m.ClientID, err = pd.getString(); err != nil {
		return err
	}

	if m.ClientHost, err = pd.getString(); err != nil {
		return err
	}

	if m.MemberMetadata, err = pd.getBytes(); err != nil {
		return err
	}

	if m.MemberAssignment, err = pd.getBytes(); err != nil {
		return err
	}

	if m.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DescribedGroup contains each described group.
type DescribedGroup struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the describe error, or 0 if there was no error.
	ErrorCode int16
	// GroupID contains the group ID string.
	GroupID string
	// GroupState contains the group state string, or the empty string.
	GroupState string
	// ProtocolType contains the group protocol type, or the empty string.
	ProtocolType string
	// ProtocolData contains the group protocol data, or the empty string.
	ProtocolData string
	// Members contains the group members.
	Members []DescribedGroupMember
	// AuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this group.
	AuthorizedOperations int32
}

func (g *DescribedGroup) encode(pe packetEncoder, version int16) (err error) {
	g.Version = version
	pe.putInt16(g.ErrorCode)

	if err := pe.putString(g.GroupID); err != nil {
		return err
	}

	if err := pe.putString(g.GroupState); err != nil {
		return err
	}

	if err := pe.putString(g.ProtocolType); err != nil {
		return err
	}

	if err := pe.putString(g.ProtocolData); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(g.Members)); err != nil {
		return err
	}
	for _, block := range g.Members {
		if err := block.encode(pe, g.Version); err != nil {
			return err
		}
	}

	if g.Version >= 3 {
		pe.putInt32(g.AuthorizedOperations)
	}

	if g.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (g *DescribedGroup) decode(pd packetDecoder, version int16) (err error) {
	g.Version = version
	if g.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if g.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if g.GroupState, err = pd.getString(); err != nil {
		return err
	}

	if g.ProtocolType, err = pd.getString(); err != nil {
		return err
	}

	if g.ProtocolData, err = pd.getString(); err != nil {
		return err
	}

	var numMembers int
	if numMembers, err = pd.getArrayLength(); err != nil {
		return err
	}
	g.Members = make([]DescribedGroupMember, numMembers)
	for i := 0; i < numMembers; i++ {
		var block DescribedGroupMember
		if err := block.decode(pd, g.Version); err != nil {
			return err
		}
		g.Members[i] = block
	}

	if g.Version >= 3 {
		if g.AuthorizedOperations, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if g.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Groups contains each described group.
	Groups []DescribedGroup
}

func (r *DescribeGroupsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 5 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if err := pe.putArrayLength(len(r.Groups)); err != nil {
		return err
	}
	for _, block := range r.Groups {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeGroupsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 5 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numGroups int
	if numGroups, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Groups = make([]DescribedGroup, numGroups)
	for i := 0; i < numGroups; i++ {
		var block DescribedGroup
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Groups[i] = block
	}

	if r.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeGroupsResponse) GetKey() int16 {
	return 15
}

func (r *DescribeGroupsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeGroupsResponse) GetHeaderVersion() int16 {
	if r.Version >= 5 {
		return 1
	}
	return 0
}

func (r *DescribeGroupsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 5
}

func (r *DescribeGroupsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeGroupsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
